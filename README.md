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
|get|474636|428163|467606|338100|17291|3051031|2433409|4966106|5485222|6740588|6555700|
|setmixed|1619|9469|7842|17220|309|6001|107852|17673|18513|18249|20234|
|getmixed|333540|127328|241168|344907|12359|359620|1976448|1275432|1519957|2240288|2597800|
|del|113590|4054|3523|219681|6379|188057|670461|461692|908052|797977|1405268|
|set|97441|11922|12414|130572|2267|99370|483053|203632|555432|243506|1120192|

**time (latency)**

| |badger|bbolt|bolt|leveldb|kv|buntdb|rocksdb|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|
|setmixed|617409|105605|127517|58071|3235598|166629|9271|56581|54015|54796|49421|
|del|220|6165|7096|113|3918|132|37|54|27|31|17|
|set|256|2096|2013|191|11023|251|51|122|45|102|22|
|get|52|58|53|73|1445|8|10|5|4|3|3|
|getmixed|74|196|103|72|2022|69|12|19|16|11|9|

### fsync

**throughputs**

| |badger|bbolt|bolt|leveldb|kv|buntdb|rocksdb|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|
|set|1071|67|70|1466|5084|66|1321|66|435274|67|660196|
|del|1309|81|73|1490|7287|162|1386|164|737082|163|802890|
|get|401010|276105|323697|510568|5664|1728309|1529519|1973164|2735229|2673796|3229974|
|setmixed|64|60|31|54|124|66|-1|67|20827|66|23213|
|getmixed|4201|300291|315338|541976|4938|3082|1961553|2144|890075|2816|1406865|

**time (latency)**

| |badger|bbolt|bolt|leveldb|kv|buntdb|rocksdb|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|
|getmixed|5950|83|79|46|5062|8110|12|11655|28|8874|17|
|del|19089|307542|337844|16776|3430|154249|18037|151850|33|152854|31|
|setmixed|15557829|16650544|31712503|18451317|8035914|14950086|-1|14848393|48016|14978524|43080|
|set|23325|373100|352715|17043|4917|373841|18922|373197|57|372896|37|
|get|62|90|77|48|4413|14|16|12|9|9|7|

