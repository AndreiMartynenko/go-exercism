package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	bird_counter := 0
	for _, count := range birdsPerDay {
		bird_counter += count
	}
	return bird_counter
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	if week < 1 || week > 52 {
		panic("Week number must be between 1 and 52")
	}
	startIndex := (week - 1) * 7
	if startIndex >= len(birdsPerDay) {
		panic("Week exceeds the number of recorded days")
	}
	endIndex := startIndex + 7
	if endIndex > len(birdsPerDay) {
		endIndex = len(birdsPerDay)
	}
	bird_counter := 0
	for i := startIndex; i < endIndex; i++ {
		bird_counter += birdsPerDay[i]
	}
	return bird_counter
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}
	return birdsPerDay
}
