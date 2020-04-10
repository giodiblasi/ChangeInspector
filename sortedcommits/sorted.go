package sortedcommits

import (
	"ChangeInspector/commits"
	"sort"
)

/*OrderableFileInfo ...*/
type OrderableFileInfo struct {
	FileName string
	Info     commits.FileInfo
}

/*GetSorted ...*/
func GetSorted(fileInfos commits.FileInfos) []OrderableFileInfo {
	fileInfoArray := make([]OrderableFileInfo, 0)
	for name, info := range fileInfos {
		fileInfoArray = append(fileInfoArray, OrderableFileInfo{
			FileName: name,
			Info:     info,
		})
	}

	sort.Slice(fileInfoArray, func(i, j int) bool {
		return fileInfoArray[i].Info.TotalChanges > fileInfoArray[j].Info.TotalChanges
	})

	return fileInfoArray
}
