# [WIP]pic
## Description
pic is split ~/.ssh/config

## Usage
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

## Install
### homebrew
```
$ brew tap pyama86/pic
$ brew install pic
```
### go get

```bash
$ go get -d github.com/pyama86/pic
```

## Contribution

1. Fork ([https://github.com/pyama86/pic/fork](https://github.com/pyama86/pic/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request
