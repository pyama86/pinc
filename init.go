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
		fmt.Fprintf(os.Stderr, "pinc: %s\n", err)
		return 1
	}

	if _, err := os.Stat(SSH_CONFIG_DIR); err != nil {
		if err := os.MkdirAll(SSH_CONFIG_DIR, 0744); err != nil {
			fmt.Fprintf(os.Stderr, "pinc: failed to create dir %s:%s\n", SSH_CONFIG_DIR, err)
			return 1
		} else {
			fmt.Printf("pinc: create directory: %s\n", SSH_CONFIG_DIR)
		}
	}

	// backup original config file to conf.d/base_config
	if _, err := os.Stat(BACKUP_SSH_CONFIG); err != nil {
		if err := os.Link(SSH_CONFIG, BACKUP_SSH_CONFIG); err != nil {
			fmt.Fprintf(os.Stderr, "pinc: failed to copy file %s:%s\n", SSH_CONFIG, err)
			return 1
		} else {
			fmt.Printf("pinc: copy config file: %s => %s\n", SSH_CONFIG, BACKUP_SSH_CONFIG)
		}
	}

	if _, err := os.Stat(PINC_CONFIG); err != nil {
		f, err := os.Create(PINC_CONFIG)
		if err != nil {
			fmt.Printf("pinc: failed to create file: %s:%s\n", PINC_CONFIG, err)
			return 1
		}
		f.WriteString("include:\n#  - /path/to/config\n")
		f.Close()
		fmt.Printf("pinc: create file: %s\n", PINC_CONFIG)
	}
	return 0
}
