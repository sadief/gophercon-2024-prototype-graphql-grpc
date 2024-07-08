# Connect GraphQL + gRPC

1. Register API with Buf Schema Registry
2. `buf push`
3. Insert client code to GraphQL resolver
4. `go run alpaca_service/server.go`
5. `go run squish_analytics_service/server/main.go`
6. Go to `http://localhost:8080/` for the GraphQL Playground and query away