package model

import (
	_ "embed"
	"encoding/json"
	"github.com/tnyidea/typeutils"
	"log"
	"os"
	"testing"
)

var env = make(map[string]string)

//go:embed us-500.json
var sampleData []byte

func init() {
	env["POSTGRES_URL"] = os.Getenv("POSTGRES_URL")
	err := typeutils.MapNoEmptyValues(env)
	if err != nil {
		log.Fatal(err)
	}
}

func TestNewGormDbSession(t *testing.T) {
	url := env["POSTGRES_URL"]

	_, err := NewGormDbSession(url)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
}

func TestLoadSampleData(t *testing.T) {

	url := env["POSTGRES_URL"]
	db, err := NewGormDbSession(url)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	// Perform a soft delete
	// db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Address{}).Commit()

	// Perform a hard delete
	db.Exec("DELETE FROM addresses")

	err = db.AutoMigrate(&Address{})
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	var addressList []Address
	err = json.Unmarshal(sampleData, &addressList)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	for _, address := range addressList {
		db.Create(&address)
	}
}
