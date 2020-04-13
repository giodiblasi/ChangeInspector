package main

import (
	"ChangeInspector/gitlog"
	"ChangeInspector/web"
	"os"
	"os/exec"
)

func getCommitsResult(path string) string {
	cmd := exec.Command("bash", "stat", path)
	out, _ := cmd.CombinedOutput()
	return string(out)
}

func main() {
	var gitFolder string = os.Args[1]
	result := getCommitsResult(gitFolder)
	gitLog := gitlog.Parse(result)
	web.StartServer(gitLog)

	// Console
	// for _, fileInfo := range sortedcommits.GetSorted(filesInfo, sortedcommits.ByCommits{}) {
	// 	fmt.Println("file:", fileInfo.FileName, "changed", fileInfo.Info.TotalChanges, "times in", len(fileInfo.Info.Commits), "commits")
	// }

}
