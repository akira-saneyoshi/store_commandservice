# store_commandservice

```zsh
go get -u github.com/onsi/ginkgo/v2@latest
go get -u github.com/onsi/gomega@latest

go get -u github.com/google/uuid

go install github.com/onsi/ginkgo/v2/ginkgo
```

```zsh
go get -u github.com/go-sql-driver/mysql
go get -u github.com/volatiletech/sqlboiler/v4
go get -u github.com/volatiletech/null/v
go get -u github.com/BurntSushi/toml

go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
```

```zsh
go get -u github.com/akira-saneyoshi/store_pb@v1.0.0
```

```zsh
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

```zsh
go build -o commandservice main.go
```

```zsh
grpcurl -plaintext localhost:8082 list

grpcurl -plaintext localhost:8082 list proto.CategoryCommand

grpcurl -plaintext localhost:8082 describe proto.CategoryCommand

grpcurl -plaintext -d '{"crud":"1", "name":"食料品"}' localhost:8082 proto.CategoryCommand.Create
```

```zsh
apk update

apk add curl

wget -O mkcert https://github.com/FiloSottile/mkcert/releases/download/v1.4.4/mkcert-v1.4.4-linux-amd64

chmod +x mkcert

mv mkcert /usr/local/bin/

mkcert -install

cd commands/presen/prepare/

mkcert commandservice
```
