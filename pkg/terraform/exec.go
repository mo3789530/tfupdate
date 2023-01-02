package terraform

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
)

const (
	outPaht = "./out"
)

type Exec interface {
	Init(workingDir string) (*tfexec.Terraform, error)
	Plan(tf *tfexec.Terraform) error
	Show(tf *tfexec.Terraform, isJson bool) (string, error)
	Apply(tf *tfexec.Terraform) error
}

type exec struct {
	version string
}

func NewExec(version string) Exec {
	return &exec{
		version: version,
	}
}

func (e *exec) Init(workingDir string) (*tfexec.Terraform, error) {
	installer := &releases.ExactVersion{
		Product: product.Terraform,
		Version: version.Must(version.NewVersion(e.version)),
	}
	execPath, err := installer.Install(context.Background())
	if err != nil {
		log.Printf("error installing terraform: %s", err)
	}

	tf, err := tfexec.NewTerraform(workingDir, execPath)
	if err != nil {
		log.Printf("error new terraform: %s", err)
	}

	return tf, nil
}

func (e *exec) Plan(tf *tfexec.Terraform) error {
	err := tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		log.Printf("error init terraform: %s", err)
	}

	planOptions := []tfexec.PlanOption{
		tfexec.Out(outPaht),
	}

	resule, err := tf.Plan(context.Background(), planOptions...)
	if err != nil {
		log.Printf("error terraform plan: %s", err)
		return err
	}

	if !resule {
		return fmt.Errorf("error exec plan terraform")
	}

	return nil
}

func (e *exec) Show(tf *tfexec.Terraform, isJson bool) (string, error) {

	if isJson {
		result, err := e.showJson(tf)
		if err != err {
			return "", err
		}
		return result, nil
	}

	result, err := e.showPlanText(tf)
	if err != err {
		return "", err
	}
	return result, nil
}

func (e *exec) showJson(tf *tfexec.Terraform) (string, error) {
	show, err := tf.ShowPlanFileRaw(context.Background(), outPaht)
	if err != nil {
		log.Printf("error terraform show: %s", err)
		return "", err
	}
	return show, nil
}

func (e *exec) showPlanText(tf *tfexec.Terraform) (string, error) {
	show, err := tf.ShowPlanFile(context.Background(), outPaht)
	if err != nil {
		log.Printf("error terraform show: %s", err)
	}
	jsonData, err := json.Marshal(show)
	if err != nil {
		log.Printf("error convert json %s", err)
		return "", err
	}
	return string(jsonData), nil
}

func (e *exec) Apply(tf *tfexec.Terraform) error {
	err := tf.Apply(context.Background())

	if err != nil {
		log.Printf("error terraform apply: %s", err)
		return err
	}
	return nil
}
