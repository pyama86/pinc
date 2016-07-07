package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"gopkg.in/yaml.v2"
)

var cmdGen = &Command{
	Run:       runGen,
	UsageLine: "gen ",
	Short:     "generate ~/.ssh/config from ~/.ssh/conf.d and ~/.ssh/pinc.yml",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdGen.Flag.BoolVar(&flagA, "a", false, "")
}

// runGen executes gen command and return exit code.
func runGen(args []string) int {
	var mergeContents string
	mergeContents = readFiles(SSH_CONFIG_DIR + "/")

	buf, err := ioutil.ReadFile(PIC_CONFIG)
	if err != nil {
		fmt.Fprintf(os.Stderr, "pinc: %s\n", err)
	}

	m := make(map[string][]string)
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		fmt.Fprintf(os.Stderr, "pinc: %s\n", err)
		return 1
	}

	for _, p := range m["include"] {
		mergeContents += readFiles(string(p))
	}

	file, err := os.Create(SSH_CONFIG)
	if err != nil {
		fmt.Fprintf(os.Stderr, "pinc: %s\n", err)
	}
	defer file.Close()
	file.Write(([]byte)(mergeContents))
	fmt.Println("update config:", SSH_CONFIG)
	return 0
}

func readFiles(root string) string {
	var mergeContents string
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			var c string
			var flg bool
			r := regexp.MustCompile(`^Host`)

			rel, err := filepath.Rel(root, path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "pinc: %s\n", err)
				return nil
			}

			fp, err := os.Open(root + "/" + rel)
			if err != nil {
				fmt.Fprintf(os.Stderr, "pinc: %s\n", err)
				return nil
			}
			defer fp.Close()

			reader := bufio.NewReaderSize(fp, 4096)
			for {
				line, _, err := reader.ReadLine()
				sl := string(line) + "\n"
				c += sl
				if r.MatchString(sl) {
					flg = true
				}

				if err == io.EOF {
					break
				} else if err != nil {
					fmt.Fprintf(os.Stderr, "pinc: %s\n", err)
					return nil
				}
			}

			if flg {
				mergeContents += string(c)
			} else {
				return nil
			}
			fmt.Println("load: " + root + rel)
			return nil
		})

	if err != nil {
		fmt.Println(err)
	}
	return mergeContents
}
