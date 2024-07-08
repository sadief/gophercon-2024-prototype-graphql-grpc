module github.com/alpaca_app/04-connect-graphql-grpc/squish_analytics_service

go 1.22.5

require (
	buf.build/gen/go/sadiefr/squishapis/connectrpc/go v1.16.2-20240708160307-7d4cca010576.1
	buf.build/gen/go/sadiefr/squishapis/protocolbuffers/go v1.34.2-20240708160307-7d4cca010576.2
	connectrpc.com/connect v1.16.2
	github.com/google/uuid v1.6.0
	golang.org/x/net v0.23.0
	google.golang.org/protobuf v1.34.2
)

require golang.org/x/text v0.14.0 // indirect
