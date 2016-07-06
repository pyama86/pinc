package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
)

var SSH_CONFIG = "/.ssh/config"
var SSH_CONFIG_DIR = "/.ssh/conf.d"
var BACKUP_SSH_CONFIG = "/.ssh/conf.d/base_config"
var PIC_CONFIG = "/.ssh/pic.yml"

var cmdInit = &Command{
	Run:       runInit,
	UsageLine: "init ",
	Short:     "initialize directory and config file",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdInit.Flag.BoolVar(&flagA, "a", false, "")
}

// runInit executes init command and return exit code.
func runInit(args []string) int {
	homeDir, _ := homedir.Dir()

	if _, err := os.Stat(homeDir + SSH_CONFIG); err != nil {
		fmt.Fprintf(os.Stderr, "pic: %s not found\n", homeDir+SSH_CONFIG)
		os.Exit(1)
	}

	if _, err := os.Stat(homeDir + SSH_CONFIG_DIR); err != nil {
		if err := os.MkdirAll(homeDir+SSH_CONFIG_DIR, 0744); err != nil {
			fmt.Fprintf(os.Stderr, "pic: failed to create dir %s:%s\n", homeDir+SSH_CONFIG_DIR, err)
			os.Exit(1)
		} else {
			fmt.Printf("pic: create directory: %s\n", homeDir+SSH_CONFIG_DIR)
		}
	}

	// backup original config file to conf.d/base_config
	if _, err := os.Stat(homeDir + BACKUP_SSH_CONFIG); err != nil {
		_ = os.Link(homeDir+SSH_CONFIG, homeDir+BACKUP_SSH_CONFIG)
		fmt.Printf("pic: copy config file: %s => %s\n", homeDir+SSH_CONFIG, homeDir+BACKUP_SSH_CONFIG)
	}

	if _, err := os.Stat(homeDir + PIC_CONFIG); err != nil {
		f, err := os.Create(homeDir + PIC_CONFIG)
		if err != nil {
			fmt.Printf("pic: failed to create file: %s:%s\n", homeDir+PIC_CONFIG, err)
			os.Exit(1)
		}
		f.WriteString("include:\n#  - /path/to/config\n")
		f.Close()
		fmt.Printf("pic: create file: %s\n", homeDir+PIC_CONFIG)
	}
	return 0
}
