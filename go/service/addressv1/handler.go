package addressv1

import (
	"github.com/tnyidea/ra-data-json-server/go/data/gorm/model"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(env map[string]string) (Handler, error) {
	db, err := model.NewGormDbSession(env["POSTGRES_URL"])
	if err != nil {
		return Handler{}, err
	}
	return Handler{
		db: db,
	}, nil
}
