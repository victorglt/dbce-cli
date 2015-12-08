package configuration

import "testing"

func TestConfig(t *testing.T) {
	SetupConfig()

	println(config.Url)
}
