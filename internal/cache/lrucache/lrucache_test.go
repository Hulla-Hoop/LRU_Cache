package lrucache_test

import (
	"os"
	"testing"

	"github.com/Hulla-Hoop/LRU_Cache/internal/cache/lrucache"
)

var c *lrucache.LRUCache

func TestMain(m *testing.M) {
	c = lrucache.New(3)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestLRUCache(t *testing.T) {
	data := []struct {
		id    string
		key   string
		value string
		want  bool
	}{
		{"1", "key1", "value1", true},
		{"2", "key2", "value2", true},
		{"3", "key3", "value3", true},
		{"4", "key4", "value4", true},
		{"5", "key4", "value4", false},
		{"6", "key5", "value5", true},
	}
	for _, d := range data {
		t.Run(d.id, func(t *testing.T) {
			result := c.Add(d.key, d.value)
			if result != d.want {
				t.Errorf("got %v, want %v", result, d.want)
			}
		})
	}

}
