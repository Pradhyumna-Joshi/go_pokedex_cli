package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {

	cases := []struct {
		name string
		key  string
		val  []byte
	}{
		{
			name: "Test Case 1",
			key:  "https://example.com",
			val:  []byte("testdata1"),
		},
		{
			name: "Test Case 2",
			key:  "https://example.com/path",
			val:  []byte("testdata2"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cache := NewCache(5 * time.Second)
			cache.Add(tc.key, tc.val)
			val, ok := cache.Get(tc.key)

			if !ok {
				t.Errorf("Expected to find key")
				return
			}

			if string(val) != string(tc.val) {
				t.Errorf("Expected to find value")
				return
			}

		})
	}

}

func TestReapLoop(t *testing.T) {
	interval := 5 * time.Millisecond
	cache := NewCache(interval)
	cache.Add("https://example.com", []byte("testdata"))

	if _, ok := cache.Get("https://example.com"); !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(7 * time.Millisecond)

	if _, ok := cache.Get("https://example.com"); ok {
		t.Errorf("expected to not find key")
		return
	}
}
