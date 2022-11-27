package addressv1

import (
	"github.com/tnyidea/ra-data-json-server/go/data/model"
)

type Handler struct {
	db model.DB
}

func NewHandler(env map[string]string) (Handler, error) {
	db, err := model.NewDatabaseConnection(env["POSTGRES_URL"])
	if err != nil {
		return Handler{}, err
	}
	return Handler{
		db: db,
	}, nil
}
