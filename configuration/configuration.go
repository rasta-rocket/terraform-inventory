package configuration

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/rasta-rocket/terraform-inventory/cmd/options"
)

const DEFAULT_TF_PATH = "."
const DEFAULT_TFSTATE_LOCATION = "terraform.tfstate"
const DEFAULT_WORKSPACE = "default"

type Configuration struct {
	TfPath     string
	Workspace  string
	TfState    string
	OutputFile string
	SshUser    string
	SshKey     string
}

func (conf *Configuration) Init(args []string, opt options.Options) {
	conf.TfPath = parseArgs(args)
	checkTfPathIsInitialized(conf.TfPath)
	conf.Workspace = getCurrentWorkspace(conf.TfPath)
	conf.TfState = getTfState(opt.TfState, conf.TfPath, conf.Workspace)
	conf.SshUser = opt.SshUser
	conf.SshKey = opt.SshKey
	conf.OutputFile = opt.OutputFile
}

func parseArgs(args []string) (tfPath string) {
	if len(args) == 0 {
		tfPath, _ = filepath.Abs(DEFAULT_TF_PATH)
	} else if len(args) == 1 {
		tfPath, _ = filepath.Abs(args[0])
		file_info, err := os.Stat(tfPath)
		if os.IsNotExist(err) {
			fmt.Printf("Error: %s does not exist\n", tfPath)
			os.Exit(1)
		} else if !file_info.Mode().IsDir() {
			fmt.Printf("Error: %s is not a directory\n", tfPath)
			os.Exit(1)
		}
	} else {
		fmt.Println("Error: Only one argument is allowed")
		os.Exit(1)
	}
	return tfPath
}

func checkTfPathIsInitialized(tfPath string) {
	path := fmt.Sprintf("%s/%s", tfPath, ".terraform")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("%s must be initialized\nRun: terraform init\n", tfPath)
		os.Exit(1)
	}
}

func getCurrentWorkspace(tfPath string) string {
	defaultEnvironmentPath := ".terraform/environment"
	environmentPath := fmt.Sprintf("%s/%s", tfPath, defaultEnvironmentPath)

	if _, err := os.Stat(environmentPath); os.IsNotExist(err) {
		return DEFAULT_WORKSPACE
	}

	byteContent, _ := ioutil.ReadFile(environmentPath)
	return string(byteContent)
}

func getTfState(tfStateOption string, tfPath string, workspace string) (tfState string) {
	if len(tfStateOption) > 0 {
		tfState, _ = filepath.Abs(tfStateOption)
		if _, err := os.Stat(tfState); os.IsNotExist(err) {
			fmt.Printf("Error: --tfstate %s does not exist\n", tfState)
			os.Exit(1)
		}
		return tfState
	}

	if workspace == DEFAULT_WORKSPACE {
		tfState = fmt.Sprintf("%s/%s", tfPath, DEFAULT_TFSTATE_LOCATION)
	} else {
		tfState = fmt.Sprintf("%s/%s/%s/%s", tfPath, "terraform.tfstate.d", workspace, DEFAULT_TFSTATE_LOCATION)
	}

	if _, err := os.Stat(tfState); os.IsNotExist(err) {
		fmt.Printf("Error: %s does not exist\n", tfState)
		os.Exit(1)
	}
	return tfState
}

var Conf Configuration
