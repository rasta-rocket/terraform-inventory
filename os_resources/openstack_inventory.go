package os_resources

import (
	"fmt"
	"github.com/rasta-rocket/terraform-inventory/inventory"
)

type AnsibleHost struct {
	name          string
	ssh_host      string
	ssh_user      string
	ssh_key       string
	has_public_ip bool
}

func (host *AnsibleHost) Init(name, ssh_host, ssh_user, ssh_key string, has_public_ip bool) {
	host.name = name
	host.ssh_host = ssh_host
	host.ssh_user = ssh_user
	host.ssh_key = ssh_key
	has_public_ip = has_public_ip
}

type OS_Inventory struct {
	hosts []AnsibleHost
}

func (inventory *OS_Inventory) Init(os_set OS_Set) {
	computes := os_set.GetComputes()
	//floatings := os_set.GetFloatingAssoc()
	//inventory.hosts = []AnsibleHost{}

	for _, compute := range computes {
		fmt.Println("Get anisble Hosts")
		name := compute.GetComputeName()
		ssh_host := compute.GetComputeIp()
		ssh_user := "cloud"
		ssh_key := compute.GetComputeKey()
		has_public_ip := false

		host := AnsibleHost{}
		host.Init(name, ssh_host, ssh_user, ssh_key, has_public_ip)
		inventory.hosts = append(inventory.hosts, host)
	}
}

func (os_inventory OS_Inventory) GenInventory(name string) {
	tf_inv := inventory.Inventory{}
	tf_inv.Init(name)
	fmt.Println("GenInventory")
	fmt.Printf("%+v\n", os_inventory.hosts)
	for _, host := range os_inventory.hosts {
		fmt.Printf("host.name = %s \n", host.name)
		tf_inv.AddHost(host.name, host.ssh_host, host.ssh_user, host.ssh_key)
	}
	tf_inv.Save()
}
