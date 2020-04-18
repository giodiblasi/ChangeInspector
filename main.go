package main

import (
	"ChangeInspector/gitlog"
	"ChangeInspector/web"
	"bufio"
	"os"
	"os/exec"
	"time"
)

func execGitLog(path string, consumer func(string)) {
	before := time.Now()
	after := before.AddDate(0, -1, 0)
	layout := "Jan-02-06"

	cmd := exec.Command("bash", "stat", path, after.Format(layout), before.Format(layout))
	stdout, _ := cmd.StdoutPipe()
	r := bufio.NewReader(stdout)
	cmd.Start()
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		consumer(string(line))
	}
	cmd.Wait()
}

func main() {

	var gitFolder string = os.Args[1]
	gitLog := gitlog.GitLog{}
	execGitLog(gitFolder, gitLog.ParseStream())
	web.StartServer(gitLog)

	// Console
	// for _, fileInfo := range sortedcommits.GetSorted(filesInfo, sortedcommits.ByCommits{}) {
	// 	fmt.Println("file:", fileInfo.FileName, "changed", fileInfo.Info.TotalChanges, "times in", len(fileInfo.Info.Commits), "commits")
	// }

}
