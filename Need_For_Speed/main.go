package main

// TODO: define the 'Car' type struct
type Car struct {
    battery int
    batteryDrain int
    speed int
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car {
        battery: 100,
        speed: speed,
        batteryDrain: batteryDrain,
        distance: 0,
    }

}

// TODO: define the 'Track' type struct
type Track struct {
    distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track {
        distance,
    }
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {

	if car.battery >= car.batteryDrain {
        car.battery -= car.batteryDrain
        car.distance += car.speed
        return car
    } else {
        return car
    }
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	
    for car.battery >= car.batteryDrain {
        car = Drive(car)
    }
    if car.distance >= track.distance {
        return true
    }
    return false
    
}
