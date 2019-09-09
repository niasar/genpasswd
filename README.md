# Genpasswd

Master branch : [![Build Status](https://dev.azure.com/niasar/genpasswd/_apis/build/status/niasar.genpasswd?branchName=master)](https://dev.azure.com/niasar/genpasswd/_build/latest?definitionId=1&branchName=master)
## SHA-512 password hash generator for /etc/shadow
### Why use this instead of mkpasswd?
Because mkpasswd dosen't require password confirmation when generating hash, this makes it difficult to use as hash generator for interactive password rotaing system in group of hosts using Ansible
### Usage 

`genpasswd [--stdout] [--filename file.name] [salt]`

By default hash will be saved in passwd.hash file or it may be saved in different file(if --filename specified) or printed to stdout (if --stdout specified)

Provided salt might be 8 to 16 characters long, if no salt provided, 16 chars long random salt will be generated
