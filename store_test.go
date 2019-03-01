package kvbench

import (
	"encoding/binary"
	"os"
	"testing"
)

var stores = []struct {
	Name    string
	Path    string
	Factory func(path string, fsync bool) (Store, error)
}{
	{"badger", "badger", newBadgerStore},
	{"bbolt", "bbolt.db", newBboltStore},
	{"bolt", "bolt.db", newBboltStore},
	{"leveldb", "leveldb.db", newLevelDBStore},
	{"kv", "kv.db", newKVStore},
	{"buntdb", "buntdb.db", newBuntdbStore},
	{"btree", "btree.db", newBTreeStore},
	{"btree/memory", ":memory:", newBTreeStore},
	{"map", "map.db", newMapStore},
	{"map/memory", ":memory:", newMapStore},
}

func prefixKey(i int) []byte {
	r := make([]byte, 8)
	binary.BigEndian.PutUint64(r, uint64(i))
	return r
}

func wrapfsync(fn func(*testing.T, Store, bool), store Store, fsync bool) func(*testing.T) {
	return func(t *testing.T) {
		fn(t, store, fsync)
	}
}
func TestStore_fsync(t *testing.T) {
	for _, s := range stores {
		store, err := s.Factory(s.Path, true)
		if err != nil {
			os.RemoveAll(s.Path)
			t.Fatal(err)
		}
		t.Run(s.Name, wrapfsync(testStore, store, true))
		os.RemoveAll(s.Path)
	}
}

func TestStore_nofsync(t *testing.T) {
	for _, s := range stores {
		store, err := s.Factory(s.Path, false)
		if err != nil {
			os.RemoveAll(s.Path)
			t.Fatal(err)
		}
		t.Run(s.Name, wrapfsync(testStore, store, false))
		os.RemoveAll(s.Path)
	}
}

func testStore(t *testing.T, store Store, fsync bool) {
	v := make([]byte, 256)

	defer store.Close()

	count := 1000
	t.Run("set", func(tt *testing.T) {
		for i := 0; i < count; i++ {
			err := store.Set(prefixKey(i), v)
			if err != nil {
				tt.Fatalf("failed to set key %d: %v", i, err)
			}
		}
	})

	t.Run("get", func(tt *testing.T) {
		for i := 0; i < count; i++ {
			_, ok, err := store.Get(prefixKey(i))
			if err != nil {
				tt.Logf("failed to get key %d: %v", i, err)
			}
			if !ok {
				tt.Logf("the key %d does not exist", i)
			}
		}
	})
}
