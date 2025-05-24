package main

import "fmt"

func getHour(roketinTotalHoursInDay int, input int) (roketinHour int, iteration int) {
	for input >= 10 {
		input /= roketinTotalHoursInDay
		iteration++
	}

	roketinHour = input
	return roketinHour, iteration
}

func getMinute(startMinute int, endMinute int, input int) (roketinMinute int) {
	fmt.Println("Input: ", input)
	fmt.Println("Start minute: ", startMinute)
	fmt.Println("End Minute: ", endMinute)
	for i := startMinute; i <= endMinute; i++ {
		fmt.Println(i)
		if i*100 >= input {
			roketinMinute = i - 1
			fmt.Println("minute: ", roketinMinute)
			return roketinMinute
		}
	}

	return roketinMinute
}

func main() {
	var hours, minutes, seconds int

	fmt.Print("Enter hours, minutes, and seconds (e.g. 13 45 22): ")

	// Validate input does not exceed 23:59:59
	if hours < 0 || hours > 23 || minutes < 0 || minutes > 59 || seconds < 0 || seconds > 59 {
		fmt.Println("Invalid input: time cannot exceed 23:59:59 and must not contain negative numbers.")
		return
	}

	_, err := fmt.Scanf("%d %d %d", &hours, &minutes, &seconds)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// convert time on earth to second
	totalElapsedSecondOnEarth := seconds + (minutes * 60) + (hours * 3600)
	totalElapsedSecondOnRoketin := totalElapsedSecondOnEarth * 100000 / 86400
	fmt.Println("totalElapsedSecondOnEarth: ", totalElapsedSecondOnEarth)

	// convert that total elapsed second in hrs:mnt:sec on Roketin Planet
	roketinTotalHoursInDay := 10
	roketinStartMinute := 0
	roketinEndMinute := 99

	resultedRoketinHour := 0
	resultedRoketinMinute := 0
	resultedRoketinSecond := 0

	roketinHour, iteration := getHour(roketinTotalHoursInDay, totalElapsedSecondOnRoketin)

	// jika puluh ribuan, maka set jam
	if iteration == 4 {
		fmt.Println("Set jam")
		resultedRoketinHour = roketinHour

		totalElapsedSecondOnRoketin -= resultedRoketinHour * 10000
		// misal sisa 3200

		roketinMinute := getMinute(roketinStartMinute, roketinEndMinute, totalElapsedSecondOnRoketin)

		resultedRoketinMinute = roketinMinute
		totalElapsedSecondOnRoketin -= resultedRoketinMinute * 100

		roketinSecond := totalElapsedSecondOnRoketin

		resultedRoketinSecond = roketinSecond

	} else if iteration == 3 {
		fmt.Println("Set minute")
		resultedRoketinMinute = getMinute(roketinStartMinute, roketinEndMinute, totalElapsedSecondOnRoketin)
		totalElapsedSecondOnRoketin -= resultedRoketinMinute * 100

		roketinSecond := totalElapsedSecondOnRoketin

		resultedRoketinSecond = roketinSecond

	} else {
		fmt.Println("Set second")
		fmt.Println("Elapsed: ", totalElapsedSecondOnRoketin)
		resultedRoketinSecond = totalElapsedSecondOnRoketin
	}

	fmt.Printf("On Earth %02d:%02d:%02d, on Roketin Planet %02d:%02d:%02d\n",
		hours, minutes, seconds,
		resultedRoketinHour, resultedRoketinMinute, resultedRoketinSecond)
}
