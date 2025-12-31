package spentcalories

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	splitStr := strings.Split(data, ",")

	if len(splitStr) != 3 {
		return 0, "", 0, fmt.Errorf("")
	}

	stepsStr := splitStr[0]
	activeType := splitStr[1]
	durationStr := splitStr[2]

	countSteps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, "", 0, err
	}

	if countSteps <= 0 {
		return 0, "", 0, fmt.Errorf("")
	}

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, "", 0, err
	}

	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("")
	}

	return countSteps, activeType, duration, nil
}

func distance(steps int, height float64) float64 {
	distanceKm := (height * stepLengthCoefficient * float64(steps)) / mInKm
	return distanceKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	avgSpeed := distance(steps, height) / float64(duration.Hours())
	return avgSpeed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
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

		resultWalk := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activeType, duration.Hours(), distanceKm, avgSpeed, countCaloriesWalk)
		return resultWalk, nil

	case "Бег":
		distanceKm := distance(num, height)
		avgSpeed := meanSpeed(num, height, duration)

		countCaloriesRun, err := RunningSpentCalories(num, weight, height, duration)
		if err != nil {
			return "", err
		}
		
		resultRun := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", activeType, duration.Hours(), distanceKm, avgSpeed, countCaloriesRun)
		return resultRun, nil

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("")
	}

	countCaloriesRun := (weight * meanSpeed(steps, height, duration) * float64(duration.Minutes())) / minInH
	return countCaloriesRun, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("")
	}

	countCaloriesWalk := (weight * meanSpeed(steps, height, duration) * float64(duration.Minutes())) / minInH * walkingCaloriesCoefficient
	return countCaloriesWalk, nil
}
