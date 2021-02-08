package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	ConfigProfile    string
	ApplicationName  string
	ConfigServer     string
	AppConfiguration Configuration
)

func LoadConfiguration() (Configuration, error) {
	var err error
	fmt.Println("Loading configuration ['" + ConfigProfile + "'] from server.")
	AppConfiguration, err = loadConfigFromServer(AppConfiguration)
	if err != nil {
		fmt.Println("Error loading configuration. " + err.Error())
		os.Exit(1)
	}
	return AppConfiguration, nil
}

func loadConfigFromServer(configuration Configuration) (Configuration, error) {
	configFile := ConfigServer + ApplicationName + "-" + ConfigProfile + ".json"
	configFileResponse, err := http.Get(configFile)
	if err != nil {
		fmt.Println("Error downloading file [\"" + ConfigProfile + "\" for \"" + ApplicationName + "\"] " + err.Error())
		return configuration, err
	}
	defer configFileResponse.Body.Close()
	byteValue, err := ioutil.ReadAll(configFileResponse.Body)
	if err != nil {
		fmt.Println("Error reading configuration from Config Server " + err.Error())
		return configuration, err
	}
	er := json.Unmarshal(byteValue, &configuration)
	if er != nil {
		fmt.Println("Error parsing json configuration file ")
		return configuration, err
	}
	fmt.Println(configuration)
	return configuration, nil
}
