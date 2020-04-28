module github.com/SleepingNext/product-service

go 1.13

require (
	github.com/SleepingNext/auth-service v0.0.0-20200311050407-6644661a0b56
	github.com/SleepingNext/product-service-cli v0.0.0-20200312032000-f9e80f3b55cf // indirect
	github.com/golang/protobuf v1.3.4
	github.com/grpc-ecosystem/grpc-gateway v1.9.5 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.3.0
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/onsi/ginkgo v1.10.1 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/prometheus/client_golang v1.3.0 // indirect
	golang.org/x/net v0.0.0-20200301022130-244492dfa37a
)

// +heroku goVersion go1.13.8
