#!/bin/sh

`rm -f benchmarks/*.csv`

STORES=("badger" "bbolt" "bolt" "leveldb" "kv" "buntdb" "rocksdb" "btree" "btree/memory" "map" "map/memory")

echo "name,set,get,set-mixed,get-mixed,del" >> benchmarks/nofsync_throughputs.csv
echo "name,set,get,set-mixed,get-mixed,del" >> benchmarks/nofsync_time.csv
echo "name,set,get,set-mixed,get-mixed,del" >> benchmarks/fsync_throughputs.csv
echo "name,set,get,set-mixed,get-mixed,del" >> benchmarks/fsync_time.csv

for i in "${STORES[@]}"
do

    data=`grep -e ^${i}/nofsync  benchmarks/test.log|awk '{print $4}'|xargs| tr ' ' ','`
    echo "${i}/nofsync,${data}" >> benchmarks/nofsync_throughputs.csv
    data=`grep -e ^${i}/nofsync  benchmarks/test.log|awk '{print $7}'|xargs| tr ' ' ','`
    echo "${i}/nofsync,${data}" >> benchmarks/nofsync_time.csv

    data=`grep -e ^${i}/fsync  benchmarks/test.log|awk '{print $4}'|xargs| tr ' ' ','`
    echo "${i}/fsync,${data}" >> benchmarks/fsync_throughputs.csv
    data=`grep -e ^${i}/nofsync  benchmarks/test.log|awk '{print $7}'|xargs| tr ' ' ','`
    echo "${i}/fsync,${data}" >> benchmarks/fsync_time.csv

	# grep "${i}/nofsync set rate" benchmarks/test.log | awk '{print $1","$4}' >> benchmarks/nofsync_set_throughputs.csv
    # grep "${i}/nofsync set rate" benchmarks/test.log | awk '{print $1","$7}' >> benchmarks/nofsync_set_time.csv

    # grep "${i}/fsync set rate" benchmarks/test.log | awk '{print $1","$4}' >> benchmarks/fsync_throughputs.csv
    # grep "${i}/fsync set rate" benchmarks/test.log | awk '{print $1","$7}' >> benchmarks/fsync_time.csv
done
