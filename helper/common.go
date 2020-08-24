package helper

func InStrArr(str string, arr []string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}

func ComputeTotalPage(total int64, pageSize int64, ) int {
	totalPage := total / pageSize
	if total % pageSize != 0 {
		totalPage++
	}
	return int(totalPage)
}
