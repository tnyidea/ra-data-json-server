package model

import (
	_ "embed"
	"github.com/tnyidea/typeutils"
	"log"
	"os"
	"testing"
)

var env = make(map[string]string)

//go:embed us-500.json
var sampleData []byte

func init() {
	env["MONGODB_URL"] = os.Getenv("MONGODB_URL")
	err := typeutils.MapNoEmptyValues(env)
	if err != nil {
		log.Fatal(err)
	}
}

func TestNewMongoDbClient(t *testing.T) {
	url := env["MONGODB_URL"]

	_, err := NewMongoDbClient(url)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
}
