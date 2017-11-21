package openstack

import (
	"github.com/rasta-rocket/terraform-inventory/ansible"
	"github.com/rasta-rocket/terraform-inventory/configuration"
)

type AnsibleHosts []ansible.AnsibleHost

func ToInventory(set Set, name string) {
	computes := set.GetComputes()

	inventory := ansible.NewInventory(name)
	for _, compute := range computes {
		name := compute.GetComputeName()
		groups := compute.GetComputeAnsibleGroups()
		ssh_host := compute.GetComputeIp()
		ssh_user := configuration.Conf.SshUser
		ssh_key := configuration.Conf.SshKey
		bastion_ip := ""

		inventory.AddHost(groups, name, ssh_host, ssh_user, ssh_key, bastion_ip)
	}

	inventory.Save()
}
