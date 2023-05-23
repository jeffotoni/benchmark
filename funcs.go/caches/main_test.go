package main

import (
	"log"
	"strconv"
	"testing"
	"time"

	gc1 "github.com/jeffotoni/gcache"
	gc2 "github.com/jeffotoni/gcache/v2"
	gocache "github.com/patrickmn/go-cache"
)

var c = gc1.New(10 * time.Second)

var c2 = gc2.New[string, int](time.Duration(time.Minute), 0)

// var c3 = gc2.New[string, string](time.Duration(time.Minute), 0)

var cache = gocache.New(10*time.Second, 1*time.Minute)

func BenchmarkGcacheSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		c.Set(key, i)
	}
}

func BenchmarkGcacheSetGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		c.Set(key, i)
		i, ok := c.Get(key)
		if !ok {
			log.Printf("Não encontrei: %v", i)
		}
	}
}

func BenchmarkGcacheSet2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		c2.Set(key, i, time.Duration(time.Minute))
	}
}
func BenchmarkGcacheSetGet2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		c2.Set(key, i, time.Duration(10*time.Second))
		i, ok := c2.Get(key)
		if ok != nil {
			log.Printf("Não encontrei: %v i=%v", ok, i)
		}
	}
}

func BenchmarkGo_cacheSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		cache.Set(key, i, time.Duration(5*time.Second))
	}
}

func BenchmarkGo_cacheSetGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		cache.Set(key, i, time.Duration(5*time.Second))
		i, ok := cache.Get(key)
		if !ok {
			log.Printf("Não encontrei: %v", i)
		}
	}
}
