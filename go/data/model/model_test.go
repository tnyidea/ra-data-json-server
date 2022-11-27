package model

import (
	"github.com/tnyidea/typeutils"
	"log"
	"os"
	"testing"
)

var env = make(map[string]string)

func init() {
	env["POSTGRES_URL"] = os.Getenv("POSTGRES_URL")
	env["ADDRESS_MODEL_SAMPLE_DATA"] = os.Getenv("ADDRESS_MODEL_SAMPLE_DATA")
	err := typeutils.MapNoEmptyValues(env)
	if err != nil {
		log.Fatal(err)
	}
}

func TestNewDatabaseConnection(t *testing.T) {
	url := env["POSTGRES_URL"]

	_, err := NewDatabaseConnection(url)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
}

func TestLoadSampleData(t *testing.T) {
	url := env["POSTGRES_URL"]
	filename := env["ADDRESS_MODEL_SAMPLE_DATA"]

	db, err := NewDatabaseConnection(url)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	err = db.LoadSampleData(filename)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

}
