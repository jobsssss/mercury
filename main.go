package main

import (
	"flag"
	"mercury/bootstrap"
	btsConfig "mercury/config"
	"mercury/pkg/config"
)

func init() {
	btsConfig.Initialize()
}

func main() {
	var env string
	flag.StringVar(&env, "env", "", "load .env, such as --env=testing, load are .env.testing")
	flag.Parse()
	config.InitConfig(env)

	bootstrap.SetupLogger()
	bootstrap.SetupDB()
	bootstrap.SetupRedis()
	bootstrap.Serve()
}
