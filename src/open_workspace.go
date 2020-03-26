package main

import (
	"fmt"
	"os/exec"
	"os"
	"strings"
	"path/filepath"
)


const (
	maxDepth = 1
)

var (
	workspaceDirs = os.Getenv("WORKSPACE_DIRS")
)

func fuzzyFindWorkspaces(args []string) (string, error) {
	splitWorkspaceDirs := strings.ReplaceAll(workspaceDirs, ",", " ")
	expandedFindCmd := fmt.Sprintf("find -L %s -maxdepth %d -type d", splitWorkspaceDirs, maxDepth)
	cmd := fmt.Sprintf("%s | /usr/local/bin/fzf -f '%s'", expandedFindCmd, strings.Join(args, " "))
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func convertAccountsToAutoCompleteItems(items string) string {
	lines := strings.Split(items, "\n")
	autoCompleteItems := make([]string, len(lines))
	for indx, line := range lines {
		if line != "" {
			autoCompleteItems[indx] = lineToAutoCompleteItem(line)
		}
	}
	return fmt.Sprintf("<items>%s</items>", strings.Join(autoCompleteItems, ""))
}

func lineToAutoCompleteItem(line string) string {
	fileName := filepath.Base(line)
	return fmt.Sprintf(
		"<item uid=\"%s\" arg=\"%s\"><title>%s</title><subtitle>%s</subtitle><autocomplete>%s</autocomplete></item>",
		line,
		line,
		fileName,
		line,
		line)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	accounts, err := fuzzyFindWorkspaces(os.Args[1:])
	check(err)
	
	items := convertAccountsToAutoCompleteItems(accounts)
	fmt.Println(items)
}