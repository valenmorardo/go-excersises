
package main

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
    func TotalBirdCount(birdsPerDay []int) int {
	var totalBirds int
    for i:=0;i<len(birdsPerDay);i++ {
        totalBirds = totalBirds + birdsPerDay[i]
    }
    return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	startWeekIndex := (week - 1)*7
	finalWeekIndex := startWeekIndex+7
	birdsInWeekParameter := birdsPerDay[startWeekIndex:finalWeekIndex]
	totalBirds := TotalBirdCount(birdsInWeekParameter)
    /* for i:=0;i<len(birdsInWeekParameter);i++ {
        totalBirds = totalBirds + birdsInWeekParameter[i]
    } */
	return totalBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i:=0;i<len(birdsPerDay);i++ {
        if i%2 == 0 {
            birdsPerDay[i] += 1
        }
    }

    return birdsPerDay
}
