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
	text := strings.Split(data, ",")
	if len(text) != 2 {
		return 0, 0, fmt.Errorf("wrong input: wrong number of args")
	}
	stepsNumber, err := strconv.Atoi(text[0])
	if err != nil {
		return 0, 0, err
	}
	if stepsNumber <= 0 {
		return 0, 0, fmt.Errorf("wrong input: steps <= 0")
	}
	duration, err := time.ParseDuration(text[1])
	if err != nil {
		return 0, 0, err
	}
	if duration <= 0 {
		return 0, 0, fmt.Errorf("wrong input: duration <= 0")
	}
	return stepsNumber, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	stepsNumber, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if stepsNumber <= 0 {
		return ""
	}
	distance := (float64(stepsNumber) * stepLength) / float64(mInKm)
	caloriesNumber, err := spentcalories.WalkingSpentCalories(stepsNumber, height, weight, duration)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", stepsNumber, distance, caloriesNumber)
}
