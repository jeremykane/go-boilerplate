# GO-BACKEND

**Requirement**
1. [Golang 1.20](https://go.dev/doc/install)
2. [Docker](https://www.docker.com/)
3. [Pre-commit](https://pre-commit.com/)
4. Run `pre-commit install` in your project folder and `go install golang.org/x/tools/cmd/goimports@latest`
5. Adjust your zshrc path for GOPATH and go env

**How to run locally**

Run dependencies using docker
```
make start-components
```
If needed, adjust local configuration (db name, user, etc) in `/config/config.go`

Download libraries
```
go mod tidy && go mod vendor
```

Build & run API
```
go build --race -o go-backend ./cmd/rest
./go-backend
```
or simply use `Make` command
```
make start-api
```

Build & run Worker
```
make start-worker
```

**How to open swagger docs**

Open in the browser `localhost:9000/swagger/index.html`


**How to write unit tests**
1. Install gomock
```
make install-gomock
```
2. Run the `./scripts/mockgen.sh` (note: every changes in interface should run this again)
3. [Example](https://github.com/jeremykane/go-boilerplate/blob/8ea507cedba76b3fa283fcfdd5da9696d7b51d08/internal/config/config_test.go#L10)