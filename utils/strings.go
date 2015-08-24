package utils

// StringInSlice - Will check if string in list.
// This is equivalent to python if x in []
func StringInSlice(str string, list []string) bool {
	for _, value := range list {
		if value == str {
			return true
		}
	}
	return false
}

// KeyInSlice - Will check if key in list.
// This is equivalent to python if x in []
func KeyInSlice(str string, list map[string]interface{}) bool {
	for key, _ := range list {
		if key == str {
			return true
		}
	}
	return false
}
