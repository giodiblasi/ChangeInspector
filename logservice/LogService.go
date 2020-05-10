package logservice

import (
	"ChangeInspector/gitlog"
	"ChangeInspector/sorter"
	"time"
)

/*LogService ...*/
type LogService struct {
	GitLog       *gitlog.GitLog
	SortableLogs *sorter.SortableLogs
}

/*NewLogService ...*/
func NewLogService(path string, before time.Time, after time.Time, filesToExclude []string) LogService {
	newgitLog := gitlog.NewGitLog(path, before, after, filesToExclude)
	newLogs := sorter.NewSorter(&newgitLog)
	return LogService{
		GitLog:       &newgitLog,
		SortableLogs: &newLogs,
	}
}

/*Update ...*/
func (logService *LogService) Update(before time.Time, after time.Time) {
	newLogService := NewLogService(logService.GitLog.Path, before, after, logService.GitLog.Ignore)

	logService.GitLog = newLogService.GitLog
	logService.SortableLogs = newLogService.SortableLogs
}
