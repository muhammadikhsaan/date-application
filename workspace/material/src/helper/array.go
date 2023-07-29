package helper

func SearchInArray[T comparable](target []T, keyword T) bool {
	for _, v := range target {
		if v == keyword {
			return true
		}
	}

	return false
}
