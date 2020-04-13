package sorter

import (
	"ChangeInspector/gitlog"
)

type criteria struct {
	Compare     func(items []LogItem) func(i, j int) bool
	SelectValue func(fileInfo gitlog.FileInfo) int64
}
type sortCriteria interface {
	getCriteria() criteria
}

/*ByChanges ...*/
type ByChanges []LogItem

func (ByChanges) getCriteria() criteria {
	return criteria{
		Compare: func(items []LogItem) func(i, j int) bool {
			return func(i, j int) bool {
				return items[i].Info.TotalChanges > items[j].Info.TotalChanges
			}
		},
		SelectValue: func(fileInfo gitlog.FileInfo) int64 {
			return fileInfo.TotalChanges
		},
	}
}

/*ByCommits ...*/
type ByCommits []LogItem

func (ByCommits) getCriteria() criteria {
	return criteria{
		Compare: func(items []LogItem) func(i, j int) bool {
			return func(i, j int) bool {
				return len(items[i].Info.Commits) > len(items[j].Info.Commits)
			}
		},
		SelectValue: func(fileInfo gitlog.FileInfo) int64 {
			return int64(len(fileInfo.Commits))
		},
	}
}
