package commits

/*CommitInfo ...*/
type CommitInfo struct {
	Hash    string
	Author  string
	Message string
}

/*FileInfo ...*/
type FileInfo struct {
	Commits        []CommitInfo
	TotalAdds      int64
	TotalRemotions int64
	TotalChanges   int64
}

func emptyFileInfo() *FileInfo {
	fileInfo := FileInfo{
		Commits:        make([]CommitInfo, 0),
		TotalAdds:      0,
		TotalRemotions: 0,
		TotalChanges:   0,
	}
	return &fileInfo
}

/*FileInfos ...*/
type FileInfos map[string]FileInfo
