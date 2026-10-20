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
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	text := strings.Split(data, ",")
	if len(text) != 3 {
		return 0, "", 0, fmt.Errorf("wrong input: wrong number of args")
	}
	stepsNumber, err := strconv.Atoi(text[0])
	if err != nil {
		return 0, "", 0, err
	}
	if stepsNumber <= 0 {
		return 0, "", 0, fmt.Errorf("wrong input: steps <= 0")
	}
	duration, err := time.ParseDuration(text[2])
	if err != nil {
		return 0, "", 0, err
	}
	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("zero duration")
	}
	return stepsNumber, text[1], duration, nil
}

func distance(steps int, height float64) float64 {
	oneStepLength := height * stepLengthCoefficient
	return (float64(steps) * oneStepLength) / float64(mInKm)
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	distance := distance(steps, height)
	return distance / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	stepsNumber, trainingType, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", err
	}
	distance := distance(stepsNumber, height)
	speed := meanSpeed(stepsNumber, height, duration)
	var caloriesNumber float64
	switch trainingType {
	case "Бег":
		caloriesNumber, err = RunningSpentCalories(stepsNumber, weight, height, duration)
	case "Ходьба":
		caloriesNumber, err = WalkingSpentCalories(stepsNumber, weight, height, duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration.Hours(), distance, speed, caloriesNumber), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("wrong args")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("wrong args")
	}
	speed := meanSpeed(steps, height, duration)
	cals := (weight * speed * duration.Minutes()) / 60
	return cals * walkingCaloriesCoefficient, nil
}
