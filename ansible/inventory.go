package ansible

import (
	"fmt"
	"io/ioutil"
)

type Inventory struct {
	name    string
	content string
}

func NewAnsibleInventory(name string) Inventory {
	return Inventory{
		name:    name,
		content: "",
	}
}

func (inventory *Inventory) AddGroup(groupname string) {
	result := fmt.Sprintf("[%s]\n", groupname)
	inventory.content += result
}

func (inventory *Inventory) AddHost(host_name, host_ip, ssh_user, ssh_key string, ssh_common_args ...string) {
	ansible_host := "ansible_host=" + host_ip
	ansible_user := "ansible_user=" + ssh_user
	ansible_ssh_key := "ansible_ssh_private_key_file=" + ssh_key

	result := fmt.Sprintf("%s %s %s %s", host_name, ansible_host, ansible_user, ansible_ssh_key)

	for _, bastion_ip := range ssh_common_args {
		ansible_ssh_common_args := "ansible_ssh_common_args='-o ProxyCommand=\"ssh -W %h:%p " + bastion_ip + "\"'"
		result = fmt.Sprintf("%s %s", result, ansible_ssh_common_args)
	}
	inventory.content += result + "\n"
}

func (inventory *Inventory) Save() {
	_ = ioutil.WriteFile(inventory.name, []byte(inventory.content), FILE_RIGHT)
}

func Test() {
	inventory := NewAnsibleInventory("inventory.ini")
	inventory.AddGroup("web")
	inventory.AddHost("apache", "10.0.0.3", "cloud", "key")
	inventory.AddHost("mysql", "10.0.0.4", "cloud", "key", "192.168.10.3")
	inventory.Save()
}
