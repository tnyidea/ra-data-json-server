package addressv1

import (
	"github.com/gin-gonic/gin"
	"github.com/tnyidea/ra-data-json-server/go/data/gorm/model"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// http://localhost:8080/address?_end=10&_order=ASC&_sort=id&_start=0

func (p *Handler) GetMany(c *gin.Context) {
	if p.serviceDb == ServiceDbMongoDb {
		p.mongodbGetMany(c)
	}

	if p.serviceDb == ServiceDbPostgres {
		p.gormGetMany(c)
	}
}

func (p *Handler) gormGetMany(c *gin.Context) {
	// Determine if the request is one of:
	// getList GET http://my.api.url/post?_sort=title&_order=ASC&_start=0&_end=24&title=bar
	// getMany GET http://my.api.url/post?id=123&id=456&id=789
	// getManyReference GET http://my.api.url/post?author_id=345

	// Query Limit and Offset
	urlQueryStartIndex := c.Query("_start")
	urlQueryEndIndex := c.Query("_end")

	var queryOffset int64
	var queryLimit int64

	var start int64
	if urlQueryStartIndex != "" {
		v, err := strconv.ParseInt(urlQueryStartIndex, 10, 64)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		start = v
	}

	var end int64
	if urlQueryEndIndex != "" {
		v, err := strconv.ParseInt(urlQueryEndIndex, 10, 64)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		end = v
	}
	queryLimit = end - start
	queryOffset = start

	// Query Sort
	urlQuerySortField := c.Query("_sort")
	urlQuerySortOrder := c.Query("_order")

	var querySortOrder string
	if urlQuerySortField != "" {
		if urlQuerySortOrder != "" {
			urlQuerySortOrder = strings.ToUpper(urlQuerySortOrder)
			if urlQuerySortOrder != "ASC" && urlQuerySortOrder != "DESC" {
				log.Println("addressv1: GetMany(invalid value for _order: " + urlQuerySortOrder + ")")
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}
		}

		sortExpression := p.gormDb.Model(&model.Address{}).NamingStrategy.ColumnName("", urlQuerySortField)
		sortOrder := urlQuerySortOrder
		querySortOrder = strings.Trim(sortExpression+" "+sortOrder, " ")
	}

	// Execute the Query
	var addresses []model.Address
	tx := p.gormDb.Model(&model.Address{})

	if queryOffset != 0 {
		tx = tx.Offset(int(queryOffset))
	}
	if queryLimit != 0 {
		tx = tx.Limit(int(queryLimit))
	}
	if querySortOrder != "" {
		tx = tx.Order(querySortOrder)
	}
	tx = tx.Find(&addresses)
	if tx.Error != nil {
		log.Println(tx.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var count int64
	tx = p.gormDb.Model(&model.Address{}).Count(&count)
	if tx.Error != nil {
		log.Println(tx.Error)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if start == 0 && end == 0 {
		end = count
	}
	countString := strconv.FormatInt(count, 10)
	startString := strconv.FormatInt(start, 10)
	endString := strconv.FormatInt(end, 10)

	c.Header("Content-Range", startString+"-"+endString+"/"+countString)
	c.Header("X-Total-Count", countString)
	c.JSON(http.StatusOK, addresses)
}

func (p *Handler) mongodbGetMany(c *gin.Context) {

}
