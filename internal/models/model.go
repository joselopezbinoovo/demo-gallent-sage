package models

import (
	"reflect"
	"strings"
	"time"
	"vg-sage/internal/utils"

	"fmt"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Filter struct {
	Where     map[string]interface{} `json:"where,omitempty"`
	DateLimit time.Time              `json:"datelimit,omitempty"`
	Limit     int                    `json:"limit,omitempty"`
	Offset    int                    `json:"offset,omitempty"`
	Order     string                 `json:"order,omitempty"`
	Skip      int                    `json:"skip,omitempty"`
}

func Find(filter Filter, point interface{}) error {
	query := GetDB()

	joinQuery, ok := DetectJoin(point, query)
	if ok {
		query = joinQuery
	}

	//Is limitDate is true we limit with the date from the model
	if !filter.DateLimit.IsZero() {
		query = DateLimit(point, query, filter.DateLimit)
	}
	if filter.Where != nil {
		query = query.Where(filter.Where)
	}
	if filter.Limit != 0 {
		query = query.Limit(filter.Limit)
	}
	if filter.Offset != 0 {
		query = query.Offset(filter.Offset)
	}
	if filter.Skip != 0 {
		query = query.Offset(filter.Skip)
	}
	if filter.Order != "" {
		query = query.Order(filter.Order)
	} else {
		query = query.Order(getOrder(point))
	}

	return query.Find(point).Error
}

func getOrder(point interface{}) string {
	reflectType := utils.GetType(point)
	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		value, ok := field.Tag.Lookup("noovo")
		if ok && strings.Contains(value, "defaultorder") {
			return field.Name
		}
	}
	return reflectType.Field(0).Name
}

// Get the column with the tad limit date
func getLimitDateColumn(point interface{}) string {
	reflectType := utils.GetType(point)
	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		value, ok := field.Tag.Lookup("noovo")
		if ok {
			args := strings.Split(value, ";")
			for _, value := range args {
				if value == "datelimit" {
					return field.Name
				}
			}
		}
	}
	return ""
}

func getPrimaryKey(point interface{}) string {
	reflectType := utils.GetType(point)
	for i := 0; i < reflectType.NumField(); i++ {
		field := reflectType.Field(i)
		value, ok := field.Tag.Lookup("noovo")
		if ok && strings.Contains(value, "primaryKey") {
			return field.Name
		}
	}
	return ""
}

func DateLimit(point interface{}, query *gorm.DB, limit time.Time) *gorm.DB {
	column := getLimitDateColumn(point)
	tableName := query.NewScope(point).TableName()
	query = query.Where(tableName+"."+column+" > ?", limit)
	return query
}

func DetectJoin(point interface{}, query *gorm.DB) (*gorm.DB, bool) {
	reflectType := utils.GetType(point)
	reflectValue := reflect.New(reflectType)
	methodValue := reflectValue.MethodByName("Join")

	if methodValue.IsValid() {
		return methodValue.Interface().(func(query *gorm.DB) *gorm.DB)(query), true
	}
	return nil, false
}

func FindAll(point []interface{}) error {
	return GetDB().Find(point).Error
}

func FindOne(filter Filter, point interface{}) error {
	return GetDB().Where(filter.Where).Order(filter.Order).First(point).Error
}

func FindOneQuery(where interface{}, point interface{}) error {
	query := GetDB()
	joinQuery, ok := DetectJoin(point, query)
	if ok {
		query = joinQuery
	}
	if where != nil {
		query = query.Where(where)
	}

	result := query.Find(point).Error

	return result
}

func Patch(point interface{}, id uuid.UUID, data interface{}) error {
	pointCopy := copyInterface(point)

	fmt.Println(GetDB().Debug().Model(pointCopy).Where(map[string]interface{}{getPrimaryKey(pointCopy): id}).Updates(data).Error)

	return GetDB().Model(pointCopy).Where(map[string]interface{}{getPrimaryKey(pointCopy): id}).Updates(data).Error
}

func copyInterface(point interface{}) interface{} {
	copy := reflect.New(reflect.ValueOf(point).Elem().Type()).Interface()
	return copy
}
