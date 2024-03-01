package function

import (
	"context"
	"io"
	"log"
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// Handle an HTTP Request.
func Handle(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	/*
	 * YOUR CODE HERE
	 *
	 * Try running `go test`.  Add more test as you code in `handle_test.go`.
	 */
	body, _ := io.ReadAll(req.Body)
	event := cloudevents.NewEvent()
	event.SetID("0")
	event.SetType("com.example.event")
	event.SetSource("event-sender")
	event.SetData("text/plain", string(body))

	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	sendCtx := cloudevents.ContextWithTarget(context.Background(), "http://parallel-kn-parallel-kn-channel.flows.svc.cluster.local")
	if result := c.Send(sendCtx, event); cloudevents.IsUndelivered(result) {
		log.Fatalf("failed to send, %v", result)
	} else {
		log.Printf("sent: %v", event)
		res.WriteHeader(http.StatusAccepted)
	}
}
