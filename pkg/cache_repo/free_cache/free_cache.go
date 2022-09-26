package freecache

import (
	cacherepo "micro_service_phone/pkg/cache_repo"
	"micro_service_phone/pkg/file_logger"
	"sync"

	"github.com/coocood/freecache"
)

type cacheRepository struct {
	cache             *freecache.Cache
	defaultExpireTime int
	sync.Mutex
	cacherepo.ICacheRepository
}

// --------------------------------------------------------------------------------------
// get stored value by key
func (r *cacheRepository) Get(key []byte) ([]byte, error) {
	// lock repository from concurent access
	r.Lock()
	defer r.Unlock()
	return r.cache.Get(key)
}

// --------------------------------------------------------------------------------------
// store `value` with key `key` and default value of expiration time
func (r *cacheRepository) Set(key []byte, value []byte) error {
	// lock repository from concurent access
	r.Lock()
	defer r.Unlock()
	return r.cache.Set(key, value, r.defaultExpireTime)
}

// --------------------------------------------------------------------------------------
// store `value` with key `key` and expiration time `expireIn`
func (r *cacheRepository) SetWithExpireValue(key []byte, value []byte, expireIn int) error {
	// lock repository from concurent access
	r.Lock()
	defer r.Unlock()
	return r.cache.Set(key, value, expireIn)
}

// --------------------------------------------------------------------------------------
// delete stored value by key
func (r *cacheRepository) Del(key []byte) bool {
	// lock repository from concurent access
	r.Lock()
	defer r.Unlock()
	return r.cache.Del(key)
}

// --------------------------------------------------------------------------------------
func InitCacheRepository(defaultExpireTime uint, defaultCacheSize uint, logger *file_logger.FileLogger) cacherepo.ICacheRepository {
	logger.Debugln("Start cache repository initialization")
	// In bytes, where 1024 * 1024 represents a single Megabyte, and 10 * 1024*1024 represents 10 Megabytes.
	cacheSize := defaultCacheSize * 1024 * 1024
	logger.Debugf("cacheSize=%d, defaultExpireTime=%d", cacheSize, defaultExpireTime)
	cacheInstance := freecache.NewCache(int(cacheSize))

	cacheRepo := cacheRepository{
		cache:             cacheInstance,
		defaultExpireTime: int(defaultExpireTime),
	}

	logger.Debugln("Cache repository initialized")
	return &cacheRepo
}
