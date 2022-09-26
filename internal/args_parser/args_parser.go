package argsparser

import "flag"

// settings for logging level
var LogLevelValue = flag.String("l", "debug", "Log level [panic|fatal|error|warn|info|debug|trace]")

// settings for cache expire time
var CacheExpireTime = flag.Uint("ce", 3600, "Default cache expire time")

// settings for cache size
var Cachesize = flag.Uint("cs", 30, "Cache size in megabytes")

// settings for cache size
var ServerPort = flag.Uint("p", 8000, "Server port")

// initialize command-line arguments
func Init() {
	flag.Parse()
}
