package main

var cmdGen = &Command{
	Run:       runGen,
	UsageLine: "gen ",
	Short:     "genelate ~/.ssh/config",
	Long: `

	`,
}

func init() {
	// Set your flag here like below.
	// cmdGen.Flag.BoolVar(&flagA, "a", false, "")
}

// runGen executes gen command and return exit code.
func runGen(args []string) int {

	return 0
}
