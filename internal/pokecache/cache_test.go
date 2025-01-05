package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	interval := 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("test data"),
		},
		{
			key: "https://example.com/path",
			val: []byte("path test data"),
		},
		{
			key: "",
			val: []byte("emptykeydata"),
		},
		{
			key: "emptyvalue",
			val: []byte{},
		},
		{
			key: "https://example.com",
			val: []byte("test data overwritten"),
		},
		{
			key: "largedata",
			val: make([]byte, 1<<20),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find a key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find a value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("test data"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected key to not exist")
		return
	}
}
