package sortedcommits

import (
	"ChangeInspector/commits"
	"sort"
)

func toSortResult(items []OrderableFileInfo, getValue func(fileInfo commits.FileInfo) int64) []SortResult {
	result := make([]SortResult, 0)
	for _, item := range items {
		result = append(result, SortResult{
			FileName: item.FileName,
			Value:    getValue(item.Info),
		})
	}
	return result
}

/*Sort ...*/
func Sort(src []OrderableFileInfo, sortCriteria sortCriteria) []SortResult {
	criteria := sortCriteria.getCriteria()
	copyArray := make([]OrderableFileInfo, len(src))
	copy(copyArray, src)
	sort.Slice(copyArray, criteria.Compare(copyArray))
	return toSortResult(copyArray, criteria.SelectValue)
}
