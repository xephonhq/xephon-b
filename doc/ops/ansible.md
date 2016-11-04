# Ansible

- [Official Doc](http://docs.ansible.com/)

## Install

From the source

- `git clone git://github.com/ansible/ansible.git --recursive`
- `git submodule update --init --recursive` if you forgot `--recursive` when first clone
- add `source /path/to/ansiblie-repo/hacking/env-setup -q` to your `.bashrc` or `.zshrc`

<!-- - create soft link
  - `ln -s /path/to/ansible /path/to/your/bin/ansible`
  - `ln -s /path/to/ansible-playbook /path/to/bin/ansible-playbook` -->

## Inventory

Define the target machines by group. [doc](http://docs.ansible.com/ansible/intro_inventory.html)

- default location is `/etc/ansible/hosts`
- use `-i` to specify file location
- use `inventory` in `ansible.cfg`, ie: `inventory = /etc/ansible/hosts`, [doc](http://docs.ansible.com/ansible/intro_configuration.html#inventory-file)

<!-- db.com is a bank .... /w\ -->
````
[cassandra]
node[1:11].db.com

[kairosdb]
ts[1:3].db.com
````

so if you want to ping all the cassandra machines `ansible cassandra -m ping`

## Playbooks

- the simplest way is to put everything in one yaml
- for complex project, see https://github.com/ansible/ansible-examples

````
- hosts: all
  become: true # use root
  tasks:
  - name: update repository
    apt:
      update_cache: yes
  - name: install essential tools
    apt:
      name: git
      state: present
````


## With Vagrant

[Vagrant: use ansible for provisioning](https://www.vagrantup.com/docs/provisioning/ansible.html)

````ruby
# Run Ansible from the Vagrant Host
config.vm.provision "ansible" do |ansible|
  ansible.playbook = "site.yml"
end
````

## Errors

- `import Error` need to use `hacking/env-setup`, soft link the binary will fail
- `Shared connection to 127.0.0.1 closed` seems to be my local ssh timeout config problem, but I switched to `ansible_local` to solve that problem
