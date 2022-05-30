package cache

import (
	lru "github.com/hashicorp/golang-lru"
	"sync"
)

var cache map[string]*lru.Cache

var once sync.Once

//单例模式
func getInstance() map[string]*lru.Cache {
	once.Do(func() {
		cache = make(map[string]*lru.Cache)
	})
	return cache
}

//插入缓存
func CreateCache(mark string, num int) {
	cache, _ := lru.New(num)
	mapCache := getInstance()
	mapCache[mark] = cache
}

//存入缓存
func SetCacheValue(mark string, key string, value interface{}) {
	mapCache := getInstance()
	cache := mapCache[mark]
	if cache == nil && IsAutoCreateCache {
		cache, _ = lru.New(AutoCreateCacheNum)
		cache.Add(key, value)
		mapCache[mark] = cache
	} else {
		cache.Add(key, value)
	}
}

//获得缓存Value
func GetCacheValue(mark string, key string) interface{} {
	mapCache := getInstance()
	cache := mapCache[mark]
	re, _ := cache.Get(key)
	return re
}

//获得缓存所有keys
func GetCacheKeys(mark string) []interface{} {
	mapCache := getInstance()
	cache := mapCache[mark]
	return cache.Keys()
}

//获得cache高度
func GetCacheLen(mark string) int {
	mapCache := getInstance()
	cache := mapCache[mark]
	return cache.Len()
}

//移除value
func RemoveCacheValue(mark string, key string) {
	mapCache := getInstance()
	cache := mapCache[mark]
	cache.Remove(key)
}

//删除缓存
func DeleteCache(mark string) {
	mapCache := getInstance()
	cache := mapCache[mark]
	if cache != nil {
		delete(mapCache, mark)
	}
}

func Start() {
	getInstance()
}

func Stop() {
	mapCache := getInstance()
	for k, _ := range mapCache {
		delete(mapCache, k)
	}
}
