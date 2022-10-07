package setting

import (
	"github.com/spf13/viper"
)

func ReadConfig() (*SampleParameter, error) {
	viper.SetConfigFile("configs/config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	//viper.WatchConfig()
	//viper.OnConfigChange(func(in fsnotify.Event) {
	//	fmt.Println("config修改了...")
	//})

	result := &SampleParameter{
		MongoHost:     viper.GetString("mongo.host"),
		MongoUserName: viper.GetString("mongo.username"),
		MongoPassword: viper.GetString("mongo.password"),
	}

	return result, nil
}

type SampleParameter struct {
	MongoHost     string
	MongoUserName string
	MongoPassword string
}
