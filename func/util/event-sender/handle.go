package function

import (
	"context"
	"io"
	"log"
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
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
	event.SetID(uuid.New().String())
	event.SetType("com.example.hello")
	event.SetSource("event-sender")
	event.SetData("text/plain", string(body))

	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	sendCtx := cloudevents.ContextWithTarget(context.Background(), "http://broker-ingress.knative-eventing.svc.cluster.local/flows-example/default")
	if result := c.Send(sendCtx, event); cloudevents.IsUndelivered(result) {
		log.Fatalf("failed to send, %v", result)
		res.WriteHeader(http.StatusInternalServerError)
	} else {
		log.Printf("sent: %v", event)
		res.WriteHeader(http.StatusAccepted)
	}
}
