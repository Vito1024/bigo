package utils

func RemoveElementInStringSlice(old []string, s string) []string {
	_new := make([]string, 0, len(old))
	for _, v := range old {
		if v != s {
			_new = append(_new, v)
		}
	}
	return _new
}
