//Package weather forecasts weather.
package weather

//CurrentCondition describes current condition.
var CurrentCondition string

//CurrentLocation describes current location.
var CurrentLocation string

//Forecast forecasts the weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
