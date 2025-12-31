package daysteps

import (
	"fmt"
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
	// TODO: реализовать функцию
	splitStr := strings.Split(data, ",")

	if len(splitStr) != 2 {
		return 0, 0, fmt.Errorf("некорректный формат данных: %s", data)
	}

	numStr := splitStr[0]
	durationStr := splitStr[1]

	// Преобразование строки в int
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, 0, err
	}

	if num <= 0 {
		return 0, 0, fmt.Errorf("количество шагов %d должно быть больше 0", num)
	}

	// Преобразование строки в duration
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, 0, err
	}

	if duration <= 0 {
		return 0, 0, fmt.Errorf("длительность %d должна быть больше 0", duration)
	}

	return num, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	countSteps, duration, err := parsePackage(data)
	if err != nil {
		errStirng := fmt.Sprint(err)
		return errStirng
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

	result := fmt.Sprintf(`
Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.`, countSteps, distance, countCalories)

	return result
}
