package cmd

import (
	"fmt"
	"os"

	"github.com/rasta-rocket/terraform-inventory/cmd/options"
	"github.com/rasta-rocket/terraform-inventory/configuration"
	"github.com/rasta-rocket/terraform-inventory/openstack"

	"github.com/spf13/cobra"
)

var opt options.Options

var RootCmd = &cobra.Command{
	Use:   "terraform-inventory [path]",
	Short: "terraform-inventory generates Ansible inventory based on Terraform deployment",
	Run: func(cmd *cobra.Command, args []string) {
		displayVersion(opt.IsVersion)
		configuration.Conf.Init(args, opt)
		run(configuration.Conf)
	},
}

func run(conf configuration.Configuration) {
	tfstate := conf.TfState
	inventory := conf.OutputFile
	bastion := opt.Bastion

	set := openstack.NewSet(tfstate)
	openstack.ToInventory(set, inventory, bastion)
}

func displayVersion(isVersion bool) {
	if isVersion {
		fmt.Printf("terraform-inventory version %s\n", options.VERSION)
		os.Exit(0)
	}
}

func init() {
	RootCmd.Flags().BoolVarP(&opt.IsVersion, "version", "v", false, "version")
	RootCmd.Flags().StringVarP(&opt.TfState, "tfstate", "t", options.DEFAULT_TFSTATE, "tfstate file path")
	RootCmd.Flags().StringVarP(&opt.OutputFile, "output", "o", options.DEFAULT_OUTPUT, "ansible inventory file path to output")
	RootCmd.Flags().StringVarP(&opt.SshUser, "ssh-user", "u", options.DEFAULT_SSH_USER, "ansible ssh user")
	RootCmd.Flags().StringVarP(&opt.SshKey, "ssh-key", "k", options.DEFAULT_SSH_KEY, "ansible ssh key")
	RootCmd.Flags().StringVarP(&opt.Bastion, "bastion", "b", options.DEFAULT_BASTION, "force the IP address of the bastion you want to use")
}
