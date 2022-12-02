package addressv1

import (
	"errors"
	gormmodel "github.com/tnyidea/ra-data-json-server/go/data/gorm/model"
	mongodbmodel "github.com/tnyidea/ra-data-json-server/go/data/mongodb/model"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

const (
	ServiceDbMongoDb  = "MONGODB"
	ServiceDbPostgres = "POSTGRES"
)

type Handler struct {
	serviceDb string
	gormDb    *gorm.DB
	mongoDb   *mongo.Database
}

func NewHandler(env map[string]string) (Handler, error) {
	serviceDb := env["SERVICE_DB"]
	if serviceDb != ServiceDbMongoDb && serviceDb != ServiceDbPostgres {
		return Handler{}, errors.New("addressv1: NewHandler(SERVICE_DB must be one of 'MONGODB' or 'POSTGRES')")
	}
	handler := Handler{
		serviceDb: serviceDb,
	}

	if serviceDb == ServiceDbPostgres {
		db, err := gormmodel.NewGormDbSession(env["POSTGRES_URL"])
		if err != nil {
			return Handler{}, err
		}
		handler.gormDb = db
	}
	if serviceDb == ServiceDbMongoDb {
		db, err := mongodbmodel.NewMongoDbSession(env["MONGODB_URL"])
		if err != nil {
			return Handler{}, err
		}
		handler.mongoDb = db
	}

	return handler, nil
}
