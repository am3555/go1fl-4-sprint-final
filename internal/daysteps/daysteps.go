package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	splitStr := strings.Split(data, ",")

	if len(splitStr) != 2 {
		return 0, 0, fmt.Errorf("")
	}

	numStr := splitStr[0]
	durationStr := splitStr[1]

	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, 0, err
	}

	if num <= 0 {
		return 0, 0, fmt.Errorf("")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, 0, err
	}

	if duration <= 0 {
		return 0, 0, fmt.Errorf("")
	}

	return num, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	countSteps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return fmt.Sprint(err)
	}

	if countSteps <= 0 {
		return ""
	}

	distance := float64(countSteps) * stepLength / mInKm

	countCalories, err := spentcalories.WalkingSpentCalories(countSteps, weight, height, duration)
	if err != nil {
		errStirng := fmt.Sprint(err)
		return errStirng
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", countSteps, distance, countCalories)
	return result
}
