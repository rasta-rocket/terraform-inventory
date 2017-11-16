package cmd

import (
	"github.com/rasta-rocket/terraform-inventory/configuration"
	"github.com/rasta-rocket/terraform-inventory/openstack"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "terraform-inventory",
	Short: "terraform-inventory generates Ansible inventory based on Terraform deployment",
	Run: func(cmd *cobra.Command, args []string) {
		configuration.Conf.Init()
		run()
	},
}

func run() {
	tfstate := configuration.Conf.Tfstate
	inventory := configuration.Conf.OutputFile
	var set openstack.Set
	set.Init(tfstate)
	openstack.ToInventory(inventory, set)
}
