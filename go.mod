module github.com/tomoyane/grant-n-z

go 1.13

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/coreos/etcd v3.3.20+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.7.4
	github.com/jinzhu/gorm v1.9.12
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-sqlite3 v2.0.1+incompatible
	go.etcd.io/etcd v3.3.20+incompatible
	go.uber.org/zap v1.14.1 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	google.golang.org/genproto v0.0.0-20220310185008-1973136f34c6 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.5
