package main

func ValueInSlice(slice []string, valueToCheck string) bool {
	for _, val := range slice {
		if val == valueToCheck {

			return true
		}
	}

	return false
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
