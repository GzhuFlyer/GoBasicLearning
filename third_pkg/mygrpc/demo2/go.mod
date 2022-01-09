module demo2

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	google.golang.org/grpc v1.43.0
)

replace github.com/coreos/bbolt v1.3.6 => go.etcd.io/bbolt v1.3.6
