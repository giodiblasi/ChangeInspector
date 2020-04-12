package sortedcommits

import "ChangeInspector/commits"

/*OrderableFileInfo ...*/
type OrderableFileInfo struct {
	FileName string
	Info     commits.FileInfo
}

type orderableItems []OrderableFileInfo
