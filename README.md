# LRU-Cache

Least Recently Used (LRU) Cache implementation in GoLang that provides a fixed-size cache with fast retrieval and storage of key-value pairs The cache automatically evicts the least recently used elements when it reaches its capacity.

## Constructor

The Constructor(capacity int) function creates a new LRUCache with the specified capacity. It returns an instance of the LRUCache.

## Get(key int) int

The Get(key int) function retrieves the value associated with the given key from the cache. If the key is not found, it returns -1. If the key is found, it returns the corresponding value and updates the access order.

## Put(key int, value int)

The Put(key int, value int) function inserts a new key-value pair into the cache or updates the value of an existing key. If the cache is already at its capacity, it evicts the least recently used element before inserting the new one.
