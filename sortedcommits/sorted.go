package sortedcommits

import (
	"ChangeInspector/commits"
	"sort"
)

/*GetSorted ...*/
func GetSorted(fileInfos commits.FileInfos, sortCriteria sortCriteria) []OrderableFileInfo {
	fileInfoArray := make([]OrderableFileInfo, 0)
	for name, info := range fileInfos {
		fileInfoArray = append(fileInfoArray, OrderableFileInfo{
			FileName: name,
			Info:     info,
		})
	}

	sort.Slice(fileInfoArray, sortCriteria.getCriteria(fileInfoArray))

	return fileInfoArray
}
