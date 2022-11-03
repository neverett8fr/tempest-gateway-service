package auth

func contains(arr []string, search string) bool {

	for _, val := range arr {
		if val == search {
			return true
		}
	}

	return false
}
