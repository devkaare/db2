package db

import (
	"encoding/json"
	"errors"
	"log"
	"os"
    "time"
)

var dbFile = "db.json"
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

func LoadCache() {
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
	dbFileCache = append(dbFileCache, map[string]interface{}{key: value})
}

func SearchCache[T any](key string, field string, search T) []map[string]interface{} {
	var foundData []map[string]interface{}
	for _, allData := range dbFileCache {
		if data, ok := allData[key].(map[string]interface{}); ok {
			switch v := any(search).(type) {
			case int:
				if data[field] == v {
					foundData = append(foundData, data)
					break
				}
			case string:
				if data[field] == v {
					foundData = append(foundData, data)
					break
				}
			}
		}
	}
	return foundData
}

func DeleteFromCache[T any](key string, field string, search T) {
	var searchIndex int
	for i, allData := range dbFileCache {
		if data, ok := allData[key].(map[string]interface{}); ok {
			if data[field] == any(search) {
				searchIndex = i
				break
			}
		}
	}

	if searchIndex > 0 {
		dbFileCache = append(dbFileCache[:searchIndex], dbFileCache[searchIndex+1:]...)
	}
}
