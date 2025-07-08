package main
import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	message := fmt.Sprintf("Welcome to my party, %s!", name)
    return message
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	message := fmt.Sprintf("Happy birthday %s! You are now %v years old!", name, age)
    return message
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	message := fmt.Sprintf("%s\nYou have been assigned to table %03d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", Welcome(name), table, direction, distance, neighbor)
    return message
}
