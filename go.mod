module github.com/fdymylja/tmos

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/dgraph-io/badger/v3 v3.2011.1
	github.com/ghodss/yaml v1.0.0
	github.com/gogo/protobuf v1.3.3
	github.com/scylladb/go-set v1.0.2
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.9
	golang.org/x/net v0.0.0-20201209123823-ac852fbbde11 // indirect
	golang.org/x/sys v0.0.0-20201211090839-8ad439b19e0f // indirect
	golang.org/x/text v0.3.4 // indirect
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c
	k8s.io/klog/v2 v2.4.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
