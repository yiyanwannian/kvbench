#!/bin/sh

COUNT=10000
SIZE=256

STORES=("badger" "bbolt" "bolt" "leveldb" "kv" "buntdb" "rocksdb" "btree" "btree/memory" "map" "map/memory")


`rm  -f .*`
`rm  -fr *.db`
`rm -fr badger`

echo "=========== test nofsync ==========="
for i in "${STORES[@]}"
do
	./main -n ${COUNT} -size ${SIZE} -s "$i" >> benchmarks/test.log 2>&1
done

`rm  -f .*`
`rm  -fr *.db`
`rm -fr badger`

echo ""
echo "=========== test fsync ==========="

for i in "${STORES[@]}"
do
	./main -n ${COUNT} -size ${SIZE} -s "$i" -fsync >> benchmarks/test.log 2>&1
done

`rm  -f .*`
`rm  -fr *.db`
`rm -fr badger`