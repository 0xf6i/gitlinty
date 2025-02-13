package local

import "strings"

func CheckForLicense(fileSlice []string) (bool, int, []string) {
	var licenseSlice []string
	licenseFileCount := 0

	for _, file := range fileSlice {
		if strings.Contains(file, "LICENSE") {
			licenseFileCount++
			licenseSlice = append(licenseSlice, file)
		}
	}
	return licenseFileCount > 0, licenseFileCount, licenseSlice
}
