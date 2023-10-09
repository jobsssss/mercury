package tests

import (
	"fmt"
	"mercury/pkg/config"
	"mercury/pkg/console"
	"testing"
)

func Test_api(t *testing.T) {
	fmt.Println(config.Env("APP_URL"))
	console.Success("okok")
}

func Test_server(t *testing.T) {
	console.Success("server:okok")
}
