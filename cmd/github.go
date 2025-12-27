package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

const githubAPI = "https://api.github.com"

func getToken() (string, error) {
	home, _ := os.UserHomeDir()
	data, err := os.ReadFile(filepath.Join(home, ".automato", "token"))
	if err != nil {
		return "", fmt.Errorf("run `automato init <token>` first")
	}
	return string(data), nil
}

func githubGET(token, url string, out any) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set(
		"Accept",
		"application/vnd.github+json, application/vnd.github.code-scanning+json",
	)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("GitHub API error: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(out)
}
