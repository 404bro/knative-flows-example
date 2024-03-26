package function

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/cloudevents/sdk-go/v2/event"
)

// Handle an event.
func Handle(ctx context.Context, e event.Event) (*event.Event, error) {
	/*
	 * YOUR CODE HERE
	 *
	 * Try running `go test`.  Add more test as you code in `handle_test.go`.
	 */
	time.Sleep(500 * time.Millisecond)
	val, err := strconv.Atoi(string(e.Data()))
	if err != nil {
		return nil, err
	}
	val = val / 2
	if rand.Int()%100 < 15 {
		return nil, fmt.Errorf("random error in div2")
	}
	e.SetData("text/plain", strconv.Itoa(val))
	e.SetType("com.example.display")
	return &e, nil // echo to caller
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
