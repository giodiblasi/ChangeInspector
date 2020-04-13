package gitlog

import (
	"strconv"
	"strings"
)

func filterEmpty(source []string) []string {
	output := make([]string, 0)
	for _, value := range source {
		if strings.ReplaceAll(value, " ", "") != "" {
			output = append(output, value)
		}
	}
	return output
}

/*Parse ...*/
func Parse(input string) GitLog {
	fileInfoMap := make(FilesInfo)
	commitsInfoMap := make(CommitsInfo)
	commits := strings.Split(input, "---------------------------")[1:]
	for _, commitStr := range commits {
		bodyStr := strings.Split(commitStr, "*******")

		commit := CommitInfo{
			strings.ReplaceAll(bodyStr[0], "\n", ""),
			strings.ReplaceAll(bodyStr[1], "\n", ""),
			bodyStr[2]}
		_, hasCommit := commitsInfoMap[commit.Hash]
		if !hasCommit {
			commitsInfoMap[commit.Hash] = commit
		}
		changesStr := bodyStr[3]
		filesStats := strings.Split(changesStr, "\n")
		for _, fileStat := range filterEmpty(filesStats) {
			fileChanges := strings.Split(fileStat, "\t")
			fileName := fileChanges[2]
			fileAdds, _ := strconv.ParseInt(fileChanges[0], 10, 64)
			fileRemotions, _ := strconv.ParseInt(fileChanges[1], 10, 64)

			fileInfo, ok := fileInfoMap[fileName]
			if !ok {
				fileInfo = *emptyFileInfo()
			}

			fileInfoMap[fileName] = FileInfo{
				Commits:        append(fileInfo.Commits, commit.Hash),
				TotalAdds:      fileInfo.TotalAdds + fileAdds,
				TotalRemotions: fileInfo.TotalRemotions + fileRemotions,
				TotalChanges:   fileInfo.TotalChanges + fileAdds + fileRemotions,
			}
		}
	}
	return GitLog{
		Commits:   commitsInfoMap,
		FilesInfo: fileInfoMap,
	}
}
