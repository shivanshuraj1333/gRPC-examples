# gRPC-examples
this repository containes basic to advance exmples of gRPC, covering Unary, Server Streaming, Client Streaming, Bi Directional Streaming
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
