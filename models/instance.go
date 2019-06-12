package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Instance struct {
	UUID string
	DisplayName string
	Host string
	Vcpus int
	MemoryMb int
	VmState string
	CreatedAt time.Time
}


func GetInstances(pageNum int, pageSize int, maps interface{}) ([]Instance, int, error) {
	var instances []Instance
	dbNew := db
	if pageSize > 0 && pageNum > 0 {
		dbNew = db.Offset((pageNum - 1) * pageSize).Limit(pageSize)
	}

	err := dbNew.Where(maps).Find(&instances).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

        var total int
	_ = db.Model(&Instance{}).Where(maps).Count(&total).Error

	return instances, total, nil
}

