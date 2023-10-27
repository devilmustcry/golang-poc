# gRPC

### Prerequisite

- Protoc (Protobuf compiler) https://grpc.io/docs/protoc-installation/
- Golang

For workshop please checkout to branch grpc-workshop

To generate protobuf file use command

```shell
protoc --go_out=${PATH_TO_OUTPUT} --go_opt=paths=source_relative \
--go-grpc_out=${PATH_TO_OUTPUT} --go-grpc_opt=paths=source_relative \
${PATH_TO_PROTOBUF_FILE}
```

You can also write this in MakeFile to make it easier to generate code