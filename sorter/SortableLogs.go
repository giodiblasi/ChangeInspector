package sorter

import (
	"ChangeInspector/gitlog"
	"ChangeInspector/utils"
)

/*LogItem ...*/
type LogItem struct {
	FileName string
	Info     gitlog.FileInfo
}

/*SortableLogs ...*/
type SortableLogs struct {
	logs []LogItem
}

/*NewSorter ...*/
func NewSorter(gitLog *gitlog.GitLog, filter []string) SortableLogs {
	arrayInfo := make([]LogItem, 0)
	for fileName, fileInfo := range gitLog.FilesInfo {
		if !utils.Contains(filter, fileName) {
			arrayInfo = append(arrayInfo, LogItem{
				FileName: fileName,
				Info:     fileInfo,
			})
		}
	}
	return SortableLogs{
		logs: arrayInfo,
	}
}
