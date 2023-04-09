package terraform

import (
	"encoding/json"
	"fmt"
	"log"
)

type State struct {
	Resources []struct {
		Type      string                   `json:"type"`
		Name      string                   `json:"name"`
		Provider  string                   `json:"provider"`
		Instances []map[string]interface{} `json:"instances"`
	} `json:"resources"`
}

func StateParser(state []byte) {
	var stateJson State
	err := json.Unmarshal(state, &stateJson)
	if err != nil {
		log.Printf("err unmarshal %v \n", err)
	}

	for _, resource := range stateJson.Resources {
		fmt.Printf("[%s.%s]\n", resource.Type, resource.Name)
		for _, instance := range resource.Instances {
			for key, value := range instance {
				fmt.Printf("%s = %v\n", key, value)
			}
		}
	}
}
