package addressv1

import (
	"github.com/gin-gonic/gin"
	"github.com/tnyidea/ra-data-json-server/go/data/gorm/model"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
)

func (p *Handler) Delete(c *gin.Context) {
	// DELETE http://my.api.url/post/123

	var address model.Address
	tx := p.db.Model(&model.Address{}).Clauses(clause.Returning{}).Where("id = ?", c.Param("id")).Delete(&address)
	log.Println(tx.RowsAffected)
	if tx.Error != nil {
		log.Println(tx.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if tx.RowsAffected == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, address)
}
