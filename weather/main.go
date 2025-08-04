package weather

// CurrentCondition describes the weather condition.
var CurrentCondition string

// CurrentLocation describes the location of the forecast.
var CurrentLocation string

// Forecast returns the current weather forecast for a given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
