package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"gin-web-init/config"
	"gin-web-init/controller"
	"os"
)

func main() {
	config.InitConfig()

	if err := endless.ListenAndServe(":8080", controller.R); err != nil {
		fmt.Printf("Listen and serve jenscan_codeguard on %s failed, failed reason: %s\n", ":8080", err.Error())
		os.Exit(255)
	}
}
