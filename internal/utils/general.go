package utils

import (
	"log"
	"strconv"

	"gorm.io/gorm"
)

func ErrorHandler(err error) {
	if err != nil {
		log.Fatalf("error %s", err.Error())
	}
}

func ValidatePage(p string) int {
	page, err := strconv.Atoi(p)
	if page == 0 || err != nil {
		page = 1
	}

	return page
}

func ValidatePageSize(s string) int {
	pageSize, err := strconv.Atoi(s)
	if err != nil {
		pageSize = 100
	}

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize < 5:
		pageSize = 5
	}

	return pageSize
}

func QueryBuilder(m map[string]string, db *gorm.DB) *gorm.DB {
	for k, v := range m {
		if len(v) > 0 {
			db.Where(k + " = ?" + v)
		}
	}

	return db
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1)
		return db.Offset(offset).Limit(pageSize)
	}

	// usage
	// db.Scopes(Paginate(r)).Find(&articles)
}

