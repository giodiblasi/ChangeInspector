package sortedcommits

type sortCriteria interface {
	getCriteria(items orderableItems) func(i, j int) bool
}

/*ByChanges ...*/
type ByChanges []OrderableFileInfo

func (ByChanges) getCriteria(items orderableItems) func(i, j int) bool {
	return func(i, j int) bool {
		return items[i].Info.TotalChanges > items[j].Info.TotalChanges
	}
}

/*ByCommits ...*/
type ByCommits []OrderableFileInfo

func (ByCommits) getCriteria(items orderableItems) func(i, j int) bool {
	return func(i, j int) bool {
		return len(items[i].Info.Commits) > len(items[j].Info.Commits)
	}
}
