package main

import (
	"ChangeInspector/logservice"
	"ChangeInspector/web"
	"os"
	"time"
)

func main() {
	before := time.Now()
	after := before.AddDate(0, 0, -7)
	var gitFolder string = os.Args[1]
	logService := logservice.NewLogService(gitFolder, before, after)
	web.StartServer(&logService)

	// Console
	// for _, fileInfo := range sortedcommits.GetSorted(filesInfo, sortedcommits.ByCommits{}) {
	// 	fmt.Println("file:", fileInfo.FileName, "changed", fileInfo.Info.TotalChanges, "times in", len(fileInfo.Info.Commits), "commits")
	// }

}
