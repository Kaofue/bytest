package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

//bson.D{} slice
//bson.M{} map
//bson.A{} array
//bson.E{} type
//https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.9.0/bson


type Mongodb struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
	cursor     *mongo.Cursor
}

func (m *Mongodb) Connection( host,username,password string) error {
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://" + username + ":" + password + host).
		SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	m.client = client

	if err = m.client.Ping(context.TODO(), nil); err != nil {
		fmt.Println("連線失敗")
		return err
	}
	//fmt.Println("連線ok")


	m.collection = m.client.Database("test").Collection("accounts")

	return nil
}

func (m *Mongodb) ReadALL(dataBase string, collection string) interface{}{
	m.collection = m.client.Database(dataBase).Collection(collection)

	cursor, err := m.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
		return nil
	}


	var A [] interface{}
	if err = cursor.All(context.TODO(), &A); err != nil {
		fmt.Println(err)
	}

	return  A
}

func (m *Mongodb) Insert(data interface{}) error {
	_, err := m.collection.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}


