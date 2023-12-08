# grpc-go-service
This is an example of gRPC server implementation of *GreetService* from [THIS](https://github.com/nirdosh17/protorepo/blob/main/greet/greet.proto) public *protobuf* definition.


# Running service locally:
```bash
make run
```

# Testing service endpoint:
- **Opt 1**: Install [Evans gRPC client](https://github.com/ktr0731/evans) and run command: `make grpc-client`
- **Opt 2**: Running client implementation from this repo: https://github.com/nirdosh17/grpc-go-client

## GRPC endpoints:
- Unary: `Greet`
- Client streaming: `LongGreet`
- Server streaming: `GreetManyTimes`
- Bi-directional: `GreetEveryone`
- WithDeadline: `GreetWithDeadline`
