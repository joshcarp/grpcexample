module github.com/joshcarp/grpcexample

go 1.14

require (
	github.com/googleapis/gax-go/v2 v2.7.0
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)

replace google.golang.org/grpc v1.30.0 => google.golang.org/grpc v1.29.0
