package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"os/user"
)

const (
	fileName = ".dbce"
)

var (
	filePath string
)

type Configuration struct {
	Key string
	Url string
}

func WriteConfig(c Configuration) {
	file, err := os.OpenFile(filePath, os.O_WRONLY, os.ModePerm)

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(file)
	w.Write(b)
	w.Flush()
}

func GetConfig() (c *Configuration) {
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(file)

	result := make([]byte, 100)
	buffer := make([]byte, 100)
	for {
		count, err := r.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}

		if count == 0 {
			break
		}

		println(count)
		println("fuuu")
		result = append(result, buffer[0:count]...)
	}

	config := Configuration{}
	println("unmarshalling")
	json.Unmarshal(result, &config)
	return &config
}

func SetupConfig() {

	user, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	home := user.HomeDir

	filePath = home + string(os.PathSeparator) + fileName
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
