package sortedcommits

import (
	"ChangeInspector/commits"
	"sort"
)

/*GoogleChartBarResult ...*/
type GoogleChartBarResult [][]interface{}

func toSortResult(items []OrderableFileInfo, getValue func(fileInfo commits.FileInfo) int64) GoogleChartBarResult {
	result := make(GoogleChartBarResult, 0)
	for _, item := range items {
		result = append(result, []interface{}{item.FileName, getValue(item.Info)})
	}
	return result
}

/*Sort ...*/
func Sort(src []OrderableFileInfo, sortCriteria sortCriteria) GoogleChartBarResult {
	criteria := sortCriteria.getCriteria()
	copyArray := make([]OrderableFileInfo, len(src))
	copy(copyArray, src)
	sort.Slice(copyArray, criteria.Compare(copyArray))
	return toSortResult(copyArray, criteria.SelectValue)
}
