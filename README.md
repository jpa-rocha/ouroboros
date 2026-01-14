# Use steps

1. use sops to create a secrets.yaml file

```
vpn:
    ipsec:
        username: **\***
        password: **\***
    id:
        username: **\***
        password: **\***
```

2. run the ""build"" command
3. add folder to path, in .bashrc add line: export PATH=$PATH:$HOME/"name of your directory"
