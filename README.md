# Level to Rocks
## What it does
Converts LevelDB database into RocksDB

## Prerequisites
* Go >= 1.17.x
* Existing LevelDB database
* RocksDB 6.27.x installation

## How to build
```sh
$ make build
```

## How to use
1. Create a directory for RocksDB

```sh
$ mkdir -p xyz; cd xyz
```

2. Execute the program

> Usage: level-to-rocks \<db name> \<directory>

For example if you have `mantlemint.db` in directory `~/zyx`

```sh
$ level-to-rocks mantlemint ~/zyx
```
