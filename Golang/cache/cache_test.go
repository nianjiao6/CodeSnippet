package cache

import (
	"Golang/cache/cacheByChan"
	"Golang/cache/cacheByLock"
	"testing"
)

var links = []string{"https://www.baidu.com", "https://books.studygolang.com/", "https://mp.weixin.qq.com/", "https://www.baidu.com",
	"https://www.baidu.com", "https://books.studygolang.com/", "https://mp.weixin.qq.com/", "https://www.baidu.com"}

func TestLockCache(t *testing.T) {
	cacheByLock.TestCache(links)
}
func TestChanCache(t *testing.T) {
	cacheByChan.TestCache(links)
}
