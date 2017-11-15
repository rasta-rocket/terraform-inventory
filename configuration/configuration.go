package configuration

import (
	"os"
)

type Configuration struct {
	Tfstate    string
	OutputFile string
}

func (conf *Configuration) Init() {
	if isLocalTfstateExist() {
		conf.Tfstate = DEFAULT_TFSTATE_LOCATION
	}
	conf.OutputFile = DEFAULT_OUTPUT_FILE
}

func isLocalTfstateExist() bool {
	_, err := os.Stat(DEFAULT_TFSTATE_LOCATION)
	return err == nil
}

var Conf Configuration
