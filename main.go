package main

import (
	"ChangeInspector/commits"
	"ChangeInspector/sortedcommits"
	"fmt"
	"os"
	"os/exec"
)

func getCommitsResult(path string) string {
	cmd := exec.Command("bash", "stat", path)
	out, _ := cmd.CombinedOutput()
	return string(out)
}

func main() {
	var result string = getCommitsResult(os.Args[1])
	var fileInfos commits.FileInfos = commits.Parse(result)

	for _, fileInfo := range sortedcommits.GetSorted(fileInfos) {
		fmt.Println("file:", fileInfo.FileName, "changed", fileInfo.Info.TotalChanges, "times in", len(fileInfo.Info.Commits), "commits")
	}

}
