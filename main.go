package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	tmdb "github.com/tendermint/tm-db"
)

func bz(s string) []byte {
	return []byte(s)
}

func main() {
	if len(os.Args) < 3 {
		panic("Usage: level-to-rocks <dir> <name>")
	}

	o := &opt.Options{
		// The default value is nil
		Filter: filter.NewBloomFilter(10),
		// Use 1 GiB instead of default 8 MiB
		BlockCacheCapacity: opt.GiB,
		// Use 64 MiB instead of default 4 MiB
		WriteBuffer:                           64 * opt.MiB,
		CompactionTableSize:                   8 * opt.MiB,
		CompactionTotalSize:                   40 * opt.MiB,
		CompactionTotalSizeMultiplierPerLevel: []float64{1, 1, 10, 100, 1000, 10000, 100000},
		// This option is the key for the speed
		DisableSeeksCompaction: true,
	}

	db, err := tmdb.NewGoLevelDBWithOpts(os.Args[2], os.Args[1], o)

	if err != nil {
		panic(err)
	}

	newDb, newErr := tmdb.NewRocksDB(os.Args[2], ".")

	if newErr != nil {
		panic(newErr)
	}

	itr, itrErr := db.Iterator(nil, nil)

	if itrErr != nil {
		panic(itrErr)
	}

	offset := 0

	for ; itr.Valid(); itr.Next() {
		key := itr.Key()
		value := itr.Value()

		newDb.Set(key, value)

		offset++

		if offset%10000 == 0 {
			fmt.Println(offset)
			runtime.GC() // Force GC
		}
	}

	itr.Close()
	newDb.Close()
	db.Close()
}
