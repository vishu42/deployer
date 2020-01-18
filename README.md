# deployer
 
Deployment helper tool hence -> deployer 

Usegae:

Make a binary of this program and paste the binary in /usr/local/bin to make it available globally in your system.

It needs two files to run
  -> secrets.yaml
  -> deployerPayload.yaml
these two files needs to be stored in same directory in which you're running the script.

Template of deployerPayload.yaml can be seen in this repo. 

It needs no arguments. 

It won't push any secrets to helm chart, it will simply append secrets to secrets.yaml file. 
