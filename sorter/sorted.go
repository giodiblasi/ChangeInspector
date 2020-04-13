package sorter

import (
	"ChangeInspector/gitlog"
	"sort"
)

/*GoogleChartBarResult ...*/
type GoogleChartBarResult [][]interface{}

func toSortResult(items []LogItem, getValue func(fileInfo gitlog.FileInfo) int64) GoogleChartBarResult {
	result := make(GoogleChartBarResult, 0)
	for _, item := range items {
		result = append(result, []interface{}{item.FileName, getValue(item.Info)})
	}
	return result
}

/*Sort ...*/
func (logSorter LogSorter) Sort(sortCriteria sortCriteria) GoogleChartBarResult {
	criteria := sortCriteria.getCriteria()
	copyArray := make([]LogItem, len(logSorter.logs))
	copy(copyArray, logSorter.logs)
	sort.Slice(copyArray, criteria.Compare(copyArray))
	return toSortResult(copyArray, criteria.SelectValue)
}
