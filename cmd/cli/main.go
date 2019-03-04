package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/rcrowley/go-metrics"

	"github.com/smallnest/kvbench"
	"github.com/smallnest/log"
)

var (
	n      = flag.Int("n", 1000000, "count")
	c      = flag.Int("c", runtime.GOMAXPROCS(-1), "concurrent goroutines")
	size   = flag.Int("size", 256, "data size")
	fsync  = flag.Bool("fsync", false, "fsync")
	memory = flag.Bool("memory", false, "fsync")
	s      = flag.String("s", "map", "store type")
)

var (
	setRate      = metrics.GetOrRegisterTimer("set", nil)
	getRate      = metrics.GetOrRegisterTimer("get", nil)
	setMixedRate = metrics.GetOrRegisterTimer("setMixed", nil)
	getMixedRate = metrics.GetOrRegisterTimer("getMixed", nil)
	delRate      = metrics.GetOrRegisterTimer("del", nil)
)

func main() {
	flag.Parse()

	var path string
	if *memory {
		path = ":memory:"
	}
	store, path, err := getStore(*s, *fsync, path)
	if err != nil {
		panic(err)
	}
	if !*memory {
		defer os.RemoveAll(path)
	}

	data := make([]byte, *size)
	numPerG := *n / (*c)

	// test set
	{
		var wg sync.WaitGroup
		wg.Add(*c)

		for j := 0; j < *c; j++ {
			index := uint64(j)
			go func() {
				i := index
				for k := 0; k < numPerG; k++ {
					now := time.Now()
					store.Set(genKey(i), data)
					setRate.UpdateSince(now)
					i += index
				}
				wg.Done()
			}()
		}
		wg.Wait()
		fmt.Printf("set rate: %d, mean: %d ns, min: %d ns, max: %d ns\n",
			int64(setRate.Rate1()), int64(setRate.Mean()), int64(setRate.Min()), int64(setRate.Max()))
	}

	// test get
	{
		var wg sync.WaitGroup
		wg.Add(*c)

		for j := 0; j < *c; j++ {
			index := uint64(j)
			go func() {
				i := index
				for k := 0; k < numPerG; k++ {
					now := time.Now()
					store.Get(genKey(i))
					getRate.UpdateSince(now)
					i += index
				}
				wg.Done()
			}()
		}
		wg.Wait()
		fmt.Printf("get rate: %d, mean: %d ns, min: %d ns, max: %d ns\n",
			int64(getRate.Rate1()), int64(getRate.Mean()), int64(getRate.Min()), int64(getRate.Max()))
	}

	// test multiple get/one set
	{
		var wg sync.WaitGroup
		wg.Add(*c)

		ch := make(chan struct{})
		go func() {
			i := uint64(0)
			for {
				select {
				case <-ch:
					return
				default:
					now := time.Now()
					store.Set(genKey(i), data)
					setMixedRate.UpdateSince(now)
					i++
				}
			}
		}()
		for j := 0; j < *c; j++ {
			index := uint64(j)
			go func() {

				i := index
				for k := 0; k < numPerG; k++ {
					now := time.Now()
					store.Get(genKey(i))
					getMixedRate.UpdateSince(now)
					i += index
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(ch)
		fmt.Printf("setmixed rate: %d, mean: %d ns, min: %d ns, max: %d ns\n",
			int64(setMixedRate.Rate1()), int64(setMixedRate.Mean()), int64(setMixedRate.Min()), int64(setMixedRate.Max()))
		fmt.Printf("getmixed rate: %d, mean: %d ns, min: %d ns, max: %d ns\n",
			int64(getMixedRate.Rate1()), int64(getMixedRate.Mean()), int64(getMixedRate.Min()), int64(getMixedRate.Max()))
	}

	// test del
	{
		var wg sync.WaitGroup
		wg.Add(*c)

		for j := 0; j < *c; j++ {
			index := uint64(j)
			go func() {
				i := index
				for k := 0; k < numPerG; k++ {
					now := time.Now()
					store.Del(genKey(i))
					delRate.UpdateSince(now)
					i += index
				}
				wg.Done()
			}()
		}
		wg.Wait()
		fmt.Printf("del rate: %d, mean: %d ns, min: %d ns, max: %d ns\n",
			int64(delRate.Rate1()), int64(delRate.Mean()), int64(delRate.Min()), int64(delRate.Max()))
	}
}

func genKey(i uint64) []byte {
	r := make([]byte, 9)
	r[0] = 'k'
	binary.BigEndian.PutUint64(r[1:], i)
	return r
}

func getStore(s string, fsync bool, path string) (kvbench.Store, string, error) {
	var store kvbench.Store
	var err error
	switch s {
	default:
		err = fmt.Errorf("unknown store type: %v", s)
	case "map":
		if path == "" {
			path = "map.db"
		}
		store, err = kvbench.NewMapStore(path, fsync)
	case "btree":
		if path == "" {
			path = "btree.db"
		}
		store, err = kvbench.NewBTreeStore(path, fsync)
	case "bolt":
		if path == "" {
			path = "bolt.db"
		}
		store, err = kvbench.NewBoltStore(path, fsync)
	case "bbolt":
		if path == "" {
			path = "bbolt.db"
		}
		store, err = kvbench.NewBboltStore(path, fsync)
	case "leveldb":
		if path == "" {
			path = "leveldb.db"
		}
		store, err = kvbench.NewLevelDBStore(path, fsync)
	case "kv":
		log.Warningf("kv store is unstable")
		if path == "" {
			path = "kv.db"
		}
		store, err = kvbench.NewKVStore(path, fsync)
	case "badger":
		if path == "" {
			path = "badger"
		}
		store, err = kvbench.NewBadgerStore(path, fsync)
	case "buntdb":
		if path == "" {
			path = "buntdb.db"
		}
		store, err = kvbench.NewBuntdbStore(path, fsync)
	case "rocksdb":
		if path == "" {
			path = "rocksdb.db"
		}
		store, err = kvbench.NewRocksdbStore(path, fsync)
	}

	return store, path, err
}
