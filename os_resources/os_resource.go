package os_resources

type os_resource struct {
	os_type    string
	attributes map[string]string
}

func (osr *os_resource) Init() {
	osr.attributes = make(map[string]string)
}
