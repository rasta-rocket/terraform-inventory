package openstack

import (
	"github.com/rasta-rocket/terraform-inventory/ansible"
)

type AnsibleHosts []ansible.AnsibleHost

func ToInventory(name string, set OS_Set) {
	hosts := AnsibleHosts{}
	computes := set.GetComputes()

	for _, compute := range computes {
		name := compute.GetComputeName()
		ssh_host := compute.GetComputeIp()
		ssh_user := "cloud"
		ssh_key := compute.GetComputeKey()
		has_public_ip := false

		host := ansible.NewAnsibleHost(name, ssh_host, ssh_user, ssh_key, has_public_ip)
		hosts = append(hosts, host)
	}

	inventory := ansible.NewAnsibleInventory(name)
	for _, host := range hosts {
		inventory.AddHost(host.Name(), host.SshHost(), host.SshUser(), host.SshKey())
	}
	inventory.Save()
}
