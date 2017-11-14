package cmd

import (
	"github.com/rasta-rocket/terraform-inventory/openstack"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "terraform-inventory",
	Short: "terraform-inventory generates Ansible inventory based on Terraform deployment",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	tfstate := "./terraform.tfstate"
	inventory := "./host"
	var set openstack.OS_Set
	set.Init(tfstate)
	openstack.ToInventory(inventory, set)
}
