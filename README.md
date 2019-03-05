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

| |badger|bbolt|bolt|leveldb|kv|buntdb|rocksdb|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|
|del|113590|4054|3523|219681|6379|188057|670461|461692|908052|797977|1405268|
|set|97441|11922|12414|130572|2267|99370|483053|203632|555432|243506|1120192|
|get|474636|428163|467606|338100|17291|3051031|2433409|4966106|5485222|6740588|6555700|
|getmixed|333540|127328|241168|344907|12359|359620|1976448|1275432|1519957|2240288|2597800|
|setmixed|1619|9469|7842|17220|309|6001|107852|17673|18513|18249|20234|

**time (latency)**

| |badger|bbolt|bolt|leveldb|kv|buntdb|rocksdb|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|
|set|97441|11922|12414|130572|2267|99370|483053|203632|555432|243506|1120192|22|
|getmixed|333540|127328|241168|344907|12359|359620|1976448|1275432|1519957|2240288|2597800|9|
|get|474636|428163|467606|338100|17291|3051031|2433409|4966106|5485222|6740588|6555700|3|
|setmixed|1619|9469|7842|17220|309|6001|107852|17673|18513|18249|20234|49421|
|del|113590|4054|3523|219681|6379|188057|670461|461692|908052|797977|1405268|17|

### fsync

**throughputs**

| |badger|bbolt|bolt|leveldb|kv|buntdb|rocksdb|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|
|setmixed|64|60|31|54|124|66|-1|67|20827|66|23213|
|del|1309|81|73|1490|7287|162|1386|164|737082|163|802890|
|set|1071|67|70|1466|5084|66|1321|66|435274|67|660196|
|get|401010|276105|323697|510568|5664|1728309|1529519|1973164|2735229|2673796|3229974|
|getmixed|4201|300291|315338|541976|4938|3082|1961553|2144|890075|2816|1406865|

**time (latency)**

| |badger|bbolt|bolt|leveldb|kv|buntdb|rocksdb|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|
|set|1071|67|70|1466|5084|66|1321|66|435274|67|660196|37|
|get|401010|276105|323697|510568|5664|1728309|1529519|1973164|2735229|2673796|3229974|7|
|setmixed|64|60|31|54|124|66|-1|67|20827|66|23213|43080|
|getmixed|4201|300291|315338|541976|4938|3082|1961553|2144|890075|2816|1406865|17|
|del|1309|81|73|1490|7287|162|1386|164|737082|163|802890|31|


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