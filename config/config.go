package config

import "os"

const (
	PORT = ":8080"
)

var Dir string

func init() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	Dir = dir
}
