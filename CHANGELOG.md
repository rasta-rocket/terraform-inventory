# Change Log

All notable changes to this project will be documented in this file.

## [1.2.0] - 2018-08-14
### Added
- Add bastion option
- Update constraint depedencies to latest stable version
- Add ssh_user in the proxy command

## [1.1.0] - 2018-03-20
### Added
- Make the inventory completely independent from the ssh configuration of the user:
  - Add SSH_KEY in proxy command.
  - Add UserKnownHostsFile and StrictHostKeyChecking option.
- Can take as argument the terraform directory to manage.
- Manage the terraform workspaces.
- Output basic error logs regarding path existence.

## [1.0.0] - 2017-12-11
### Added
- First version of terrafrom-inventory.
- Go depedencies mangement system added (dep).
- Options management (cobra).
- Support openstack provider.
- Generate Ansible inventory file.
- Generate ssh proxy command based on bastion.
- Generate Ansible group based on openstack virtual machine metadata.
