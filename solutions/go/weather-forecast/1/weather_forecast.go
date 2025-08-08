// Package weather contains functionality for forecasting the current weather conditions in Goblinocus.
package weather

// CurrentCondition holds the current condition.
var CurrentCondition string

// CurrentLocation holds the current location.
var CurrentLocation string

// Forecast gives a forecast for a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
