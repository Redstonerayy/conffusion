# Conffusion
Sometimes its confusing to have many config files or many operating system installations. Sometimes, after installing a new system, you
want to configure it the same as another system. Focused on linux this program wants to enable backing up config files
and bootstrapping new systems with them. 

# Todo
* make bootstrapping possible and test in a Virtual Maschine
* make it possible to choose if u want to store the config from the bootstrap file
* change pkg file name to package manager name

# Usage
Create a .conffusion file in your homefolder
There different variables can be defined.
Example Contents of $HOME/.conffusion
```
CONFIGFOLDER /my/own/path
```

CONFIGFOLDER defines the folder where conffusion searched for `config.json` and `vars.txt`.