package business

import "strings"

func MonthNameToFrench(monthName string) string {
	r := strings.NewReplacer(
		"January", "Janvier",
		"February", "Février",
		"March", "Mars",
		"April", "Avril",
		"May", "Mai",
		"June", "Juin",
		"July", "Juillet",
		"August", "Août",
		"September", "Septembre",
		"October", "Octobre",
		"November", "Novembre",
		"December", "Décembre")

	return r.Replace(monthName)
}
