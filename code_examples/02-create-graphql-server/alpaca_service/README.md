# Create Graphql Server

1. Initialize go.mod
2. Using gqlgen, initialize with `go get github.com/99designs/gqlgen`
3. Initialize project layout with `go run github.com/99designs/gqlgen init`
4. Update gQL Schema
5. Update schema.resolvers.go
6. Regenerate the GQL files `go run github.com/99designs/gqlgen generate`
7. Run Alpaca server `go run main.go`