package main

import (
	"ChangeInspector/gitlog"
	"ChangeInspector/web"
	"os"
	"time"
)

func main() {
	before := time.Now()
	after := before.AddDate(0, 0, -7)
	layout := "Jan-02-06"

	var gitFolder string = os.Args[1]
	gitLog := gitlog.GitLog{Path: gitFolder}
	gitLog.Update(before.Format(layout), after.Format(layout))
	web.StartServer(gitLog)

	// Console
	// for _, fileInfo := range sortedcommits.GetSorted(filesInfo, sortedcommits.ByCommits{}) {
	// 	fmt.Println("file:", fileInfo.FileName, "changed", fileInfo.Info.TotalChanges, "times in", len(fileInfo.Info.Commits), "commits")
	// }

}
