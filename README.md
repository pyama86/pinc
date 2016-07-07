# pinc
## Description
pinc is split ~/.ssh/config

## Usage
```
$ pinc init
=> create dir ~/.ssh/conf.d
   create file ~/.ssh/pinc.yml
$ pinc gen
 merge configs  [~/.ssh/config and ~/.ssh/conf.d and ~/.ssh/pinc.yml]
```

* pinc.yml
```yaml
includes:
- ~/src/github.com/org/repos/share_ssh_config
- ~/src/github.com/org/repos/serviceA_ssh_config
- ~/src/github.com/org/repos/serviceB_ssh_config
```

## Install
### homebrew
```
$ brew tap pyama86/pinc
$ brew install pinc
```
### go get

```bash
$ go get -d github.com/pyama86/pinc
```

## Contribution

1. Fork ([https://github.com/pyama86/pinc/fork](https://github.com/pyama86/pinc/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request
