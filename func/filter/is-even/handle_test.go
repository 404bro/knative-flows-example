package function

import (
	"context"
	"testing"

	"github.com/cloudevents/sdk-go/v2/event"
)

// TestHandle ensures that Handle accepts a valid CloudEvent without error.
func TestHandleOdd(t *testing.T) {
	// Assemble
	e := event.New()
	e.SetID("id")
	e.SetType("type")
	e.SetSource("source")
	e.SetData("text/plain", "1")

	// Act
	echo, err := Handle(context.Background(), e)
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	if echo != nil {
		t.Errorf("data: 1, but received non-nil event from is-odd")
	}
}

func TestHandleEven(t *testing.T) {
	// Assemble
	e := event.New()
	e.SetID("id")
	e.SetType("type")
	e.SetSource("source")
	e.SetData("text/plain", "2")

	// Act
	echo, err := Handle(context.Background(), e)
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	if echo == nil {
		t.Errorf("data: 2, but received nil event from is-odd")
	} else {
		if string(echo.Data()) != "2" {
			t.Errorf("the received event expected data to be '2', got '%s'", echo.Data())
		}
	}
}
