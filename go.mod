module github.com/joshcarp/grpcexample

go 1.14

require (
	github.com/googleapis/gax-go/v2 v2.1.1
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)

replace google.golang.org/grpc v1.30.0 => google.golang.org/grpc v1.29.0
