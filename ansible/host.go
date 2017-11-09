package ansible

type AnsibleHost struct {
	name          string
	ssh_host      string
	ssh_user      string
	ssh_key       string
	has_public_ip bool
}

func NewAnsibleHost(name, ssh_host, ssh_user, ssh_key string, has_public_ip bool) AnsibleHost {
	return AnsibleHost{
		name:          name,
		ssh_host:      ssh_host,
		ssh_user:      ssh_user,
		ssh_key:       ssh_key,
		has_public_ip: has_public_ip,
	}
}

func (host AnsibleHost) Name() string {
	return host.name
}

func (host AnsibleHost) SshHost() string {
	return host.ssh_host
}

func (host AnsibleHost) SshUser() string {
	return host.ssh_user
}

func (host AnsibleHost) SshKey() string {
	return host.ssh_key
}

func (host AnsibleHost) HasPublicIp() bool {
	return host.has_public_ip
}
