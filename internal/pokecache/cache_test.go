package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cases := []struct {
			key string
			val []byte
		}{
			{
				key: "https://example.com",
				val: []byte("testdata"),
			},
			{
				key: "https://example.com/path",
				val: []byte("moretestdata"),
			},
		}

	interval := 20 * time.Second
	for _, tc := range cases {
		testCache := NewCache(interval)
		if testCache == nil {
			t.Fatal("failed to create new cache")
		}
		testCache.Add(tc.key, tc.val)
		value, found := testCache.Get(tc.key)
		if !found {
			t.Fatalf("expected to find key %q but did not", tc.key)
		}
		if string(value) != string(tc.val) {
			t.Errorf("expected values to match: got %q, want %q", string(value), string(tc.val))
		}
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 10 * time.Millisecond
	const someTime = 5 * time.Millisecond
	const fullWaitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}
	time.Sleep(someTime)
	cache.Add("http://second.com", []byte("secondTestData"))

	time.Sleep(fullWaitTime)
	_, ok = cache.Get("http://second.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
