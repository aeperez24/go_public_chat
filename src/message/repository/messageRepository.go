package repository

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"

	"gopkg.in/mgo.v2/bson"

	"config/mongo"
	message "message"
)

type messageRepository struct {
}
type MessageRepository interface {
	SaveMessage(message.Message)
	GetMessages(time.Time, time.Time) []message.Message
}
type MessageEntity struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Date time.Time
	Text string
	Name string
}

func (rep messageRepository) SaveMessage(m message.Message) {
	sessionProv := mongo.MongoSessionProvider
	session := sessionProv.GetSession()
	defer session.Close()
	databaseName := viper.GetString("databaseInfo.database")         //measurements1
	collectionName := viper.GetString("databaseInfo.collectionName") //temperature
	c := session.DB(databaseName).C(collectionName)
	// TODO:FIX
	// c.Insert(&MessageEntity{Date: m.Date, Name: m.Name, Text: m.Text})
	c.Insert(&MessageEntity{Date: m.Date, Name: m.Name, Text: m.Text})
}

func (rep messageRepository) GetMessages(start time.Time, end time.Time) []message.Message {
	readQuery := viper.GetString("query.mongo.chat_message.datestart_dateend")
	ssdate := fmt.Sprintf("%v", start)
	sedate := fmt.Sprintf("%v", end)
	myMap := map[string]string{"inicio": ssdate, "fin": sedate}
	replacedQuery := prepareQuery(readQuery, myMap)
	query := make(bson.M)
	bson.Unmarshal([]byte(replacedQuery), &query)
	// query2 := bson.M{"date": bson.M{"$gt": start, "$lt": end}}
	result := rep.GetMessageRepositoryByQuery(query)
	return result
}

func NewMessageRepository() MessageRepository {
	return &messageRepository{}
}

func (rep messageRepository) GetMessageRepositoryByQuery(query interface{}) []message.Message {
	sessionProv := mongo.MongoSessionProvider
	session := sessionProv.GetSession()

	defer session.Close()
	databaseName := viper.GetString("databaseInfo.database")         //chat
	collectionName := viper.GetString("databaseInfo.collectionName") //chatMessage

	c := session.DB(databaseName).C(collectionName) // var us userEntity
	result := []message.Message{}
	err := c.Find(query).All(&result)
	if err != nil {
		log.Printf("error en la consulta %v", err)
	}

	return result
}
func prepareQuery(query string, params map[string]string) string {
	for key := range params {
		value := params[key]
		query = strings.Replace(query, "$"+key, value, -1)
	}
	return query
}
