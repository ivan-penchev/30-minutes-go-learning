package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

func main() {
	err := cliCmd.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err.Error())
	}
}

// ##################################################################
// CLI root
// ##################################################################
const CLI_NAME = "myCli"

// package global variables
var (
	version = "0.0.1" // to be populated by linker

	// Cobra parameters that are common across aggregates
	url      string
	showVers bool // whether to print version info or not
	verbose  bool // whether to display request info

	// myCli root command
	cliCmd = &cobra.Command{
		Use:     CLI_NAME,
		Short:   CLI_NAME,
		Long:    ``,
		Example: "",
		RunE:    preFlight,
	}
)

func preFlight(ccmd *cobra.Command, args []string) error {
	// if --version is passed print the version info
	if showVers {
		fmt.Printf(CLI_NAME+" %s\n", version)
	} else {
		fmt.Printf("usage: " + CLI_NAME + " get\n")
	}
	return nil
}

func init() {
	cliCmd.Flags().BoolVarP(&showVers, "version", "v", false, "Display the current version of this CLI")

	getCmd.Flags().StringVarP(&url, "url", "u", "", "URL of the server to access")
	getCmd.MarkFlagRequired("url")
	cliCmd.AddCommand(getCmd)
}

// ##################################################################
// The actual command
// ##################################################################

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get command",
	Long:  `get command`,
	Run:   getVerb,
}

// same code as http client example
func getVerb(ccmd *cobra.Command, args []string) {
	var body []byte

	// we introduce a new package here
	resp, err := http.Get(url)
	if err == nil {
		defer resp.Body.Close()
	}

	if err == nil {
		body, err = ioutil.ReadAll(resp.Body)
	}

	if err == nil {
		fmt.Println(string(body))
	} else {
		fmt.Println(err.Error())
	}

}
