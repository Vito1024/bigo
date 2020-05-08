package utils

import "errors"

func RemoveElementInStringSlice(old []string, s string) []string {
	_new := make([]string, 0, len(old))
	for _, v := range old {
		if v != s {
			_new = append(_new, v)
		}
	}
	return _new
}

func IndexInString(str string, ch byte) int {
	for i := 0; i < len(str); i++ {
		if str[i] == ch {
			return i
		}
	}
	return -1 // Not found
}

func Split(str string, sep byte) ([]string, error) {
	res := make([]string, 0)

	for start := 0; start < len(str); {

		if str[start] == '"' {
			if start == len(str)-1 {
				// format error
				return nil, errors.New("format error")
			}
			end := IndexInString(str[start+1:], '"')
			if end < 0 {
				return nil, errors.New("format error")
				//format error
			} else {
				res = append(res, str[start+1:start+1+end])
				start = start + 1 + end + 1
				continue
			}
		}

		if str[start] == sep {
			start++
			continue
		}
		end := IndexInString(str[start:], sep)
		if end < 0 {
			// not found
			res = append(res, str[start:])
			return res, nil
		} else {
			res = append(res, str[start:start+end])
			start += end + 1
		}
	}
	return res, nil
}
