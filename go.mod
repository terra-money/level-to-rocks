module test

go 1.17

// Breaking changes were released with the wrong tag (use v0.6.6 or later).
retract v0.6.5

require (
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
	github.com/tendermint/tm-db v0.6.7
)

require (
	github.com/DataDog/zstd v1.4.1 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cosmos/gorocksdb v1.2.0 // indirect
	github.com/dgraph-io/badger/v2 v2.2007.2 // indirect
	github.com/dgraph-io/ristretto v0.0.3-0.20200630154024-f66de99634de // indirect
	github.com/dgryski/go-farm v0.0.0-20190423205320-6a90982ecee2 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)
