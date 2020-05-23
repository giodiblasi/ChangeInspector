package utils

/*Contains ...*/
func Contains(source []string, item string) bool {
	for _, v := range source {
		if v == item {
			return true
		}
	}
	return false
}

/*RemoveFirst ...*/
func RemoveFirst(source []string, item string) (bool, []string) {
	found, index := FindIndex(source, item)
	if !found {
		return false, nil
	}
	return true, append(source[:index], source[index+1:]...)

}

/*FindIndex ...*/
func FindIndex(source []string, item string) (bool, int) {
	for index, v := range source {
		if v == item {
			return true, index
		}
	}
	return false, -1
}
