package openstack

import (
	"fmt"
	"strings"

	"github.com/rasta-rocket/terraform-inventory/ansible"
	"github.com/rasta-rocket/terraform-inventory/configuration"
	"gopkg.in/AlecAivazis/survey.v1"
)

func ToInventory(set Set, name string, bastion string) {
	computes := set.GetComputes()
	floatings := set.GetFloatingAssoc()
	bst_id, bst_ip := "", bastion

	if len(bastion) == 0 {
		bst_id, _, bst_ip = getBastion(floatings, computes)
	}

	if len(bst_ip) > 0 {
		fmt.Printf("Inventory file is configured with %s as bastion (ssh-proxy)\n", bst_ip)
	}

	inventory := ansible.NewInventory(name)
	for _, compute := range computes {
		name := compute.GetComputeName()
		groups := compute.GetComputeAnsibleGroups()
		ssh_host := compute.GetComputeIp()
		ssh_user := configuration.Conf.SshUser
		ssh_key := configuration.Conf.SshKey
		bastion_ip := bst_ip
		if strings.Compare(bst_id, compute.GetId()) == 0 {
			ssh_host = bst_ip
			bastion_ip = ""
		}

		inventory.AddHost(groups, name, ssh_host, ssh_user, ssh_key, bastion_ip)
	}

	inventory.Save()
}

func getBastion(floatings Set, computes Set) (string, string, string) {
	bastion_list := []string{}

	// Create list of potential bastion
	for _, floating := range floatings {
		for _, compute := range computes {
			fip_id := floating.GetFloatingInstanceId()
			compute_id := compute.GetId()
			if strings.Compare(fip_id, compute_id) == 0 {
				bst := fmt.Sprintf("%-40s %-30s %s", compute_id, compute.GetComputeName(), floating.GetFloatingIp())
				bastion_list = append(bastion_list, bst)
			}
		}
	}

	if len(bastion_list) == 0 {
		return "", "", ""
	}

	// Prompt the user to get the good bastion
	choice := ""
	prompt := &survey.Select{
		Message: "Choose a bastion :",
		Options: bastion_list,
	}
	survey.AskOne(prompt, &choice, nil)

	// Split the fields of the strings and return the proper values
	fields := strings.Fields(choice)
	return fields[0], fields[1], fields[2]
}
