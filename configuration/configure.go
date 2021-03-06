package configuration

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
	Context  *Configuration
)

type Configuration struct {
	Key string
	Url string
}

func WriteConfig() {
	file, err := os.OpenFile(filePath, os.O_WRONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(Context)
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
			file, err = os.Create(filePath)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}

	}
	Context = GetConfig()
}
