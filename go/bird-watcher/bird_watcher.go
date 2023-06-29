package birdwatcher

func SumSlice(sliceToSum []int) int {
    var sum int
    for _, n := range sliceToSum {
        sum += n
    }
    return sum
}

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	return SumSlice(birdsPerDay)
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	return SumSlice(birdsPerDay[7*week-7:7*week])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
    	birdsPerDay[i] += 1
    }

    return birdsPerDay
}
