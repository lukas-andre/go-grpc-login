# GRPC + Go + Postgres <3


> this project it is verbose on purpose 

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

## TODO 
- [ ] add a explication of the PoC arquitecture of the finish project
- [ ] follow this flow: 
    - [ ] protofiles with buf and Makefile
    - [ ] start grpc server
    - [ ] configuration
    - [ ] connect to postgres with database config
    - [ ] create user model and migrate to database
    - [ ] create basic project structure grpc_service -> domain_service -> repository -> DAO
    - [ ] use wire to manage dependencies
    - [ ] implement and explain create user explanation
    - [ ] use evans CLI to test the Grpc sever
    - [ ] implement and explain login & token handler
    - [ ] implementent logging interceptor
    - [ ] implementent authorization interceptor 
    - [ ] implementent globalService scope in go
    - [ ] use globalService to validate token
    - [ ] how to register public and private grpc methods
    - [ ]  implementent methodService to validate token
- [ ] add HTTP server and swagger generation with protofile annotations and grpc-gateway




