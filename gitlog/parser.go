package gitlog

import (
	"bufio"
	"os/exec"
	"strconv"
	"strings"
)

func execGitLog(path string, before, after string, consumer func(string)) {
	cmd := exec.Command("bash", "stat", path, after, before)
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

func filterEmpty(source []string) []string {
	output := make([]string, 0)
	for _, value := range source {
		if strings.ReplaceAll(value, " ", "") != "" {
			output = append(output, value)
		}
	}
	return output
}

func (gitLog *GitLog) parseCommitString(commitStr string) {
	bodyStr := strings.Split(commitStr, "*******")

	commit := CommitInfo{
		strings.ReplaceAll(bodyStr[0], "\n", ""),
		strings.ReplaceAll(bodyStr[1], "\n", ""),
		strings.Trim(bodyStr[2], "\n")}
	_, hasCommit := gitLog.Commits[commit.Hash]
	if !hasCommit {
		gitLog.Commits[commit.Hash] = commit
	}
	changesStr := bodyStr[3]
	filesStats := strings.Split(changesStr, "\n")
	for _, fileStat := range filterEmpty(filesStats) {
		fileChanges := strings.Split(fileStat, "\t")
		fileName := fileChanges[2]
		fileAdds, _ := strconv.ParseInt(fileChanges[0], 10, 64)
		fileRemotions, _ := strconv.ParseInt(fileChanges[1], 10, 64)

		fileInfo, ok := gitLog.FilesInfo[fileName]
		if !ok {
			fileInfo = *emptyFileInfo()
		}

		gitLog.FilesInfo[fileName] = FileInfo{
			Commits:        append(fileInfo.Commits, &commit),
			TotalAdds:      fileInfo.TotalAdds + fileAdds,
			TotalRemotions: fileInfo.TotalRemotions + fileRemotions,
			TotalChanges:   fileInfo.TotalChanges + fileAdds + fileRemotions,
		}
	}
}

/*Update ...*/
func (gitLog *GitLog) Update(before string, after string) {
	commitStr := ""
	gitLog.Commits = make(CommitsInfo)
	gitLog.FilesInfo = make(FilesInfo)
	execGitLog(gitLog.Path, before, after, func(line string) {
		if line == "---------------------------" {
			if commitStr != "" {
				gitLog.parseCommitString(commitStr)
				commitStr = ""
			}
		} else {
			commitStr += "\n" + line
		}
	})
}
