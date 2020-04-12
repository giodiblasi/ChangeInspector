package sortedcommits

import (
	"ChangeInspector/commits"
)

type criteria struct {
	Compare     func(items orderableItems) func(i, j int) bool
	SelectValue func(fileInfo commits.FileInfo) int64
}
type sortCriteria interface {
	getCriteria() criteria
}

/*ByChanges ...*/
type ByChanges []OrderableFileInfo

func (ByChanges) getCriteria() criteria {
	return criteria{
		Compare: func(items orderableItems) func(i, j int) bool {
			return func(i, j int) bool {
				return items[i].Info.TotalChanges > items[j].Info.TotalChanges
			}
		},
		SelectValue: func(fileInfo commits.FileInfo) int64 {
			return fileInfo.TotalChanges
		},
	}
}

/*ByCommits ...*/
type ByCommits []OrderableFileInfo

func (ByCommits) getCriteria() criteria {
	return criteria{
		Compare: func(items orderableItems) func(i, j int) bool {
			return func(i, j int) bool {
				return len(items[i].Info.Commits) > len(items[j].Info.Commits)
			}
		},
		SelectValue: func(fileInfo commits.FileInfo) int64 {
			return int64(len(fileInfo.Commits))
		},
	}
}
