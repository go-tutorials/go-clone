package main

import (
	"encoding/json"
	"fmt"
	"go-clone/internal/clone"
	"go-clone/internal/config"
	"go-clone/internal/model"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func main() {
	var rootConfig config.RootConfig
	rootConfig = readFile()
	fmt.Println(rootConfig)

	var userConfig model.User
	json.Unmarshal([]byte(rootConfig.Config), &userConfig)
	fmt.Println(userConfig)

	user := clone.Clone(userConfig).(*model.User)
	fmt.Println(user)
	fmt.Println(*user.Name)
	fmt.Println(*user.Age)
	if user.Sex != nil {
		fmt.Println(*user.Sex)
	}
	fmt.Println(*user.Amount)


	err := yaml.Unmarshal([]byte(rootConfig.ConfigRef1), &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user)

}

func readFile() config.RootConfig {
	yamlFile, err := ioutil.ReadFile("configs/config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	var c config.RootConfig
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}