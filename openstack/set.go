package openstack

import (
	"encoding/json"
	"io/ioutil"
)

type Set []Resource

func NewSet(tf_file_name string) (set Set) {
	set.Init(tf_file_name)
	return set
}

func (set *Set) Init(tf_file_name string) {
	var tfstate map[string]interface{}
	content, _ := ioutil.ReadFile(tf_file_name)
	_ = json.Unmarshal(content, &tfstate)
	*set = append(*set, parse_tfstate(tfstate)...)
}

func parse_tfstate(tfstate map[string]interface{}) (set Set) {
	modules := tfstate["modules"].([]interface{})
	for _, module := range modules {
		osr_list := parse_module(module.(map[string]interface{}))
		set = append(set, osr_list...)
	}
	return set
}

func parse_module(module map[string]interface{}) (osr_list []Resource) {
	resources := module["resources"].(map[string]interface{})
	for _, res := range resources {
		osr := parse_resource(res.(map[string]interface{}))
		osr_list = append(osr_list, osr)
	}
	return osr_list
}

func parse_resource(resource map[string]interface{}) (osr Resource) {
	osr.Init()
	osr.os_type = resource["type"].(string)
	res_primary := resource["primary"].(map[string]interface{})
	res_attr := res_primary["attributes"].(map[string]interface{})
	for k, v := range res_attr {
		osr.attributes[k] = v.(string)
	}
	return osr
}

func (set Set) GetComputes() (computes Set) {
	for _, osr := range set {
		if osr.IsCompute() {
			computes = append(computes, osr)
		}
	}
	return computes
}

func (set Set) GetFloatingAssoc() (floatings Set) {
	for _, osr := range set {
		if osr.IsFloatingAssoc() {
			floatings = append(floatings, osr)
		}
	}
	return floatings
}
