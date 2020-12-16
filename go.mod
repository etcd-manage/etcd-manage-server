module github.com/etcd-manage/etcd-manage-server

go 1.14

replace (
	github.com/coreos/bbolt => github.com/coreos/bbolt v1.3.3
	github.com/etcd-manage/etcd-manage-ui/tpls => ../../../github.com/etcd-manage/etcd-manage-ui/tpls
	google.golang.org/grpc => google.golang.org/grpc v1.26.0 // grpc对etcd依赖问题
)

require (
	github.com/coreos/bbolt v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/etcd-manage/etcd-manage-ui/tpls v0.0.0-00010101000000-000000000000
	github.com/etcd-manage/etcdsdk v0.0.0-20191206100937-45fc0eca65f0
	github.com/gin-gonic/autotls v0.0.0-20200518075542-45033372a9ad
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/google/uuid v1.1.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/jinzhu/gorm v1.9.13
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/naoina/go-stringutil v0.1.0 // indirect
	github.com/naoina/toml v0.1.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/prometheus/client_golang v1.8.0 // indirect
	github.com/shiguanghuxian/etcd-manage v2.0.0-beta+incompatible
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200427203606-3cfed13b9966 // indirect
	go.uber.org/zap v1.15.0
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
	google.golang.org/genproto v0.0.0-20200618031413-b414f8b61790 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)
