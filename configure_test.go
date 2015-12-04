package main

import "testing"

func TestConfig(t *testing.T) {
	SetupConfig()

	c := Configuration{"abc", "go.cloud.exchange"}

	WriteConfig(c)

	c2 := GetConfig()

	println(c2.Url)
	println(c2.Key)
}
