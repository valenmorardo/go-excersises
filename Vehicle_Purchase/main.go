package main

import "fmt"
// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if (kind == "car" || kind == "truck") {
        return true
    }
    return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if (option1 < option2) {
        return fmt.Sprintf("%s is clearly the better choice.", option1)
    } else {
        return fmt.Sprintf("%s is clearly the better choice.", option2)
    }
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	/* if (age < 3) {
        return (80*originalPrice)/100
    } else if (age <= 10) {
        return (50*originalPrice)/100
    } else {
        return (70*originalPrice)/100
    } */

    if (age >= 3 && age < 10) {
        return (70*originalPrice)/100
    } else if (age < 3) {
        return (80*originalPrice)/100
    } else {
        return (50*originalPrice)/100
    }
}
