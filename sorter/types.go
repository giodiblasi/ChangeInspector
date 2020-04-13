package sorter

import "ChangeInspector/gitlog"

/*LogItem ...*/
type LogItem struct {
	FileName string
	Info     gitlog.FileInfo
}

/*LogSorter ...*/
type LogSorter struct {
	logs []LogItem
}

/*CreateSorter ...*/
func CreateSorter(gitLog gitlog.GitLog) LogSorter {
	arrayInfo := make([]LogItem, 0)
	for fileName, fileInfo := range gitLog.FilesInfo {
		arrayInfo = append(arrayInfo, LogItem{
			FileName: fileName,
			Info:     fileInfo,
		})
	}
	return LogSorter{
		logs: arrayInfo,
	}
}
