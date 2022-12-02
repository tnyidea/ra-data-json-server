package model

import (
	"context"
	_ "embed"
	"encoding/json"
	"github.com/tnyidea/typeutils"
	"log"
	"os"
	"testing"
)

var env map[string]string

//go:embed testdata/us-500.json
var sampleData []byte

func init() {
	env = map[string]string{
		"MONGODB_URL": os.Getenv("MONGODB_URL"),
	}
	err := typeutils.MapNoEmptyValues(env)
	if err != nil {
		log.Fatal(err)
	}
}

func TestNewMongoDbClient(t *testing.T) {
	url := env["MONGODB_URL"]

	db, err := NewMongoDbSession(url)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	err = db.Client().Ping(context.TODO(), nil)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
}

func TestLoadSampleData(t *testing.T) {
	url := env["MONGODB_URL"]

	db, err := NewMongoDbSession(url)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	defer func() {
		_ = db.Client().Disconnect(context.Background())
	}()

	var addressList []interface{}
	err = json.Unmarshal(sampleData, &addressList)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	// begin insertMany
	collection := db.Collection("address")
	err = collection.Drop(context.Background())
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	_, err = collection.InsertMany(context.TODO(), addressList)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
}
