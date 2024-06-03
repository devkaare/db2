<h1>Not for production use and not recommended for large datasets!</h1>
<div>
  <p>
    This <code>db2</code> package provides basic functions for managing
    key-value pairs stored in a JSON file.
  </p>
</div>
<div>
  <p>
    <strong>Important:</strong><br />
    Call <code>LoadCache(dbFilePath string)</code> at the start of your main
    function to initialize the cache with the JSON file located at
    <code>dbFilePath</code>, and <code>defer SaveCache()</code> to save changes
    when the program exits. Additionally, a background process automatically
    saves the cache every 5 minutes.
  </p>
</div>
<div>
  <h3>Functions:</h3>
  <p>
    <strong>1. LoadCache(dbFilePath string):</strong><br />
    Loads the cache from the specified JSON file. Call this at the start of your
    main function.
  </p>
  <p>
    <strong>2. SaveCache():</strong><br />
    Saves the cache to the JSON file. Defer this at the start of your main
    function.
  </p>
  <p>
    <strong>3. AddToCache(key string, value map[string]interface{}):</strong
    ><br />
    Adds a value to the cache under the specified key.
  </p>
  <p>
    <strong
      >4. SearchCache[T any](key string, field string, search T)
      map[string]interface{}:</strong
    ><br />
    Searches the cache for an item matching the specified key and field. Returns
    the first matching item.
  </p>
  <p>
    <strong
      >5. DeleteFromCache[T any](key string, field string, search T):</strong
    ><br />
    Deletes an item from the cache based on the specified key and field.
  </p>
</div>
