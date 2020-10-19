package setting

import (
	"project/src/pkg/log_project"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"os"
	"strings"
)

// GlobalSettings 全局配置
type GlobalSettings struct {
	FilePath FilePath
	DB       Database
	Serve    Serve
}

type Serve struct {
	ServerAddr string
	ServerPort string
}

// Database 数据库连接参数
type Database struct {
	Host         string
	Port         string
	Driver       string
	DatabaseName string
	Username     string
	Password     string
	MaxIdleConn  int
	MaxOpenConn  int
}

// FilePath 文件保存
type FilePath struct {
	LogDir       string // 日志保存地址
	LogPatent    string // 日志格式
}

var (
	GlobalConfig = new(GlobalSettings)
	DevEnv       = false
)

func init()  {
	var (
		err        error
		configFile string
	)
	if err = gotenv.Load("./config/.env"); err != nil {
		log_project.Fatal(err)
	}
	env := os.Getenv("ENVIRONMENT")
	switch {
	case strings.Contains(env, "prod"):
		configFile = "./config/prod-global-config.yaml"
	case strings.Contains(env, "test"):
		configFile = "./config/test-global-config.yaml"
	default: // default is dev
		DevEnv = true
		configFile = "./config/dev-global-config.yaml"

	}
	viper.SetConfigFile(configFile)
	_ = viper.ReadInConfig()
	if err = viper.Unmarshal(GlobalConfig); err != nil {
		log_project.Fatal(err)
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		log_project.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
		if err = viper.Unmarshal(GlobalConfig); err != nil {
			log_project.Fatal(err)
		}
	})
	log_project.Infoln("setting init success !")
}