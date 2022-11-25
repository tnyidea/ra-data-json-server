package endpointsv1

import (
	"github.com/gin-gonic/gin"
	datatypes "github.com/tnyidea/react-admin/dataprovider/go/data/types"
	"net/http"
)

func GetAllAddresses(c *gin.Context) {
	// https://jsonplaceholder.typicode.com/users?_end=10&_order=ASC&_sort=id&_start=0
	var addresses []datatypes.Address
	c.IndentedJSON(http.StatusOK, addresses)
}

func GetAddress(c *gin.Context) {
	var address datatypes.Address
	c.IndentedJSON(http.StatusOK, address)
}

//func FindAllUsersV1(r *http.Request, db models.DB) response.DataProviderResponseData {
//	log.Println("=== Executing FindAllUsersV1 ===")
//	defer log.Println("=== FindAllUsersV1 Execution Complete ===")
//
//
//	query := r.URL.Query()
//	listStart, err := strconv.Atoi(query.Get("_start"))
//	if err != nil {
//		return response.DataProviderResponseData{
//			Status: http.StatusBadRequest,
//		}
//	}
//	listEnd, err := strconv.Atoi(query.Get("_end"))
//	if err != nil {
//		return response.DataProviderResponseData{
//			Status: http.StatusBadRequest,
//		}
//	}
//	listSortColumn := query.Get("_sort")
//	if listSortColumn == "" {
//		return response.DataProviderResponseData{
//			Status: http.StatusBadRequest,
//		}
//	}
//	listSortOrder := query.Get("_order")
//	if listSortOrder == "" {
//		return response.DataProviderResponseData{
//			Status: http.StatusBadRequest,
//		}
//	}
//
//	users, err := db.FindAllUsers()
//	if err != nil {
//		return response.DataProviderResponseData{
//			Status: http.StatusInternalServerError,
//		}
//	}
//	totalCount, err := db.Count()
//	if err != nil {
//		return response.DataProviderResponseData{
//			Status: http.StatusInternalServerError,
//		}
//	}
//
//	// TODO determine best way to sort on other struct fields
//	if listSortOrder == "ASC" {
//		sort.Slice(users, func(i, j int) bool {
//			return users[i].Id < users[j].Id
//		})
//	}
//	if listSortOrder == "DESC" {
//		sort.Slice(users, func(i, j int) bool {
//			return users[i].Id > users[j].Id
//		})
//	}
//
//	return response.DataProviderResponseData{
//		Status:     http.StatusOK,
//		Items:      users[listStart:listEnd],
//		TotalCount: totalCount,
//	}
//}
