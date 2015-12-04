package main

import (
	"bufio"
	"encoding/json"
	"io"
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

func WriteConfig(c *Configuration) {
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

	var result []byte
	buffer := make([]byte, 128)
	for {
		count, err := r.Read(buffer[0:])
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		result = append(result, buffer[0:count]...)
	}

	config := Configuration{}
	println("unmarshalling")
	json.Unmarshal(result, &config)
	return &config
}

func SetupConfig() (c *Configuration) {

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
	return GetConfig()
}
