package openstack

import (
	"strings"
)

type Resource struct {
	os_type    string
	attributes map[string]string
}

func (osr *Resource) Init() {
	osr.attributes = make(map[string]string)
}

func (osr Resource) GetId() string {
	return osr.attributes["id"]
}

func (osr Resource) GetComputeName() (name string) {
	if osr.IsCompute() {
		name = osr.attributes["name"]
	}
	return name
}

func (osr Resource) GetComputeIp() (ip string) {
	if osr.IsCompute() {
		ip = osr.attributes["access_ip_v4"]
	}
	return ip
}

func (osr Resource) GetComputeKey() (key string) {
	if osr.IsCompute() {
		key = osr.attributes["key_pair"]
	}
	return key
}

func (osr Resource) GetComputeAnsibleGroups() (groups []string) {
	if osr.IsCompute() {
		tmp := strings.Replace(osr.attributes["metadata.ansible_group"], GROUP_IFS, " ", -1)
		groups = strings.Fields(tmp)
	}
	return groups
}

func (osr Resource) IsCompute() bool {
	return osr.os_type == COMPUTE_TYPE
}

func (osr Resource) IsFloatingAssoc() bool {
	return osr.os_type == FLOATING_ASSOC_TYPE
}
