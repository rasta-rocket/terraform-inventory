# Terraform Inventory

Terraform Inventory is a tool that generate an Ansible inventory based on a Terraform deployment.

### Download
You can download the released [binaries](https://github.com/rasta-rocket/terraform-inventory/releases)

### Installation

[Go](https://golang.org/) and [dep](https://github.com/golang/dep) are required to build this tool.
Then install terraform inventory from github sources

```sh
$ git clone https://github.com/rasta-rocket/terraform-inventory
$ dep ensure
$ go install
```

### Usage
```sh
$ terraform apply
$ terraform-inventory
? Choose a bastion :
❯ e1b52c07-8a75-4527-b43f-850aaec4c3ef     bastion_0                      84.39.63.203
$ cat hosts
$ ansible -m ping all
```
