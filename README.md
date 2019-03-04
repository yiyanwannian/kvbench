# KVBench

cloned from [tidwall/kvbench](https://github.com/tidwall/kvbench)

KVBench is a Redis server clone backed by a few different Go databases. 

It's intended to be used with the `redis-benchmark` command to test the performance of various Go databases.
It has support for redis pipelining.

this cloned version adds more kv databases and automatic scripts.

Features:

- Databases
  - [badger](https://github.com/dgraph-io/badger)
  - [BboltDB](https://github.com/etcd-io/bbolt)
  - [BoltDB](https://github.com/boltdb/bolt)
  - [buntdb](https://github.com/tidwall/buntdb)
  - [LevelDB](https://github.com/syndtr/goleveldb)
  - [cznic/kv](https://github.com/cznic/kv)
  - [rocksdb](https://github.com/tecbot/gorocksdb)
  - map (in-memory) with [AOF persistence](https://redis.io/topics/persistence)
  - btree (in-memory) with [AOF persistence](https://redis.io/topics/persistence)
- Option to disable fsync
- Compatible with Redis clients

## Benchmarks

Test concurrency cases:

- set: use 40 goroutines to write
- get: use 40 goroutines to read
- mixed: use one goroutine to write and use 40 goroutines to read
- del: use 40 goroutines to delete

Use 10000 items to test. Size of value of items is 256 bytes and size of key of items is 9 bytes.

### fsync disabled

#### throughouts

![set](cmd/cli/benchmarks/nofsync-set-throughputs.png)
![get](cmd/cli/benchmarks/nofsync-get-throughputs.png)
![mixed](cmd/cli/benchmarks/nofsync-mixed-throughputs.png)
![del](cmd/cli/benchmarks/nofsync-del-throughputs.png)

#### time

![set](cmd/cli/benchmarks/nofsync-set-time.png)
![get](cmd/cli/benchmarks/nofsync-get-time.png)
![mixed](cmd/cli/benchmarks/nofsync-mixed-time.png)
![del](cmd/cli/benchmarks/nofsync-del-time.png)

### fsync enabled

#### throughouts

![set](cmd/cli/benchmarks/fsync-set-throughputs.png)
![get](cmd/cli/benchmarks/fsync-get-throughputs.png)
![mixed](cmd/cli/benchmarks/fsync-mixed-throughputs.png)
![del](cmd/cli/benchmarks/fsync-del-throughputs.png)

#### time

![set](cmd/cli/benchmarks/fsync-set-time.png)
![get](cmd/cli/benchmarks/fsync-get-time.png)
![mixed](cmd/cli/benchmarks/fsync-mixed-time.png)
![del](cmd/cli/benchmarks/fsync-del-time.png)