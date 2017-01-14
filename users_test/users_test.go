package cmn4j_test

import (
	"testing"
	"time"

	"github.com/kc1116/cmool-events/events"
)

var db = events.Db

func TestCreateUserNode(t *testing.T) {
	var testEvent events.Event
	testEvent.Properties.Name = "Test"
	testEvent.Properties.DateCreated = time.Now()
	testEvent.Properties.Description = "This is a test event."
	testEvent.Properties.Keywords = []string{"key", "words"}
	testEvent.Properties.Rating = 3.5
	testEvent.Properties.TypeOfEvent = "Just a Test"
	testEvent.Properties.Emblem = "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcRjWFifmZY2WER6nCFNZOCtF2WSRm2vkDr3erTHUdTWFI8tCoQDJaXNJ5c"
	testEvent.Properties.Location.City = "Harrison"
	testEvent.Properties.Location.State = "New Jersey"
	testEvent.Properties.Location.StreetAddress = "1 Harrison ave"
	testEvent.Properties.Location.ZipCode = "07029"

	event, err := events.CreateEventNode(testEvent)
	if err != nil {
		t.Error("Expected an test event got an error:", err.Error())
	} else {
		t.Logf("TestCreateEventNode:%+v\n", event.Properties.UniqueID)
	}

}

func TestGetEventNode(t *testing.T) {
	uuid := "3e8c0e55-3971-436a-b76e-7414be2db023"

	event, err := events.GetEventNode(uuid)
	if err != nil {
		t.Error("Expected an test event got an error:", err.Error())
	} else {
		t.Logf("TestGetEventNode:%+v\n", event)
	}

}
