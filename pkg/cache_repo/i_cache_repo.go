package cacherepo

type ICacheRepository interface {
	// get stored value by key
	Get(key []byte) ([]byte, error)

	// store `value` with key `key` and expiration time `expireIn`
	SetWithExpireValue(key []byte, value []byte, expireIn int) error

	// store `value` with key `key` and default value of expiration time
	Set(key []byte, value []byte) error

	// delete stored value by key
	Del(key []byte) bool
}
