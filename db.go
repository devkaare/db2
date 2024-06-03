package db

import (
	"encoding/json"
	"errors"
	"log"
	"os"
    "time")

var dbFile string
var dbFileCache []map[string]interface{}

// This function is not intended to be used outside this file, hence not included in docs
func DoesFileExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return true, err
}

func LoadCache(dbFilePath string) {
    dbFile = dbFilePath // Set the global dbFile variable to the provided path
	result, _ := DoesFileExist(dbFile)
	if !result {
		os.Create(dbFile)
		os.WriteFile(dbFile, []byte("[]"), 0644)
	}

	data, err := os.ReadFile(dbFile)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(data, &dbFileCache); err != nil {
		log.Fatal(err)
	}

	// Start a goroutine to save the cache every 5 minutes
    go func() {
        ticker := time.NewTicker(5 * time.Minute)
        defer ticker.Stop()

        for range ticker.C {
            SaveCache()
        }
    }()
}

func SaveCache() {
	newData, err := json.Marshal(dbFileCache)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(dbFile, newData, 0644); err != nil {
		log.Fatal(err)
	}
}

func AddToCache(key string, value map[string]interface{}) {
    for i, allData := range dbFileCache {
        if _, ok := allData[key]; ok {
            dbFileCache[i][key] = append(dbFileCache[i][key].([]interface{}), value)
            return
        }
    }

    // If no existing key is found, create a new one
    dbFileCache = append(dbFileCache, map[string]interface{}{key: []interface{}{value}})
}

func SearchCache[T any](key string, field string, search T) map[string]interface{} {
	var foundData []map[string]interface{}
	for _, allData := range dbFileCache {
		if data, ok := allData[key].([]interface{}); ok {
			for _, item := range data {
				if itemMap, ok := item.(map[string]interface{}); ok {
					switch v := any(search).(type) {
					case int:
						if itemMap[field] == v {
							foundData = append(foundData, itemMap)
						}
					case string:
						if itemMap[field] == v {
							foundData = append(foundData, itemMap)
						}
					}
				}
			}
		}
	}
    if len(foundData) == 0 {
        return nil
    }
	return foundData[0]
}

func DeleteFromCache[T any](key string, field string, search T) {
	for i, allData := range dbFileCache {
		if data, ok := allData[key].([]interface{}); ok {
			for j, item := range data {
				if itemMap, ok := item.(map[string]interface{}); ok {
					if itemMap[field] == any(search) {
						data = append(data[:j], data[j+1:]...)
						allData[key] = data
						dbFileCache[i] = allData
						return
					}
				}
			}
		}
	}
}
