using Go = import "/go.capnp";
@0xab27b8dad5a14d75;
$Go.package("proto");
$Go.import("github.com/Oncilla/go-capnproto2-fuzz/proto");

struct L {
	entries @0 :List(Entry);
}

struct Entry {
	data @0 :Data;
}
