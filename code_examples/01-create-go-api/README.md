# Create Go API

1. Create Alpaca HTTP request 
2. Run Alpaca server `go run main.go`
3. Make CURL request with new Alpaca name `curl http://localhost:8080/alpaca/ \
-d '{"Name": "Fernando"}'`