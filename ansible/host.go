package ansible

import "fmt"

type AnsibleHost struct {
	group      string
	name       string
	ssh_host   string
	ssh_user   string
	ssh_key    string
	bastion_ip string
}

func NewAnsibleHost(group, name, ssh_host, ssh_user, ssh_key, bastion_ip string) AnsibleHost {
	return AnsibleHost{
		group:      group,
		name:       name,
		ssh_host:   ssh_host,
		ssh_user:   ssh_user,
		ssh_key:    ssh_key,
		bastion_ip: bastion_ip,
	}
}

func (host AnsibleHost) String() string {
	var (
		ansible_host    string
		ansible_user    string
		ansible_ssh_key string
	)

	ansible_name := host.name
	no_host_check := "-o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no"
	ansible_ssh_common_args := " ansible_ssh_common_args='" + no_host_check

	if host.ssh_host != "" {
		ansible_host = " ansible_host=" + host.ssh_host
	}
	if host.ssh_user != "" {
		ansible_user = " ansible_user=" + host.ssh_user
	}
	if host.ssh_key != "" {
		ansible_ssh_key = " ansible_ssh_private_key_file=" + host.ssh_key
	}
	if host.bastion_ip != "" {
		ansible_ssh_common_args += " -o ProxyCommand=\"ssh " + no_host_check + " -W %h:%p -i " + host.ssh_key + " " + host.bastion_ip + "\""
	}
	ansible_ssh_common_args += "'"
	content := fmt.Sprintf("%s%s%s%s%s\n", ansible_name, ansible_host, ansible_user, ansible_ssh_key, ansible_ssh_common_args)
	return content
}
