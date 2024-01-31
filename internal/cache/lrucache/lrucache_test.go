package lrucache_test

import (
	"os"
	"testing"

	"github.com/Hulla-Hoop/LRU_Cache/internal/cache/lrucache"
)

// один кэш чтоб править всеми
var c *lrucache.LRUCache

func TestMain(m *testing.M) {
	//один инициализированный кэш
	c = lrucache.New(3)
	exitCode := m.Run()
	os.Exit(exitCode)
}

// тест Add(key string, value string) bool
func TestLRUCacheAdd(t *testing.T) {
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

// тест Get(key string) (string, bool)
func TestLRUCacheGet(t *testing.T) {
	data := []struct {
		id    string
		key   string
		value string
		want  bool
	}{
		{"1", "key1", "", false},
		{"2", "key2", "", false},
		{"3", "key3", "value3", true},
		{"4", "key4", "value4", true},
		{"5", "key4", "value4", true},
		{"6", "key5", "value5", true},
	}
	for _, d := range data {
		c.Add(d.key, d.value)
	}

	for _, d := range data {
		t.Run(d.id, func(t *testing.T) {

			value, result := c.Get(d.key)
			if result != d.want {
				t.Errorf("got %v|%s , want %s|%v", result, value, d.value, d.want)
			}
		})
	}

}

// тест Remove(key string) bool
func TestLRUCacheRemove(t *testing.T) {
	data := []struct {
		id    string
		key   string
		value string
		want  bool
	}{
		{"1", "key1", "value1", true},
		{"2", "key2", "value2", true},
		{"3", "key3", "value3", true},
		{"4", "key4", "", false},
		{"5", "key4", "", false},
		{"6", "key5", "", false},
	}

	c.Add(data[0].key, data[0].value)
	c.Add(data[1].key, data[1].value)
	c.Add(data[2].key, data[2].value)

	for _, d := range data {
		t.Run(d.id, func(t *testing.T) {
			result := c.Remove(d.key)
			if result != d.want {
				t.Errorf("got %v , want %v", result, d.want)
			}
		})
	}
}
