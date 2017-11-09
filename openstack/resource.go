package openstack

type os_resource struct {
	os_type    string
	attributes map[string]string
}

func (osr *os_resource) Init() {
	osr.attributes = make(map[string]string)
}

func (osr os_resource) GetId() string {
	return osr.attributes["id"]
}

func (osr os_resource) GetComputeName() (name string) {
	if osr.IsCompute() {
		name = osr.attributes["name"]
	}
	return name
}

func (osr os_resource) GetComputeIp() (ip string) {
	if osr.IsCompute() {
		ip = osr.attributes["access_ip_v4"]
	}
	return ip
}

func (osr os_resource) GetComputeKey() (ip string) {
	if osr.IsCompute() {
		ip = osr.attributes["key_pair"]
	}
	return ip
}

func (osr os_resource) IsCompute() bool {
	return osr.os_type == COMPUTE_TYPE
}

func (osr os_resource) IsFloatingAssoc() bool {
	return osr.os_type == FLOATING_ASSOC_TYPE
}
