package ansible

import (
	"fmt"
	"io/ioutil"
)

type Group []AnsibleHost

type Inventory struct {
	name   string
	groups map[string]Group
}

func NewInventory(name string) Inventory {
	return Inventory{
		name:   name,
		groups: make(map[string]Group),
	}
}

func (inventory *Inventory) String() string {
	content := ""
	for group, hosts := range inventory.groups {
		if len(group) != 0 {
			content += fmt.Sprintf("[%s]\n", group)
		}
		for _, host := range hosts {
			content += host.String()
		}
	}
	return content
}

func (inventory *Inventory) Save() {
	content := inventory.String()
	_ = ioutil.WriteFile(inventory.name, []byte(content), FILE_RIGHT)
}

func (inventory *Inventory) AddHost(groups []string, name, ssh_host, ssh_user, ssh_key, bastion_ip string) {
	if len(groups) == 0 {
		groups = []string{""}
	}
	for _, group := range groups {
		host := NewAnsibleHost(group, name, ssh_host, ssh_user, ssh_key, bastion_ip)
		if _, ok := inventory.groups[group]; !ok {
			inventory.groups[group] = Group{}
		}
		inventory.groups[group] = append(inventory.groups[group], host)
	}
}
