package terraform

import (
	"bufio"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"strings"

	tfjson "github.com/hashicorp/terraform-json"
)

//go:embed templates/*
var templateFS embed.FS

type data struct {
	Heading     string
	Description string
	Outputs     map[string]*tfjson.StateOutput
}

type TfState interface {
	CreateDoc(string, *string) error
}

type tfstate struct {
	format string
}

func NewTfState(format string) TfState {
	return &tfstate{
		format: format,
	}
}

func getValue(output *tfjson.StateOutput) string {
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

func prettyPrint(output tfjson.StateOutput) template.HTML {
	if output.Sensitive {
		return template.HTML("sensitive")
	}
	pretty, err := json.MarshalIndent(output.Value, "", " ")
	if err != nil {
		log.Fatalf("error %s", err)
	}

	return template.HTML(string(pretty))
}

func dataType(output tfjson.StateOutput) string {
	return fmt.Sprintf("%T", output.Value)
}

func (t *tfstate) CreateDoc(path string, jsonData *string) error {
	var stateReader io.Reader = bufio.NewReader(os.Stdin)
	if jsonData != nil && *jsonData != "" {
		stateReader = strings.NewReader(*jsonData)
	}
	if path != "" {
		b, err := os.ReadFile(path)
		if err != nil {
			log.Printf("error read state file %s \n", err)
			return err
		}
		stateReader = strings.NewReader(string(b))
	}

	fmt.Println(stateReader)

	var state *tfjson.State

	if err := json.NewDecoder(stateReader).Decode(&state); err != nil {
		log.Printf("error %s", err)
		return err
	}

	tmpl, err := getTemplate("md")
	if err != nil {
		log.Printf("error get template %s", err)
		return err
	}

	n, err := template.New("md.tmpl").Funcs(template.FuncMap{
		"value":       getValue,
		"dataType":    dataType,
		"prettyPrint": prettyPrint,
	}).ParseFS(templateFS, tmpl)

	if err != nil {
		log.Printf("error new template %s", err)
		return err
	}

	// fmt.Println(state.Values.RootModule.Resources)
	fmt.Println("aaaa")
	for _, v := range state.Values.RootModule.Resources {
		a := fmt.Sprintf("%v", v.AttributeValues)
		fmt.Println(a)
		b, _ := json.Marshal(v.AttributeValues)
		fmt.Println(string(b))
		for _, v2 := range v.AttributeValues {
			fmt.Println(v2)
		}
	}

	outputs := map[string]*tfjson.StateOutput{}
	if state.Values != nil {
		outputs = state.Values.Outputs
	}

	err = n.Execute(os.Stdout, data{
		Outputs:     outputs,
		Description: "test",
	})

	if err != nil {
		log.Printf("error %s", err)
		return err
	}

	return nil

}
