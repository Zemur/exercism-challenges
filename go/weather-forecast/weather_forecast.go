// Package weather forecasts the weather for a location.
package weather


// CurrentCondition: The current condition as a string.
var CurrentCondition string
// CurrentLocation: The location as a string.
var CurrentLocation string

// Forecast takes the city and condtion provided and returns a formatted string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
