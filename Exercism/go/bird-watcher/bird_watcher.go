package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var sum int
	for _, v := range birdsPerDay {
		sum += v
	}

	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var weeks int = len(birdsPerDay) / 7
	var result int
	if week <= weeks {
		for i := (week * 7) - 7; i < week*7; i++ {
			result += birdsPerDay[i]
		}

	} else {
		result = -1
	}

	return result
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}
