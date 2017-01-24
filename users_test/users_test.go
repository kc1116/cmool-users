package cmn4j_test

import (
	"testing"

	"github.com/kc1116/cmool-events/events"
	"github.com/kc1116/cmool-users/users"
)

var db = events.Db

/*func TestCreateUserNode(t *testing.T) {
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

func TestPostComment(t *testing.T) {
	comment := "This party is great!"
	eventID := "3e8c0e55-3971-436a-b76e-7414be2db023"
	userID := "daa72fb3-9964-43be-8288-dece8389eba6"

	c, err := users.PostComment(comment, eventID, userID)
	if err != nil {
		t.Error("Expected an comment got an error:", err.Error())
	} else {
		t.Logf("TestPostComment:%+v\n", c)
	}
}

func TestPostPhoto(t *testing.T) {
	photoURI := "https://www.google.com/url?sa=i&rct=j&q=&esrc=s&source=images&cd=&cad=rja&uact=8&ved=0ahUKEwin7rWwjNrRAhXFdSYKHX1PBAIQjRwIBw&url=https%3A%2F%2Ftwitter.com%2Fneo4j&bvm=bv.144686652,d.eWE&psig=AFQjCNHrTojApeJcEGsumzMdtloVirC0og&ust=1485323771031443"
	eventID := "3e8c0e55-3971-436a-b76e-7414be2db023"
	userID := "daa72fb3-9964-43be-8288-dece8389eba6"

	c, err := users.PostPhoto(photoURI, eventID, userID)
	if err != nil {
		t.Error("Expected an photo URI got an error:", err.Error())
	} else {
		t.Logf("TestPostPhoto:%+v\n", c)
	}
}*/
func TestLikeEvent(t *testing.T) {
	eventID := "ff5cabd1-c941-44b7-b6dc-2d31c5ac69a4"
	userID := "daa72fb3-9964-43be-8288-dece8389eba6"

	err := users.LikeEvent(eventID, userID)
	if err != nil {
		t.Error("Expected nil got an error:", err.Error())
	} else {
		t.Logf("TestLikeEvent: Success")
	}
}
