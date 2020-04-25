module github.com/echo-marche/hack-tech-tips-api

go 1.13

require (
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.4.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/jinzhu/gorm v1.9.12
	github.com/mwitkow/go-proto-validators v0.3.0
	golang.org/x/crypto v0.0.0-20200423211502-4bdfaf469ed5 // indirect
	golang.org/x/net v0.0.0-20200421231249-e086a090c8fd // indirect
	google.golang.org/genproto v0.0.0-20200424135956-bca184e23272
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.21.0
	gopkg.in/yaml.v2 v2.2.3 // indirect
)

// for realize
replace gopkg.in/urfave/cli.v2 => github.com/urfave/cli/v2 v2.2.0
