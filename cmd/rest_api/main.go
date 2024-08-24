package main

import (
	"payment/api/rest"
	"payment/pkg/config"
)

func main() {
	config.ConnectPgsql()
	rest.StartApp().Run(":8080")
}