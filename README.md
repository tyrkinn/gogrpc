# Installing dependencies
Run this command in root directory
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
go mod tidy
```

# Generating proto-files
## Linux/MacOS
```bash
sh gen_protoc.sh
```
## Windows
```bash
gen_protoc.bat
```

# Running
## Server
```bash
go run server/main.go
```

## Client
```bash
go run client/main.go
```