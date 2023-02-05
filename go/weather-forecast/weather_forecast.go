// Package weather is for showing the current weather conditions at a specific location.
package weather

// CurrentCondition described the current weather condition.
var CurrentCondition string

// CurrentLocation describes the geolocation who's forecast is displayed.
var CurrentLocation string

// Forecast provides a listing of the current weather condition and location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
