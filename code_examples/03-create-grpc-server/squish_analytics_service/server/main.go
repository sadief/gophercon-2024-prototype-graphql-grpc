package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	connect "connectrpc.com/connect"
	squish_analytics_v1 "github.com/alpaca_app/03-create-grpc-server/gen/squish_analytics/v1"
	squish_analytics_v1connect "github.com/alpaca_app/03-create-grpc-server/gen/squish_analytics/v1/squish_analyticsv1connect"
	"github.com/google/uuid"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const address = "localhost:8000"

func main() {
	mux := http.NewServeMux()
	path, handler := squish_analytics_v1connect.NewSquishAnalyticsServiceHandler(&squishAnalyticsServiceServer{})
	mux.Handle(path, handler)
	fmt.Println("... Listening on", address)
	http.ListenAndServe(
		address,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

// squishAnalyticsServiceServer implements the SquishAnalyticsService API.
type squishAnalyticsServiceServer struct {
	squish_analytics_v1connect.UnimplementedSquishAnalyticsServiceHandler
}

// CreateSquishRating adds a new rating entry for an Alpaca
func (s *squishAnalyticsServiceServer) CreateSquishRating(
	ctx context.Context,
	req *connect.Request[squish_analytics_v1.CreateSquishRatingRequest],
) (*connect.Response[squish_analytics_v1.CreateSquishRatingResponse], error) {
	squishRating := req.Msg.GetSquishRating()
	log.Printf("Received request for new Squish Rating: %v", squishRating)

	newRating := &squish_analytics_v1.CreateSquishRatingResponse{
		SquishRatingId: uuid.New().String(),
	}
	// Save to DB

	return connect.NewResponse(newRating), nil
}
