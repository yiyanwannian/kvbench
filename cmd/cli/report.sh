#!/bin/sh

STORES=("badger" "bbolt" "bolt" "leveldb" "kv" "buntdb" "rocksdb" "btree" "btree/memory" "map" "map/memory")

for i in "${STORES[@]}"
do
	grep "${i}/nofsync set rate" benchmarks/test.log | awk '{print $1","$4}' >> benchmarks/nofsync_throughputs.csv
    grep "${i}/nofsync set rate" benchmarks/test.log | awk '{print $1","$7}' >> benchmarks/nofsync_time.csv

    grep "${i}/fsync set rate" benchmarks/test.log | awk '{print $1","$4}' >> benchmarks/fsync_throughputs.csv
    grep "${i}/fsync set rate" benchmarks/test.log | awk '{print $1","$7}' >> benchmarks/fsync_time.csv
done
