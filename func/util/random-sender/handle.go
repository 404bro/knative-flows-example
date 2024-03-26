package function

import (
	"context"
	"fmt"
	"math/rand"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/google/uuid"
)

// Handle an event.
func Handle(ctx context.Context, e event.Event) (*event.Event, error) {
	/*
	 * YOUR CODE HERE
	 *
	 * Try running `go test`.  Add more test as you code in `handle_test.go`.
	 */
	event := cloudevents.NewEvent()
	event.SetID(uuid.New().String())
	event.SetType("com.example.collatz")
	event.SetSource("random-sender")
	num := rand.Int() % 100000
	if rand.Intn(300) < 100 {
		if num%2 == 1 {
			num += 1
		}
	} else {
		if num%2 == 0 {
			num += 1
		}
	}
	event.SetData("text/plain", fmt.Sprintf("%d", num))
	return &event, nil // echo to caller
}

/*
Other supported function signatures:

	Handle()
	Handle() error
	Handle(context.Context)
	Handle(context.Context) error
	Handle(event.Event)
	Handle(event.Event) error
	Handle(context.Context, event.Event)
	Handle(context.Context, event.Event) error
	Handle(event.Event) *event.Event
	Handle(event.Event) (*event.Event, error)
	Handle(context.Context, event.Event) *event.Event
	Handle(context.Context, event.Event) (*event.Event, error)

*/
