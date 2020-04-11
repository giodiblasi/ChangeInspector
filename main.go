package main

import (
	"ChangeInspector/commits"
	"ChangeInspector/sortedcommits"
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
	var result string = getCommitsResult(gitFolder)
	var filesInfo commits.FileInfos = commits.Parse(result)
	data := sortedcommits.GetSorted(filesInfo, sortedcommits.ByChanges{})
	web.StartServer(data)

	// Console
	// for _, fileInfo := range sortedcommits.GetSorted(filesInfo, sortedcommits.ByCommits{}) {
	// 	fmt.Println("file:", fileInfo.FileName, "changed", fileInfo.Info.TotalChanges, "times in", len(fileInfo.Info.Commits), "commits")
	// }

}
