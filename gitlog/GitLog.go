package gitlog

/*CommitInfo ...*/
type CommitInfo struct {
	Hash    string
	Author  string
	Message string
}

/*FileInfo ...*/
type FileInfo struct {
	Commits        []*CommitInfo
	TotalAdds      int64
	TotalRemotions int64
	TotalChanges   int64
}

func emptyFileInfo() *FileInfo {
	fileInfo := FileInfo{
		Commits:        make([]*CommitInfo, 0),
		TotalAdds:      0,
		TotalRemotions: 0,
		TotalChanges:   0,
	}
	return &fileInfo
}

/*FilesInfo ...*/
type FilesInfo map[string]FileInfo

/*CommitsInfo ...*/
type CommitsInfo map[string]CommitInfo

/*GitLog ...*/
type GitLog struct {
	FilesInfo FilesInfo
	Commits   CommitsInfo
}

/*GetFileInfo ...*/
func (gitLog GitLog) GetFileInfo(fileName string) FileInfo {
	info, _ := gitLog.FilesInfo[fileName]
	return info
}

/*GetFileCommits ...*/
func (gitLog GitLog) GetFileCommits(fileName string) []CommitInfo {
	info := gitLog.FilesInfo[fileName]
	commits := make([]CommitInfo, 0)
	for _, commit := range info.Commits {
		commits = append(commits, *commit)
	}
	return commits
}
