package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Service struct {
	UUID       string
	Host       string
	LastSeenUp string
	Binary     string
	CreatedAt  time.Time
}

func GetServices(pageNum int, pageSize int, maps interface{}) ([]Service, int, error) {
	var services []Service
	dbNew := db
	if pageSize > 0 && pageNum > 0 {
		dbNew = db.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	}

	b := []string{"nova-compute"}
	err := dbNew.Where(maps).Where("services.binary in(?)", b).Find(&services).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	var total int
	_ = db.Model(&Service{}).Where(maps).Where("services.binary in(?)", b).Count(&total).Error

	return services, total, nil
}
