package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type codeQLAlert struct {
	State string `json:"state"`
	Rule struct {
		Severity              string `json:"severity"`
		SecuritySeverityLevel string `json:"security_severity_level"`
	} `json:"rule"`
}

var codeqlCmd = &cobra.Command{
	Use:   "codeql <owner/repo>",
	Short: "Check CodeQL security alerts",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repo := args[0]

		token, err := getToken()
		if err != nil {
			fmt.Println("❌", err)
			return
		}

		url := fmt.Sprintf(
			"%s/repos/%s/code-scanning/alerts?state=open&per_page=100",
			githubAPI,
			repo,
		)

		var alerts []codeQLAlert
		if err := githubGET(token, url, &alerts); err != nil {
			fmt.Println("❌", err)
			return
		}

		count := map[string]int{}

		for _, a := range alerts {
			if a.State != "open" {
				continue
			}

			sev := a.Rule.SecuritySeverityLevel
			if sev == "" {
				sev = "unknown"
			}
			count[sev]++
		}

		fmt.Println("🦾 Automato CodeQL Report")

		if len(count) == 0 {
			fmt.Println("✅ No open CodeQL security alerts")
			return
		}

		order := []string{"critical", "high", "medium", "low", "unknown"}
		for _, sev := range order {
			if n, ok := count[sev]; ok {
				fmt.Printf("%s: %d\n", sev, n)
			}
		}
	},
}
