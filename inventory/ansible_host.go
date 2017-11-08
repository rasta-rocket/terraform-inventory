package inventory

type AnsibleHost struct {
	Name          string
	Ssh_host      string
	Ssh_user      string
	Ssh_key       string
	Has_public_ip bool
}

func (host *AnsibleHost) Init(name, ssh_host, ssh_user, ssh_key string, has_public_ip bool) {
	host.Name = name
	host.Ssh_host = ssh_host
	host.Ssh_user = ssh_user
	host.Ssh_key = ssh_key
	host.Has_public_ip = has_public_ip
}
