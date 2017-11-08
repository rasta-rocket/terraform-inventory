package main

import (
	//	"fmt"
	//	"github.com/rasta-rocket/terraform-inventory/inventory"
	"github.com/rasta-rocket/terraform-inventory/os_resources"
)

func main() {
	var tf_file_name string = "./terraform.tfstate"
	var os_res os_resources.OS_Set
	var os_inv os_resources.OS_Inventory
	os_res.Init(tf_file_name)
	os_inv.Init(os_res)
	os_inv.GenInventory("test.ini")
}
