// Package main is the entry point of the application
package main

import (
	"github.com/rakib-09/golang-clean-architecture/cmd"
	_ "github.com/spf13/viper/remote"
)

func main() {
	cmd.Execute()
}
