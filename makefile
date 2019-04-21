test: build-proto
	go test -run Panic

build-proto: proto
	capnp compile -I ~/go/src/zombiezen.com/go/capnproto2/std -ogo proto/*.capnp

fuzz-all: fuzz-build fuzz-run

fuzz-run:
	go-fuzz -bin=./bin/list-fuzz.zip -workdir=workdir/list

fuzz-build:
	go-fuzz-build -o bin/list-fuzz.zip ./
