package main

func AppendUnique(uniqueSlice []string, newValue string) ([]string, bool) {
	for _, value := range uniqueSlice {
		if value == newValue {

			return uniqueSlice, false
		}
	}

	return append(uniqueSlice, newValue), true
}

func ValueInSlice(valueToCheck string, slice []string) bool {
	for _, val := range slice {
		if val == valueToCheck {

			return true
		}
	}

	return false
}

func UniqueSlice(inputSlice []string) []string {
	uniqueSlice := make([]string, 0, len(inputSlice))
	marker := make(map[string]bool)
	for _, val := range inputSlice {
		if _, ok := marker[val]; !ok {
			marker[val] = true
			uniqueSlice = append(uniqueSlice, val)
		}
	}

	return uniqueSlice
}
