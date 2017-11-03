package os_resources

import (
	"encoding/json"
	"io/ioutil"
)

type OS_Set []os_resource

func (os_set *OS_Set) Init(tf_file_name string) {
	var tfstate map[string]interface{}
	content, _ := ioutil.ReadFile(tf_file_name)
	_ = json.Unmarshal(content, &tfstate)
	*os_set = append(*os_set, parse_tfstate(tfstate)...)
}

func parse_tfstate(tfstate map[string]interface{}) (os_set OS_Set) {
	modules := tfstate["modules"].([]interface{})
	for _, module := range modules {
		osr_list := parse_module(module.(map[string]interface{}))
		os_set = append(os_set, osr_list...)
	}
	return os_set
}

func parse_module(module map[string]interface{}) (osr_list []os_resource) {
	resources := module["resources"].(map[string]interface{})
	for _, res := range resources {
		osr := parse_resource(res.(map[string]interface{}))
		osr_list = append(osr_list, osr)
	}
	return osr_list
}

func parse_resource(resource map[string]interface{}) (osr os_resource) {
	osr.Init()
	osr.os_type = resource["type"].(string)
	res_primary := resource["primary"].(map[string]interface{})
	res_attr := res_primary["attributes"].(map[string]interface{})
	for k, v := range res_attr {
		osr.attributes[k] = v.(string)
	}
	return osr
}

func (os_set OS_Set) GetComputes() (computes OS_Set) {
	for _, osr := range os_set {
		if osr.IsCompute() {
			computes = append(computes, osr)
		}
	}
	return computes
}

func (os_set OS_Set) GetFloatingAssoc() (floatings OS_Set) {
	for _, osr := range os_set {
		if osr.IsFloatingAssoc() {
			floatings = append(floatings, osr)
		}
	}
	return floatings
}
