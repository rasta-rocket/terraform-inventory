package main

import (
	"github.com/rasta-rocket/terraform-inventory/openstack"
)

func main() {
	tfstate := "./terraform.tfstate"
	inventory := "./host"
	var set openstack.OS_Set
	set.Init(tfstate)
	openstack.ToInventory(inventory, set)
}
