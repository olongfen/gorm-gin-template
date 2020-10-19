package models

import (
	"fmt"
	"project/src/models/model_log"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"project/src/pkg/log_project"
	"project/src/setting"
)

var(
	DB *gorm.DB
)

// init 初始化数据库模型
func init()  {
	var(
		err error
	)
	model_log.LogModel = log_project.NewLogFile(log_project.ParamLog{Path: setting.GlobalConfig.FilePath.LogDir + "/" + "model", Stdout: setting.DevEnv, P: setting.GlobalConfig.FilePath.LogPatent})
	dataSourceName := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", setting.GlobalConfig.DB.Driver, setting.GlobalConfig.DB.Username,
		setting.GlobalConfig.DB.Password, setting.GlobalConfig.DB.Host, setting.GlobalConfig.DB.Port, setting.GlobalConfig.DB.DatabaseName)
	fmt.Println(dataSourceName)
	if DB, err = gorm.Open(postgres.Open(dataSourceName), nil); err != nil {
		logrus.Fatal(err)
	}
	if setting.DevEnv {
		DB = DB.Debug()
	}
	err = DB.AutoMigrate()
	if err != nil {
		panic(err)
	}

	logrus.Infoln("database init success !")
}