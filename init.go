package main

import (
	"fmt"
	"os"
)

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

	if _, err := os.Stat(SSH_CONFIG); err != nil {
		fmt.Fprintf(os.Stderr, "pic: %s\n", err)
		os.Exit(1)
	}

	if _, err := os.Stat(SSH_CONFIG_DIR); err != nil {
		if err := os.MkdirAll(SSH_CONFIG_DIR, 0744); err != nil {
			fmt.Fprintf(os.Stderr, "pic: failed to create dir %s:%s\n", SSH_CONFIG_DIR, err)
			os.Exit(1)
		} else {
			fmt.Printf("pic: create directory: %s\n", SSH_CONFIG_DIR)
		}
	}

	// backup original config file to conf.d/base_config
	if _, err := os.Stat(BACKUP_SSH_CONFIG); err != nil {
		_ = os.Link(SSH_CONFIG, BACKUP_SSH_CONFIG)
		fmt.Printf("pic: copy config file: %s => %s\n", SSH_CONFIG, BACKUP_SSH_CONFIG)
	}

	if _, err := os.Stat(PIC_CONFIG); err != nil {
		f, err := os.Create(PIC_CONFIG)
		if err != nil {
			fmt.Printf("pic: failed to create file: %s:%s\n", PIC_CONFIG, err)
			os.Exit(1)
		}
		f.WriteString("include:\n#  - /path/to/config\n")
		f.Close()
		fmt.Printf("pic: create file: %s\n", PIC_CONFIG)
	}
	return 0
}
