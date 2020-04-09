package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type commitInfo struct {
	Hash    string
	Author  string
	Message string
}

/*FileInfo ...*/
type FileInfo struct {
	Commits        []commitInfo
	TotalAdds      int64
	TotalRemotions int64
	TotalChanges   int64
}

func getCommitsResult(path string) string {
	cmd := exec.Command("bash", "stat", path)
	out, _ := cmd.CombinedOutput()
	return string(out)

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

/*FileInfos ...*/
type FileInfos map[string]FileInfo

func parse(input string) FileInfos {
	fileInfoMap := make(FileInfos)
	commits := strings.Split(input, "---------------------------")[1:]
	for _, commitStr := range commits {
		bodyStr := strings.Split(commitStr, "*******")

		commit := commitInfo{
			bodyStr[0],
			strings.ReplaceAll(bodyStr[1], "\n", ""),
			bodyStr[2]}

		changesStr := bodyStr[3]
		filesStats := strings.Split(changesStr, "\n")
		for _, fileStat := range filterEmpty(filesStats) {
			fileChanges := strings.Split(fileStat, "\t")
			fileName := fileChanges[2]
			fileAdds, _ := strconv.ParseInt(fileChanges[0], 10, 64)
			fileRemotions, _ := strconv.ParseInt(fileChanges[1], 10, 64)

			fileInfo, ok := fileInfoMap[fileName]
			if !ok {
				fileInfo = FileInfo{
					Commits:        make([]commitInfo, 0),
					TotalAdds:      0,
					TotalRemotions: 0,
					TotalChanges:   0,
				}
			}

			fileInfoMap[fileName] = FileInfo{
				Commits:        append(fileInfo.Commits, commit),
				TotalAdds:      fileInfo.TotalAdds + fileAdds,
				TotalRemotions: fileInfo.TotalRemotions + fileRemotions,
				TotalChanges:   fileInfo.TotalChanges + fileAdds + fileRemotions,
			}
		}
	}
	return fileInfoMap
}

func main() {
	var result string = getCommitsResult(os.Args[1])
	var fileInfos FileInfos = parse(result)

	for fileName, fileInfo := range fileInfos {
		fmt.Println(fileName)
		fmt.Println(fileInfo.TotalChanges)
	}

}
