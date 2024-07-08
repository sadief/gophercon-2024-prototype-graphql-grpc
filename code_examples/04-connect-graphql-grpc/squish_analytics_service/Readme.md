# Create gRPC Server

1. Install Buf CLI
2. Create `buf.yaml` with `buf config init`
3. Define your path to the `.proto` file:
    `modules
        path: proto`
4. Build with `buf build`
5. Create `buf.gen.yaml`  to use Go and Connect-Go plugins
6. Edit `.proto schema`
7. Run `buf generate`
8. Start server with `go run server/main.go`
9. Test with buf curl command `buf curl \
  --schema . \
  --data '{"alpaca_id": "1234", "squish_rating": "10"}' \
  http://localhost:8080/squish_analytics.v1.SquishAnalyticsService/CreateSquishRating`