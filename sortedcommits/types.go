package sortedcommits

import "ChangeInspector/commits"

/*OrderableFileInfo ...*/
type OrderableFileInfo struct {
	FileName string
	Info     commits.FileInfo
}

/*SortResult ...*/
type SortResult struct {
	FileName string
	Value    int64
}

type orderableItems []OrderableFileInfo
