package main

func AppendUniqueSlice(uniqueSlice []string, newValue string) ([]string, bool) {
	for _, value := range uniqueSlice {
		if value == newValue {

			return uniqueSlice, false
		}
	}

	return append(uniqueSlice, newValue), true
}

func ValueInSlice(slice []string, valueToCheck string) bool {
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

func IntersectionOfTwoSlices(sliceA []string, sliceB []string) []string {
	intersectionMap := make(map[string]bool)
	var intersectionSlice []string

	for _, value := range sliceA {
		intersectionMap[value] = true
	}
	for _, value := range sliceB {
		if _, ok := intersectionMap[value]; ok {
			intersectionSlice = append(intersectionSlice, value)
		}
	}
	return intersectionSlice
}

func sliceToInterface(slice []string) interface{} {
	newInterface := make([]interface{}, len(slice), len(slice))
	for value := range slice {
		newInterface[value] = slice[value]
	}

	return newInterface
}
