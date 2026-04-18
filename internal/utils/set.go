package utils

func SliceToSet(s []string) map[string]struct{} {
	set := make(map[string]struct{}, len(s))
	for _, v := range s {
		set[v] = struct{}{}
	}
	return set
}

func GetMissingKeys(a, b map[string]struct{}) []string {
	missing := []string{}
	for k := range a {
		if _, ok := b[k]; !ok {
			missing = append(missing, k)
		}
	}

	return missing
}
