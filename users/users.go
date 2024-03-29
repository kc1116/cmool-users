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
	postVideo() bool
	sendFriendRequest() bool
	acceptFriendRequest() bool
	rate()
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
var UserRelationships = map[string]string{
	"Liked":        "LIKED",
	"FriendOf":     "FRIEND_OF",
	"EventComment": "EVENT_COMMENT",
	"EventPhoto":   "EVENT_PHOTO",
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
func (user User) Attending(eventID string) (string, error) {
	rel := events.EventRelationships["IsAttending"]
	stmt := `
		MATCH (user:User),(event:Event)
        WHERE user.UniqueID = {userid} AND event.UniqueID = {eventid}
        CREATE UNIQUE (user)-[r:` + rel + `]->(event)
        RETURN r
	`
	params := neoism.Props{
		"userid":  user.Properties.UniqueID,
		"eventid": eventID,
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
		return eventID, err
	}
	//r := res[0]
	return eventID, nil

}

//PostComment ...
func PostComment(comment string, eventID string, userID string) (string, error) {
	uid := uuid.NewV4().String()
	rel := UserRelationships["EventComment"]
	stmt := `MATCH (event:Event)
            WHERE event.UniqueID = {eid}
		    MERGE (comment:EventComment {User:{uid},DatePosted:{date},Comment:{comment},UniqueID:{uniqueID}})-[r:` + rel + `]->(event)
		    RETURN comment`

	params := neoism.Props{
		"uid":      userID,
		"eid":      eventID,
		"uniqueID": uid,
		"date":     time.Now(),
		"comment":  comment,
	}

	// query results
	res := []struct {
		Comment neoism.Node
	}{}

	cq := neoism.CypherQuery{
		Statement:  stmt,
		Parameters: params,
		Result:     &res,
	}

	// execute query
	err := Db.Cypher(&cq)
	if err != nil {
		return "", err
	}
	if len(res) == 0 {
		err := errors.New("Error adding comment to event in graph.")
		return "", err
	}

	return comment, nil
}

//PostPhoto ...
func PostPhoto(photoURI string, eventID string, userID string) (string, error) {
	uid := uuid.NewV4().String()
	rel := UserRelationships["EventPhoto"]
	stmt := `MATCH (event:Event)
            WHERE event.UniqueID = {eid}
		    MERGE (photo:EventPhoto {User:{uid},DatePosted:{date},Photo:{photoURI},UniqueID:{uniqueID}})-[r:` + rel + `]->(event)
		    RETURN photo`

	params := neoism.Props{
		"uid":      userID,
		"eid":      eventID,
		"uniqueID": uid,
		"date":     time.Now(),
		"photoURI": photoURI,
	}

	// query results
	res := []struct {
		Photo neoism.Node
	}{}

	cq := neoism.CypherQuery{
		Statement:  stmt,
		Parameters: params,
		Result:     &res,
	}

	// execute query
	err := Db.Cypher(&cq)
	if err != nil {
		return "", err
	}
	if len(res) == 0 {
		err := errors.New("Error while adding photo to event graph.")
		return "", err
	}

	return photoURI, nil
}

//LikeEvent ...
func LikeEvent(eventID string, userID string) error {
	rel := UserRelationships["Liked"]
	stmt := `
		MATCH (user:User),(event:Event)
        WHERE user.UniqueID = {uid} AND event.UniqueID = {eid}
		SET event.Likes = event.Likes + 1
        CREATE UNIQUE (user)-[r:` + rel + `]->(event)
        RETURN r`

	params := neoism.Props{
		"uid": userID,
		"eid": eventID,
	}

	// query results
	res := []struct {
		R neoism.Node
	}{}

	cq := neoism.CypherQuery{
		Statement:  stmt,
		Parameters: params,
		Result:     &res,
	}

	// execute query
	err := Db.Cypher(&cq)
	if err != nil {
		return err
	}
	if len(res) == 0 {
		err := errors.New("Error while liking event.")
		return err
	}

	return nil
}
