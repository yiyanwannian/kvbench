package kvbench

import (
	"encoding/binary"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
)

var count = flag.Int("count", 1000, "item count for test")
var valueSize = flag.Int("*valueSize", 1024, "item value size for test")

type entry struct {
	Key   []byte
	Value []byte
	Meta  byte
}

var entries = GenerateDatas(*count)

var stores = []struct {
	Name    string
	Path    string
	Factory func(path string, fsync bool) (Store, error)
}{
	{"badger", "data/badger.db", NewBadgerStore},
	{"bbolt", "data/bbolt.db", NewBboltStore},
	{"bolt", "data/bolt.db", NewBboltStore},
	{"leveldb", "data/leveldb.db", NewLevelDBStore},
	{"kv", "data/kv.db", NewKVStore},
	{"buntdb", "data/buntdb.db", NewBuntdbStore},
	{"rocksdb", "data/rocksdb.db", NewRocksdbStore},
	{"pebble", "data/pebble.db", NewPebbleStore},
	{"pogreb", "data/pogreb.db", NewPogrebStore},
	{"btree", "data/btree.db", NewBTreeStore},
	{"btree/memory", ":memory:", NewBTreeStore},
	{"nutsdb", "data/nutsdb.db", NewNutsdbStore},
	{"map", "data/map.db", NewMapStore},
	{"map/memory", ":memory:", NewMapStore},
}

func prefixKey(i int) []byte {
	r := make([]byte, 8)
	binary.BigEndian.PutUint64(r, uint64(i))
	return r
}

func fillEntryWithIndex(e *entry, index int) {
	k := rand.Intn(*count * 100)
	key := fmt.Sprintf("vsz=%036d-k=%010d-%010d", *valueSize, k, index) // 64 bytes.
	if cap(e.Key) < len(key) {
		e.Key = make([]byte, 2*len(key))
	}
	e.Key = e.Key[:len(key)]
	copy(e.Key, key)

	rCnt := *valueSize
	p := make([]byte, rCnt)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < rCnt; i++ {
		p[i] = ' ' + byte(r.Intn('~'-' '+1))
	}
	e.Value = p[:*valueSize]
}

func GenerateDatas(num int) []*entry {
	mockEntries := make([]*entry, 0 , *count)
	for i := 0; i < num; i ++ {
		e := new(entry)
		fillEntryWithIndex(e, i)
		mockEntries = append(mockEntries, e)
	}

	return mockEntries
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
			//os.RemoveAll(s.Path)
			t.Fatal(err)
		}
		t.Run(s.Name, wrapfsync(testStore, store, true))
		//os.RemoveAll(s.Path)
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
	//v := make([]byte, 256)

	defer store.Close()

	t.Run("set", func(tt *testing.T) {
		for i := 0; i < *count; i++ {
			err := store.Set(entries[i].Key, entries[i].Value)
			//err := store.Set(prefixKey(i), v)
			if err != nil {
				tt.Fatalf("failed to set key %d: %v", i, err)
			}
		}
	})

	t.Run("get", func(tt *testing.T) {
		for i := 0; i < *count; i++ {
			_, ok, err := store.Get(entries[i].Key)
			if err != nil {
				tt.Fatalf("failed to get key %d: %v", i, err)
			}
			if !ok {
				tt.Fatalf("the key %d does not exist", i)
			}
		}
	})
}
