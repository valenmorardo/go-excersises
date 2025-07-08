package main

// Package weather provides tools to process and work with functions relationed with weather.
// package weather

// CurrentCondition indicates an string value of current temperature condition.
var CurrentCondition string
// CurrentLocation indicates an string value of current Location of the temperature.
var CurrentLocation string


// Forecast returns an string value about the current location and current condition about the weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
