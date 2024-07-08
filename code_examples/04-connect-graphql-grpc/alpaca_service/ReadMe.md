# Create Graphql Server

1. Initialize go.mod
2. Using gqlgen, initialize with `go get github.com/99designs/gqlgen`
3. Update gQL Schema
4. Regenerate the GQL files `go run github.com/99designs/gqlgen generate`
5. Run Alpaca server `go run main.go`
6. Make CURL request with new Alpaca name `curl -X POST http://localhost:8000/alpaca/ -H "Content-Type: application/json" -d '{"Name": "Fernando"}'`