package config

import (
	"mercury/pkg/helpers"
	"os"

	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
)

var viper *viperlib.Viper

type ConfigFunc func() map[string]interface{}

var ConfigFuncs map[string]ConfigFunc

func init() {

	// Make a viper instance
	viper = viperlib.New()

	// Set config type, support "env","dotenv","json", "toml", "yaml", "yml", "properties","props", "prop"
	viper.SetConfigType("env")

	// The env file relative path of main.go
	viper.AddConfigPath(".")

	// Set env prefix in order to diff go system env
	viper.SetEnvPrefix("appenv")

	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

// Initialize config information
func InitConfig(env string) {

	loadEnv(env)

	loadConfig()
}

func loadConfig() {
	for name, fn := range ConfigFuncs {
		viper.Set(name, fn())
	}
}

func loadEnv(suffix string) {

	path := ".env"

	if len(suffix) > 0 {
		filepath := ".env." + suffix
		if _, err := os.Stat(filepath); err == nil {
			path = filepath
		}
	}

	viper.SetConfigName(path)
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}

	viper.WatchConfig()
}

func Env(name string, defaultVal ...interface{}) interface{} {
	if len(defaultVal) > 0 {
		return selfGet(name, defaultVal[0])
	}

	return selfGet(name)
}

func Add(name string, fn ConfigFunc) {
	ConfigFuncs[name] = fn
}

func Get(path string, defaultVal ...interface{}) string {
	return GetString(path, defaultVal...)
}

// @todo 该函数获取配置的值，如果配置不存在返回Nil
// 如果配置中缺少了一个值，程序能正常启动，但是在运行过程中会报一个服务器错误而不给出具体原因，加大调试困难,后续要重构
func selfGet(path string, defaultVal ...interface{}) interface{} {
	if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return nil
	}
	return viper.Get(path)
}

func GetInt(path string, defaultVal ...interface{}) int {
	return cast.ToInt(selfGet(path, defaultVal...))
}

func GetFloat64(path string, defaultVal ...interface{}) float64 {
	return cast.ToFloat64(selfGet(path, defaultVal...))
}

func GetInt64(path string, defaultValue ...interface{}) int64 {
	return cast.ToInt64(selfGet(path, defaultValue...))
}

func GetUint(path string, defaultValue ...interface{}) uint {
	return cast.ToUint(selfGet(path, defaultValue...))
}

func GetBool(path string, defaultValue ...interface{}) bool {
	return cast.ToBool(selfGet(path, defaultValue...))
}

func GetString(path string, defaultVal ...interface{}) string {
	return cast.ToString(selfGet(path, defaultVal...))
}

func GetStringMapString(path string) map[string]string {
	return viper.GetStringMapString(path)
}
