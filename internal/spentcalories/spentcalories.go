package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	//"golang.org/x/text/width"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	splitStr := strings.Split(data, ",")

	if len(splitStr) != 3 {
		return 0, "", 0, fmt.Errorf("некорректный формат данных: %s", data)
	}

	stepsStr := splitStr[0]
	activeType := splitStr[1]
	durationStr := splitStr[2]

	// Преобразование строки в int
	countSteps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, "", 0, err
	}

	if countSteps <= 0 {
		return 0, "", 0, fmt.Errorf("количество шагов %d должно быть больше 0", countSteps)
	}

	// Преобразование строки в duration
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, "", 0, err
	}

	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("длительность %d должна быть больше 0", duration)
	}

	return countSteps, activeType, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	distanceKm := (height * stepLengthCoefficient * float64(steps)) / mInKm

	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	avgSpeed := distance(steps, height) / float64(duration.Hours())

	return avgSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	num, activeType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}

	switch activeType {
	case "Ходьба":
		distanceKm := distance(num, height)
		avgSpeed := meanSpeed(num, height, duration)
		countCaloriesWalk, err := WalkingSpentCalories(num, weight, height, duration)
		if err != nil {
			return "", err
		}
		resultWalk := fmt.Sprintf(`
Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f`, activeType, duration.Hours(), distanceKm, avgSpeed, countCaloriesWalk)
		return resultWalk, nil

	case "Бег":
		distanceKm := distance(num, height)
		avgSpeed := meanSpeed(num, height, duration)
		countCaloriesRun, err := RunningSpentCalories(num, weight, height, duration)
		if err != nil {
			return "", err
		}
		resultRun := fmt.Sprintf(`
Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f`, activeType, duration.Hours(), distanceKm, avgSpeed, countCaloriesRun)
		return resultRun, nil

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов %d должно быть больше 0", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен %.2f быть больше 0", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен %.2f быть больше 0", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("длительность %v должна быть больше 0", duration)
	}

	countCaloriesRun := (weight * meanSpeed(steps, height, duration) * float64(duration.Minutes())) / minInH

	return countCaloriesRun, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов %d должно быть больше 0", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес %.2f должен быть больше 0", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост %.2f должен быть больше 0", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("длительность %v должна быть больше 0", duration)
	}

	countCaloriesWalk := (weight * meanSpeed(steps, height, duration) * float64(duration.Minutes())) / minInH * walkingCaloriesCoefficient
	return countCaloriesWalk, nil
}
