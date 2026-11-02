package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://domain.com",
			val: []byte("testdata"),
		},
		{
			key: "https://domain.com/path",
			val: []byte("othertestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			data, ok := cache.Get(c.key)
			if !ok {
				t.Fatalf("Expected to find key")
			}
			if string(data) != string(c.val) {
				t.Fatalf("Expected to find value")
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://domain.com", []byte("testdata"))

	_, ok := cache.Get("https://domain.com")
	if !ok {
		t.Fatalf("Expected to find key")
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://domain.com")
	if ok {
		t.Fatalf("Expected not to find key")
	}
}
