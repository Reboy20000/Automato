package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type workflowRuns struct {
	Runs []struct {
		Conclusion string `json:"conclusion"`
		HeadSHA    string `json:"head_sha"`
		Name       string `json:"name"`
	} `json:"workflow_runs"`
}

var codeciCmd = &cobra.Command{
	Use:   "codeci <owner/repo>",
	Short: "Check GitHub Actions CI status",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repo := args[0]

		token, err := getToken()
		if err != nil {
			fmt.Println("❌", err)
			return
		}

		url := fmt.Sprintf(
			"%s/repos/%s/actions/runs?per_page=1",
			githubAPI,
			repo,
		)

		var runs workflowRuns
		if err := githubGET(token, url, &runs); err != nil {
			fmt.Println("❌", err)
			return
		}

		if len(runs.Runs) == 0 {
			fmt.Println("⚠️ No CI runs found")
			return
		}

		run := runs.Runs[0]

		fmt.Println("🦾 Automato CI Report")
		fmt.Println("Workflow:", run.Name)
		fmt.Println("Commit:", run.HeadSHA[:7])
		fmt.Println("Status:", run.Conclusion)

		if run.Conclusion == "success" {
			fmt.Println("✅ CI passed")
		} else {
			fmt.Println("❌ CI failed")
		}
	},
}
