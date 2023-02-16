/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/Gocyber-world/navigator-demo/cmd"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPRIVATE=github.com/Gocyber-world
//go:generate go install github.com/swaggo/swag/cmd/swag@v1.7.9
//go:generate swag init
//go:generate go mod tidy
//go:generate go mod download

// @Navigator Demo Service API
// @version 1.0
// @description Navigator demo service api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
