module github.com/tomoyane/grant-n-z

go 1.13

require (
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.7.4
	github.com/jinzhu/gorm v1.9.12
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/stretchr/testify v1.4.0
	go.etcd.io/etcd v3.3.18+incompatible
	go.uber.org/zap v1.14.0 // indirect
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
	google.golang.org/grpc v1.27.1 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/yaml.v2 v2.2.8
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
