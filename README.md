# grpc-go-service
This is an example of gRPC server implementation of *GreetService* from [THIS](https://github.com/nirdosh17/protorepo/blob/main/greet/greet.proto) public *protobuf* definition.

## GRPC endpoints:
- Unary: `Greet`
- Client streaming: `LongGreet`
- Server streaming: `GreetManyTimes`
- Bi-directional: `GreetEveryone`
- WithDeadline: `GreetWithDeadline`
