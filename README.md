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


## Latest Test Result

### nofsync

**throughputs**

|   |badger|bbolt|bolt|leveldb|kv|buntdb|rocksdb|btree|btree/memory|map|map/memory|
|---|------|-----|----|-------|--|------|-------|-----|------------|---|--------|
|set|98239|12005|12307|130255|2388|97386|488594|196261|559014|254365|1180283|
|set|532633|436289|535859|351586|17834|3494389|2531786|3169883|5325324|6808325|6700930|
|mixed set|1644|8509|7882|17743|290|5075|105005|17100|19085|20020|18279|
|mixed get|324082|116445|230899|335067|11622|299941|2021079|1211994|1487325|2436765|2739350|
|del|104381|4096|3367|215534|6401|198273|701337|446954|983082|753055|1577595|


** time (latency) **

|   |badger|bbolt|bolt|leveldb|kv|buntdb|rocksdb|btree|btree/memory|map|map/memory|
|---|------|-----|----|-------|--|------|-------|-----|------------|---|--------|
|set|254|2082|2031|191|10467|256|51|127|44|98|21|
|set|46|57|46|71|1401|7|9|7|4|3|3|
|mixed set|||||3|||||||
|mixed get|77|214|108|74|2150|83|12|20|16|10|9|
|del|239|6103|7423|115|3905|126|35|55|25|33|15|



## Benchmarks Charts

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