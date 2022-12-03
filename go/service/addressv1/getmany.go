package addressv1

import (
	"context"
	"github.com/gin-gonic/gin"
	gormmodel "github.com/tnyidea/ra-data-json-server/go/data/gorm/model"
	mongomodel "github.com/tnyidea/ra-data-json-server/go/data/mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

		sortExpression := p.gormDb.Model(&gormmodel.Address{}).NamingStrategy.ColumnName("", urlQuerySortField)
		sortOrder := urlQuerySortOrder
		querySortOrder = strings.Trim(sortExpression+" "+sortOrder, " ")
	}

	// Execute the Query
	var addresses []gormmodel.Address
	tx := p.gormDb.Model(&gormmodel.Address{})

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
	tx = p.gormDb.Model(&gormmodel.Address{}).Count(&count)
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

	var querySortField string
	var querySortOrder int
	if urlQuerySortField != "" {
		if urlQuerySortOrder != "" {
			urlQuerySortOrder = strings.ToUpper(urlQuerySortOrder)
			if urlQuerySortOrder != "ASC" && urlQuerySortOrder != "DESC" {
				log.Println("addressv1: GetMany(invalid value for _order: " + urlQuerySortOrder + ")")
				c.AbortWithStatus(http.StatusBadRequest)
				return
			}

			querySortField = urlQuerySortField
			querySortOrder = 1 // ASC
			if urlQuerySortOrder == "DESC" {
				querySortOrder = -1
			}

		}
	}

	// Execute the Query
	collection := p.mongoDb.Collection("address")

	//findOptions := options.Find().SetSort(bson.D{{"rating", -1}}).SetLimit(2).SetSkip(1)
	findOptions := options.Find()
	if querySortField != "" {
		findOptions = findOptions.SetSort(bson.D{{querySortField, querySortOrder}})
	}
	if queryOffset != 0 {
		findOptions = findOptions.SetSkip(queryOffset)
	}
	if queryLimit != 0 {
		findOptions = findOptions.SetLimit(queryLimit)
	}

	var addresses []mongomodel.Address
	cursor, err := collection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = cursor.All(context.Background(), &addresses)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	count, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		log.Println(err)
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
