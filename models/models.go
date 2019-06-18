package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/luohao-brian/openstack-admin/pkg/logging"
	"github.com/luohao-brian/openstack-admin/pkg/setting"
)

var db *gorm.DB

func Setup() {
	var err error
	logging.Info(fmt.Sprintf("NovaDB: %s:*@tcp(%s:%s)/%s", setting.NovaSetting.User, setting.NovaSetting.Host, setting.NovaSetting.Port, setting.NovaSetting.Name))

	var connstr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.NovaSetting.User,
		setting.NovaSetting.Password,
		setting.NovaSetting.Host,
		setting.NovaSetting.Port,
		setting.NovaSetting.Name)
	db, err = gorm.Open(setting.NovaSetting.Type, connstr)
	if err != nil {
		logging.Error(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName + "s"
	}

	db.LogMode(true)

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
