package os_resources

import (
	"fmt"
	"github.com/rasta-rocket/terraform-inventory/inventory"
)

type OS_Inventory struct {
	hosts []inventory.AnsibleHost
}

func (os_inventory *OS_Inventory) Init(os_set OS_Set) {
	computes := os_set.GetComputes()
	//floatings := os_set.GetFloatingAssoc()
	//inventory.hosts = []AnsibleHost{}

	for _, compute := range computes {
		name := compute.GetComputeName()
		ssh_host := compute.GetComputeIp()
		ssh_user := "cloud"
		ssh_key := compute.GetComputeKey()
		has_public_ip := false

		host := inventory.AnsibleHost{}
		host.Init(name, ssh_host, ssh_user, ssh_key, has_public_ip)
		os_inventory.hosts = append(os_inventory.hosts, host)
	}
}

func (os_inventory OS_Inventory) GenInventory(name string) {
	tf_inv := inventory.Inventory{}
	tf_inv.Init(name)
	fmt.Printf("%+v\n", os_inventory.hosts)
	for _, host := range os_inventory.hosts {
		tf_inv.AddHost(host.Name, host.Ssh_host, host.Ssh_user, host.Ssh_key)
	}
	tf_inv.Save()
}
