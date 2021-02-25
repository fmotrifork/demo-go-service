# Building the demo-go-service

## Using Make

The following commands are avaible through `make`:

```bash
Management commands for demo-go-service:

Usage:
    make build           Compile the project.
    make get-deps        runs dep ensure, mostly used for ci.
    make clean           Clean the directory tree.
```

### Build and run example

```bash
$ make build
building demo-go-service 0.1.0
GOPATH=/home/mogensen/go
go build -ldflags "-X github.com/fmotrifork/demo-go-service/version.GitCommit=6ad1a13c2213a45b3c5b591723fce45294d36edb+CHANGES -X github.com/fmotrifork/demo-go-service/version.BuildDate=2021-02-25-15:18:55" -o bin/demo-go-service

$ ./bin/demo-go-serive server
...
2021/02/25 15:20:35 Starting webserver on: localhost:3333
...
```

