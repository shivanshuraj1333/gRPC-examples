# gRPC-examples
this repository containes basic to advance exmples of gRPC, covering Unary, Server Streaming, Client Streaming, Bi Directional Streaming
all the gRPC examples are in go
```
|── BiDirectionalStreaming
│   ├── GreetEveryone
│   │   ├── client
│   │   │   └── client.go
│   │   ├── proto
│   │   │   ├── greet.pb.go
│   │   │   ├── greet.proto
│   │   │   └── greet_grpc.pb.go
│   │   └── server
│   │       └── server.go
│   └── findMaximumApi
│       ├── client
│       │   └── client.go
│       ├── proto
│       │   ├── max.pb.go
│       │   ├── max.proto
│       │   └── max_grpc.pb.go
│       └── server
│           └── server.go
├── ClientStreaming
│   ├── computeAverageApi
│   │   ├── client
│   │   │   └── client.go
│   │   ├── proto
│   │   │   ├── compute.pb.go
│   │   │   ├── compute.proto
│   │   │   └── compute_grpc.pb.go
│   │   └── server
│   │       └── server.go
│   └── longGreet
│       ├── client
│       │   └── client.go
│       ├── proto
│       │   ├── longGreet.pb.go
│       │   ├── longGreet.proto
│       │   └── longGreet_grpc.pb.go
│       └── server
│           └── server.go
├── ServerStreaming
│   ├── primeNumberDecomposition
│   │   ├── client
│   │   │   └── client.go
│   │   ├── proto
│   │   │   ├── prime.pb.go
│   │   │   ├── prime.proto
│   │   │   └── prime_grpc.pb.go
│   │   └── server
│   │       └── server.go
│   └── videoStreaming
│       ├── client
│       │   └── client.go
│       ├── proto
│       │   ├── video.pb.go
│       │   ├── video.proto
│       │   └── video_grpc.pb.go
│       └── server
│           └── server.go
├── Unary
│   ├── CalculatorService
│   │   ├── client
│   │   │   └── client.go
│   │   ├── proto
│   │   │   ├── calc.pb.go
│   │   │   ├── calc.proto
│   │   │   └── calc_grpc.pb.go
│   │   └── server
│   │       └── server.go
│   ├── GreetService
│   │   ├── client
│   │   │   └── client.go
│   │   ├── greetpb
│   │   │   ├── greet.pb.go
│   │   │   ├── greet.proto
│   │   │   └── greet_grpc.pb.go
│   │   └── server
│   │       └── server.go
│   └── UserManagementService
│       ├── usermgmt
│       │   ├── usermgmt.pb.go
│       │   ├── usermgmt.proto
│       │   └── usermgmt_grpc.pb.go
│       ├── usermgmt_client
│       │   └── usermgmt_client.go
│       └── usermgmt_server
│           └── usermgmt_server.go
├── go.mod
└── go.sum
```
1. typed of APIs:
    1. Unary: 1 req - 1 res
    2. Server Streaming: 1 req - n res
    3. Client Streaming: n req - 1 res
    4. Bi Directional Streaming: n req - n res
2. Security as a first class citizen
3. Easy and understandable status codes https://grpc.github.io/grpc/core/md_doc_statuscodes.html
