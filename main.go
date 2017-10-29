package main

import (
	"fmt"
	"github.com/rasta-rocket/terraform-inventory/os_resources"
)

/*
   .modules[1].resources["openstack_compute_instance_v2.instance"].primary.attributes.access_ip_v4
   .modules[1].resources["openstack_compute_instance_v2.instance"].type == openstack_compute_instance_v2
*/

func main() {
	var tf_file_name string = "./terraform.tfstate"
	var os_res os_resources.OS_Set
	os_res.Init(tf_file_name)
	fmt.Printf("%+v\n", os_res)
}
