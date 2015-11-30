package main

import (
	"log"
	"os"
	"os/user"
)

const (
	CONFIG_FILE = ".dbce"
)

type Key struct {
	name string
	key  string
}

type Environment struct {
	url  string
	keys []Key
}

type Configuration struct {
	environments []Environment
}

func WriteConfig() {

}

func GetConfig() {

}

func SetupConfig() {

	user, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	home := user.HomeDir

	filePath := home + string(os.PathSeparator) + CONFIG_FILE
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(filePath)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}

	}
}
