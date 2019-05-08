package main

import (
	"testing"

	"github.com/spf13/viper"
)

func Test_main(t *testing.T) {
	//main()
}

func Test_setupLogging(t *testing.T) {
	levels := []string{
		"info",
		"warn",
		"error",
		"debug",
		"default",
	}

	for _, level := range levels {
		viper.Set("level", level)
		setupLogging()
	}
}
