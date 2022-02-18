# GRPC + Go + Postgres <3

First create a go module with this command 
```sh
go mod init
```

First install protobuf compailer 

Linux, using apt or apt-get, fox example
```sh
$ apt install -y protobuf-compiler
$ protoc --version  # Ensure compiler version is 3+
```

MacOS, using Homebrew:
```sh
$ brew install protobuf
$ protoc --version  # Ensure compiler version is 3+
```
[More instalation option](https://grpc.io/docs/protoc-installation/)

then install protoc-gen-go with command:

```sh 
brew install protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

then when protoc-gen-go-grpc is installed you can generate code with command:

```sh 
protoc --go_out=. \
    --go-grpc_out=. \
    protos/login.proto
```

for simplicity make a shell script to generate code:

```sh
echo "protoc --go_out=. \
    --go-grpc_out=. \
    protos/login.proto" > generate_protos.sh
```

en execute the script: 

```sh 
sh ./generate_protos.sh
```

that will generate code in loginpb directory (you can change the name if you want in login.proto replacing the name of the go package)


create grpc_login database in postgres
```sql
CREATE DATABASE grpc_login;
```



