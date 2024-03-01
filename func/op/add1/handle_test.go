package function

import (
	"context"
	"testing"

	"github.com/cloudevents/sdk-go/v2/event"
)

// TestHandle ensures that Handle accepts a valid CloudEvent without error.
func TestHandle(t *testing.T) {
	// Assemble
	e := event.New()
	e.SetID("id")
	e.SetType("type")
	e.SetSource("source")
	e.SetData("text/plain", "5")

	// Act
	echo, err := Handle(context.Background(), e)
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	if echo == nil {
		t.Errorf("received nil event") // fail on nil
	} else if string(echo.Data()) != "6" {
		t.Errorf("the received event expected data to be '6', got '%s'", echo.Data())
	}
}
