package logservice

import (
	"ChangeInspector/gitlog"
	"ChangeInspector/sorter"
	"ChangeInspector/utils"
	"time"
)

/*LogService ...*/
type LogService struct {
	GitLog       *gitlog.GitLog
	SortableLogs *sorter.SortableLogs
	Filter       []string
}

/*NewLogService ...*/
func NewLogService(path string, before time.Time, after time.Time, filesToExclude []string) LogService {
	newgitLog := gitlog.NewGitLog(path, before, after)
	newLogs := sorter.NewSorter(&newgitLog, filesToExclude)
	return LogService{
		GitLog:       &newgitLog,
		SortableLogs: &newLogs,
		Filter:       filesToExclude,
	}
}

/*Update ...*/
func (logService *LogService) Update(before time.Time, after time.Time) {
	newLogService := NewLogService(logService.GitLog.Path, before, after, logService.Filter)

	logService.GitLog = newLogService.GitLog
	logService.SortableLogs = newLogService.SortableLogs
}

/*AddFileToFilter ...*/
func (logService *LogService) AddFileToFilter(fileName string) {
	logService.Filter = append(logService.Filter, fileName)
	sorter := sorter.NewSorter(logService.GitLog, logService.Filter)
	logService.SortableLogs = &sorter
}

/*RemoveFileFromFilter ...*/
func (logService *LogService) RemoveFileFromFilter(fileName string) bool {
	ok, updatedFilter := utils.RemoveFirst(logService.Filter, fileName)
	if ok {
		logService.Filter = updatedFilter
		sorter := sorter.NewSorter(logService.GitLog, logService.Filter)
		logService.SortableLogs = &sorter
	}
	return ok
}
