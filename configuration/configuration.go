package configuration

import (
	"io/ioutil"
	"os"
)

type Configuration struct {
	TfPath     string
	Workspace  string
	Tfstate    string
	OutputFile string
	SshUser    string
	SshKey     string
}

func (conf *Configuration) Init() {
	if isLocalTfstateExist() {
		conf.Tfstate = DEFAULT_TFSTATE_LOCATION
	}
	conf.OutputFile = DEFAULT_OUTPUT_FILE
}

func isPathExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsExist(err)
}

func getCurrentWorkspace(path string) string {
	workspace := "default"
	defaultEnvironmentPath := ".terraform/environment"
	environmentPath := path + "/" + defaultEnvironmentPath

	if !isPathExist(environmentPath) {
		return workspace
	}

	byteContent, _ := ioutil.ReadFile(environmentPath)
	return string(byteContent)
}

func isLocalTfstateExist() bool {
	_, err := os.Stat(DEFAULT_TFSTATE_LOCATION)
	return err == nil
}

var Conf Configuration
