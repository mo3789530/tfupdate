package terraform

import (
	"encoding/json"
	"fmt"
	"log"
)

type ResourceChange struct {
	Address string `json:"address"`
	Type    string `json:"type"`
	Change  struct {
		Actions []string `json:"actions"`
	} `json:"change"`
}

func PlanParser(plan []byte) {
	// JSONを構造体にパース
	var data struct {
		ResourceChanges []ResourceChange `json:"resource_changes"`
	}
	if err := json.Unmarshal(plan, &data); err != nil {
		log.Panic(err)
	}

	for _, rc := range data.ResourceChanges {
		fmt.Printf("%s %s: %v\n", rc.Type, rc.Address, rc.Change.Actions)
	}
}


