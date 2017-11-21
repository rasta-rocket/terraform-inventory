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
		//_ = ansible.NewConfiguration("./ansible.cfg")
		run()
	},
}

func run() {
	tfstate := configuration.Conf.Tfstate
	inventory := configuration.Conf.OutputFile
	set := openstack.NewSet(tfstate)
	openstack.ToInventory(set, inventory)
}

func init() {
	RootCmd.Flags().StringVarP(&configuration.Conf.Tfstate, "tfstate", "t", configuration.DEFAULT_TFSTATE_LOCATION, "local tfstate file")
	RootCmd.Flags().StringVarP(&configuration.Conf.OutputFile, "output", "o", configuration.DEFAULT_OUTPUT_FILE, "ansible inventory file to output")
	RootCmd.Flags().StringVarP(&configuration.Conf.SshUser, "ssh-user", "u", configuration.DEFAULT_SSH_USER, "ansible ssh user")
	RootCmd.Flags().StringVarP(&configuration.Conf.SshKey, "ssh-key", "k", configuration.DEFAULT_SSH_KEY, "ansible ssh key")
}
