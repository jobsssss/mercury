package tests

import (
	btsConfig "mercury/config"
	"mercury/pkg/config"
	"mercury/pkg/console"
	"os"
	"testing"
)

func initialize() {
	os.Chdir("../")
	btsConfig.Initialize()
	config.InitConfig("")
}

func Test_api(t *testing.T) {
	initialize()
	t.Log(config.Env("APP_URL"))
}

func Test_server(t *testing.T) {
	console.Success("server:okok")
}
