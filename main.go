package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/nxvishal/yamlparser"
)

func main() {
	secrets := yamlparser.Parse("./secrets.yaml")["secrets"].(map[interface{}]interface{})

	payload := yamlparser.Parse("./deployerPayload.yaml")
	targetTeamEnv := payload["targetTeamEnv"].(string)
	if len(targetTeamEnv) == 0 {
		log.Fatal("Empty targetTeamEnv")
	}
	pickTeamEnv := payload["pickTeamEnv"].(string)
	if len(targetTeamEnv) == 0 {
		log.Fatal("Empty targetTeamEnv")
	}
	fmt.Println("copying", pickTeamEnv, "secrets to", targetTeamEnv)
	overrideKeys := payload["overrideKeysWith"].(map[interface{}]interface{})

	targetSecretsMap := make(map[string]string, 15)
	for k, v := range secrets {
		if strings.Contains(k.(string), pickTeamEnv) {
			targetKey := strings.Replace(k.(string), pickTeamEnv, targetTeamEnv, 1)
			targetSecretsMap[targetKey] = v.(string)
		}
	}

	for k, v := range overrideKeys {
		targetSecretsMap[strings.Join([]string{targetTeamEnv, k.(string)}, "")] = v.(string)
	}
	result, err := yaml.Marshal(targetSecretsMap)
	if err != nil {
		log.Fatal(err)
	}

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("./secrets.yaml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write(result); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%v", targetSecretsMap)
}
