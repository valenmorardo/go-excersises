package main

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (successRate*float64(productionRate))/100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)/60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	cantAutos95k := (carsCount / 10)
    cantAutos10k := carsCount%10

    return uint((cantAutos95k*95000) + (cantAutos10k*10000))
   // return ((carsCount / 10) * 95000) + ((carsCount%10)*10000)
}
