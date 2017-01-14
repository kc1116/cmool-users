package cmn4j_test

import (
	"testing"
	"time"

	"github.com/kc1116/cmool-events/events"
	"github.com/kc1116/cmool-users/users"
)

var db = events.Db

func TestCreateUserNode(t *testing.T) {
	var testUser users.User
	testUser.Properties.Name = "Test"
	testUser.Properties.DateJoined = time.Now()
	testUser.Properties.Description = "This is a test users."
	testUser.Properties.EventInterest = []string{"parties", "career", "meetup"}
	testUser.Properties.ProfilePhoto = "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcRjWFifmZY2WER6nCFNZOCtF2WSRm2vkDr3erTHUdTWFI8tCoQDJaXNJ5c"
	testUser.Properties.City = "Harrison"
	testUser.Properties.State = "New Jersey"

	user, err := users.CreateUserNode(testUser)
	if err != nil {
		t.Error("Expected an test user got an error:", err.Error())
	} else {
		t.Logf("TestCreateUserNode:%+v\n", user.Properties.UniqueID)
	}

}

func TestGetUserNode(t *testing.T) {
	uuid := "daa72fb3-9964-43be-8288-dece8389eba6"

	user, err := users.GetUserNode(uuid)
	if err != nil {
		t.Error("Expected an test user got an error:", err.Error())
	} else {
		t.Logf("TestGetUserNode:%+v\n", user)
	}

}

func TestUserAttending(t *testing.T) {
	var testUser users.User
	testUser.Properties.Name = "Test"
	testUser.Properties.DateJoined = time.Now()
	testUser.Properties.Description = "This is a test users."
	testUser.Properties.EventInterest = []string{"parties", "career", "meetup"}
	testUser.Properties.ProfilePhoto = "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcRjWFifmZY2WER6nCFNZOCtF2WSRm2vkDr3erTHUdTWFI8tCoQDJaXNJ5c"
	testUser.Properties.City = "Harrison"
	testUser.Properties.State = "New Jersey"
	testUser.Properties.UniqueID = "daa72fb3-9964-43be-8288-dece8389eba6"

	testEventID := "3e8c0e55-3971-436a-b76e-7414be2db023"

	eid, err := testUser.Attending(testEventID)
	if err != nil {
		t.Error("Expected an event ID got an error:", err.Error())
	} else {
		t.Logf("TestUserAttending:%+v\n", eid)
	}
}
