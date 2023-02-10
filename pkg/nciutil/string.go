package nciutil

func FirstNonEmpty(strings []string) string {
	for _, str := range strings {
		if str != "" {
			return str
		}
	}
	return ""
}

func GetValueFromMap(m map[string]string, key string) string {
	value, ok := m[key]
	if !ok {
		return ""
	}
	return value
}
