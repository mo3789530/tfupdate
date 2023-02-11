package terraform

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"

	tfjson "github.com/hashicorp/terraform-json"
)

type TfState interface {
}

type tfstate struct {
	path string
}

func NewTfState() TfState {
	return &tfstate{}
}

func (t *tfstate) getValue(output tfjson.StateOutput) string {
	if output.Sensitive {
		return "sensitive value"
	}
	return fmt.Sprintf("%v", output.Value)
}

func getTemplate(format string) (string, error) {
	switch format {
	case "md":
		return "templates/md.tmpl", nil
	default:
		return "", fmt.Errorf("%s is not supported output format", format)
	}
}

func (t *tfstate) prettyPrint(output tfjson.StateOutput) template.HTML {
	if output.Sensitive {
		return template.HTML("sensitive")
	}
	pretty, err := json.MarshalIndent(output.Value, "", " ")
	if err != nil {
		log.Fatalf("error %s", err)
	}

	return template.HTML(string(pretty))
}
