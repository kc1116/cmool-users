package users

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/kc1116/cmool-events/events"
	neoism "gopkg.in/jmcvetta/neoism.v1"
)

var (
	Db *neoism.Database
)

func init() {
	var err error
	Db, err = neoism.Connect("http://neo4j:password@localhost:7474/db/data")
	if err != nil {
		panic(err)
	}
}

//API ... interface to be implemented by user struct
type API interface {
	attending() events.Event
	postVideo() bool
	postComment() bool
	postPhoto() bool
	sendFriendRequest() bool
	acceptFriendRequest() bool
}

// User ... event struct for neo4j event nodes
type User struct {
	Properties Properties
}

// Properties ... an users properties
type Properties struct {
	Name          string    `json:"Name"`
	DateJoined    time.Time `json:"Date"`
	Description   string    `json:"Description"`
	EventInterest []string  `json:"Keywords"`
	ProfilePhoto  string    `json:"Emblem"`
	City          string    `json:"City"`
	State         string    `json:"State"`
	UniqueID      string    `json:"UniqueID"`
}

// UserRelationships ... neo4j relationships associated with Event nodes
var UserRelationships = map[string]interface{}{
	"Liked":    "LIKED",
	"FriendOf": "FRIEND_OF",
}

// CreateUserNode . . . create a new user node from Event struct
func CreateUserNode(user User) (User, error) {
	uid := uuid.NewV4().String()
	node, err := Db.CreateNode(neoism.Props{
		"Name":          user.Properties.Name,
		"DateJoined":    user.Properties.DateJoined,
		"Description":   user.Properties.Description,
		"EventInterest": user.Properties.EventInterest,
		"ProfilePhoto":  user.Properties.ProfilePhoto,
		"City":          user.Properties.City,
		"State":         user.Properties.State,
		"UniqueID":      uid,
	})
	if err != nil {
		return user, err
	}

	node.AddLabel("User")
	user.Properties.UniqueID = uid
	return user, nil
}

// GetUserNode . . . get an user node. returns properties assiciated with that node
func GetUserNode(identifier string) (map[string]interface{}, error) {
	stmt := `
		MATCH (user:User)
		WHERE user.UniqueID = {uid}
		RETURN user
	`
	params := neoism.Props{
		"uid": identifier,
	}

	res := []struct {
		User neoism.Node
	}{}

	cq := neoism.CypherQuery{
		Statement:  stmt,
		Parameters: params,
		Result:     &res,
	}

	err := Db.Cypher(&cq)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		err := errors.New("User node not found.")
		return nil, err
	}

	return res[0].User.Data, nil
}

//Attending ...
func (user User) Attending(event events.Event) (events.Event, error) {
	rel := events.EventRelationships["IsAttending"]
	stmt := `
		MATCH (user:User),(event:Event)
        WHERE user.UniqueID = {userid} AND event.UniqueID = {eventid}
        CREATE UNIQUE (user)-[r:` + rel + `]->(event)
        RETURN r
	`
	params := neoism.Props{
		"userid":  user.Properties.UniqueID,
		"eventid": event.Properties.UniqueID,
	}

	// query results
	res := []struct {
		User  string `json:"user.Name"` // `json` tag matches column name in query
		Rel   string `json:"type(r)"`
		Event string `json:"event.Name"`
	}{}

	cq := neoism.CypherQuery{
		Statement:  stmt,
		Parameters: params,
		Result:     &res,
	}

	// execute query
	err := Db.Cypher(&cq)
	if err != nil {
		return event, err
	}
	//r := res[0]
	return event, nil

}
