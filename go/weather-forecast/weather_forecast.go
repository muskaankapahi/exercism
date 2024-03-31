// Package weather provides weather stools. 
package weather

// CurrentCondition represents current weather condition.
var CurrentCondition string
// CurrentLocation represents current location condition.
var CurrentLocation string

// Forecast provides tool to append current location to current condition. 
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
