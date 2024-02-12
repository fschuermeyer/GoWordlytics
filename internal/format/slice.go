package format

func UniqueSlice(slice []string) []string {
	list := []string{}

	for _, item := range slice {
		if !SliceContains(list, item) {
			list = append(list, item)
		}
	}

	return list
}

func SliceContains(slice []string, item string) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}

	return false
}
