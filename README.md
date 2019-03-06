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
  - [pebble](https://github.com/petermattis/pebble)
  - map (in-memory) with [AOF persistence](https://redis.io/topics/persistence)
  - btree (in-memory) with [AOF persistence](https://redis.io/topics/persistence)
- Option to disable fsync
- Compatible with Redis clients


## HDD benchmark

### nofsync

**throughputs**

| |badger|bbolt|bolt|leveldb|kv|buntdb|pebble|rocksdb|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|--|
|set|97309|11934|12516|134054|2535|106322|481964|495220|195407|570192|259629|1073199|
|get|535912|467077|538754|356332|17408|1424205|2356922|2538238|5107017|5319205|6658498|6584665|
|setmixed|1755|9680|8073|18019|318|10123|105829|101276|16574|20383|19298|19274|
|del|102790|4161|3367|217777|6409|213214|675658|681636|465202|1017582|842456|1347452|
|getmixed|317882|127199|234670|336547|12732|581805|2013081|1963066|1235444|1422469|2257621|2308104|

**time (latency)**

| |badger|bbolt|bolt|leveldb|kv|buntdb|pebble|rocksdb|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|--|
|setmixed|569585|103295|123867|55494|3140349|98775|9449|9873|60332|49058|51818|51880|
|del|243|6007|7424|114|3900|117|37|36|53|24|29|18|
|set|256|2094|1997|186|9859|235|51|50|127|43|96|23|
|get|46|53|46|70|1436|17|10|9|4|4|3|3|
|getmixed|78|196|106|74|1963|42|12|12|20|17|11|10|

### fsync

**throughputs**

| |badger|bbolt|bolt|leveldb|kv|buntdb|pebble|rocksdb|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|--|
|getmixed|4478|326285|286623|569800|4886|2911|1923446|2349072|1967|994233|2960|1281065|
|del|1171|82|62|1060|7317|162|1335|1306|164|842247|164|847816|
|get|429793|279845|272294|559221|7859|1811922|1522997|1688903|3181673|2851439|3055300|2849814|
|setmixed|64|32|57|56|123|65|-1|-1|66|19785|67|20625|
|set|1099|68|67|2275|5131|66|1347|1359|67|429553|67|631432|


**time (latency)**

| |badger|bbolt|bolt|leveldb|kv|buntdb|pebble|rocksdb|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|--|
|set|22745|363186|371857|10986|4872|373860|18551|18393|370929|58|371504|39|
|setmixed|15505773|30648176|17444811|17550529|8088880|15330992|-1|-1|15037276|50546|14880390|48490|
|get|58|89|91|44|3180|13|16|14|7|8|8|8|
|getmixed|5582|76|87|43|5116|8585|12|10|12706|25|8444|19|
|del|21340|302808|396976|23582|3416|153810|18723|19140|152106|29|151896|29|

## SDD benchmark

### nofsync

**throughputs**

| |badger|bbolt|bolt|leveldb|kv|buntdb|pebble|rocksdb|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|--|
|getmixed|396374|72794|226852|273125|13685|273452|1926559|1925973|659800|941788|1464767|1786358|
|del|82684|3424|3054|148731|7008|191465|629728|633072|380794|813068|597367|1137109|
|get|438464|412654|454839|405765|17011|2961620|1689377|1617706|4041922|3665312|4420964|4015257|
|setmixed|4410|8283|9241|17141|685|13497|138785|148465|30585|38453|43119|47488|
|set|74325|12597|13396|123448|3613|102831|447157|450753|179979|597334|207175|866142|

**time (latency)**

| |badger|bbolt|bolt|leveldb|kv|buntdb|pebble|rocksdb|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|--|
|set|672|3969|3732|405|13836|486|111|110|277|83|241|57|
|get|114|121|109|123|2939|16|29|30|12|13|11|12|
|setmixed|226754|120726|108204|58336|1459733|74090|7205|6735|32695|26005|23191|21057|
|getmixed|126|686|220|183|3653|182|25|25|75|53|34|27|
|del|604|14599|16371|336|7134|261|79|78|131|61|83|43|

### fsync

**throughputs**

| |badger|bbolt|bolt|leveldb|kv|buntdb|pebble|rocksdb|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|--|
|get|433933|264725|274340|584829|5637|1886080|2386065|1901863|3041362|2814522|3636363|3424657|
|getmixed|131276|204951|241943|622393|4672|101736|3575259|2841716|103878|681942|118910|1353729|
|del|40289|6973|5976|73800|7291|15474|67268|59760|18274|802632|17738|830909|
|set|33149|5519|5609|52800|4868|7593|53698|49733|7692|517625|7635|829187|
|setmixed|4358|4508|4427|5663|237|6134|3932|3694|6388|28777|6932|51712|


**time (latency)**

| |badger|bbolt|bolt|leveldb|kv|buntdb|pebble|rocksdb|btree|btree/memory|map|map/memory|
|--|--|--|--|--|--|--|--|--|--|--|--|--|
|set|1508|9059|8913|946|10269|6584|931|1005|6500|96|6548|60|
|getmixed|380|243|206|80|10700|491|13|17|481|73|420|36|
|del|1241|7169|8365|677|6857|3231|743|836|2736|62|2818|60|
|get|115|188|182|85|8869|26|20|26|16|17|13|14|
|setmixed|229444|221785|225861|176560|4204335|163008|254310|270760|156530|34749|144248|19338|

