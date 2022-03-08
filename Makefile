build: go.sum
	go build -tags rocksdb -o level-to-rocks ./main.go

install: ./level-to-rocks
	cp level-to-rocks ~/bin
