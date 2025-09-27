package daysteps

import (
	"errors"
	"fmt"
	"internal/spentcalories"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	dataSlice := strings.Split(data, ",")
	if len(dataSlice) != 2 {
		return 0, 0, errors.New("Неверный формат входных данных parsePackage")
	}

	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil || steps <= 0 {
		return 0, 0, errors.New("Ошибка преобразования шагов: " + err.Error())
	}

	trainDur, err := time.ParseDuration(dataSlice[1])
	if err != nil {

		return 0, 0, errors.New("Ошибка преобразования длительности тренировки: " + err.Error())
	}
	return steps, trainDur, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, trainDur, err := parsePackage(data)
	if err != nil || steps <= 0 {
		return ""
	}
	distance := (stepLength * float64(steps)) / mInKm
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, trainDur)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distance, calories)
}
