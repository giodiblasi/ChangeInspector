package main

import (
	"ChangeInspector/logservice"
	"ChangeInspector/web"
	"os"
	"strings"
	"time"
)

func main() {
	before := time.Now()
	after := before.AddDate(0, 0, -7)
	var gitFolder string = os.Args[1]
	filesToExclude := ""
	if len(os.Args) > 2 {
		filesToExclude = os.Args[2]
	}
	logService := logservice.NewLogService(gitFolder, before, after, strings.Split(filesToExclude, ";"))
	web.StartServer(&logService)

	// Console
	// for _, fileInfo := range sortedcommits.GetSorted(filesInfo, sortedcommits.ByCommits{}) {
	// 	fmt.Println("file:", fileInfo.FileName, "changed", fileInfo.Info.TotalChanges, "times in", len(fileInfo.Info.Commits), "commits")
	// }

}
