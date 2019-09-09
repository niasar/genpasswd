# Genpasswd

Master branch : [![Build Status](https://dev.azure.com/niasar/genpasswd/_apis/build/status/niasar.genpasswd?branchName=master)](https://dev.azure.com/niasar/genpasswd/_build/latest?definitionId=1&branchName=master)
## SHA-512 password hash generator for /etc/shadow
### Why use this insted of mkpasswd?
Because mkpasswd dosen't require password confirmation when generating hash, this makes it difficult use it as hash generator for 
### Usage 
`genpasswd [salt]`
hash will be saved in passwd.hash file

if no salt provided, it will be randomly generated
