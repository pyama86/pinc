# pic
pic is split ~/.ssh/config

# usage
```
$ pic init
=> create dir ~/.ssh/conf.d
   create file ~/.ssh/pic.yml
$ pic merge
 merge configs  [~/.ssh/config and ~/.ssh/conf.d and ~/.ssh/pic.yml]
```

* pic.yml
```yaml
includes:
- ~/src/github.com/org/repos/share_ssh_config
- ~/src/github.com/org/repos/serviceA_ssh_config
- ~/src/github.com/org/repos/serviceB_ssh_config
```

# author
* pyama86
