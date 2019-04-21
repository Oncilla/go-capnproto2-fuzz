# Fuzzing go-capnproto2

During fuzzing [go-capnproto2](github.com/capnproto/go-capnproto2)  with [go-fuzz](https://github.com/dvyukov/go-fuzz) an error was found when parsing invalid composite lists.

When creating a `capnp.List` from a `capnp.Ptr`, the `uint32` value `Ptr.lenOrCap` is cast to a `int32`,
possibly resulting in a negative value. During the validity check, the non-negativity of the list length is not considered, resulting in a panic in the pogs library 
[extract.go:L261](https://github.com/capnproto/go-capnproto2/blob/e1ae1f982d9908a41db464f02861a850a0880a5a/pogs/extract.go#L261)

Running the panicking tests on your machine:
````
make test
````

To run the fuzzer locally:
````
make fuzz-all
````
