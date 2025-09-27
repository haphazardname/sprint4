package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	//lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	dataSlice := strings.Split(data, ",")
	if len(dataSlice) != 3 {
		return 0, "", 0, errors.New("неверный формат входных данных parseTraining")
	}

	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil || steps <= 0 {
		return 0, "", 0, errors.New("ошибка преобразования шагов")
	}

	typeTrain := dataSlice[1]

	trainDur, err := time.ParseDuration(dataSlice[2])
	if err != nil || trainDur.Seconds() <= 0 {
		return 0, "", 0, errors.New("ошибка преобразования длительности тренировки")
	}

	return steps, typeTrain, trainDur, nil
}

func distance(steps int, height float64) float64 {
	if steps > 0 && height > 0 {
		return ((float64(stepLengthCoefficient) * height) * float64(steps)) / float64(mInKm)
	}
	return 0
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration > 0 && steps > 0 && height > 0 {
		return distance(steps, height) / duration.Hours()
	}
	return 0
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, typeTrain, trainDur, err := parseTraining(data)
	if err != nil {
		//log.Println(err)
		return "", err
	}

	if weight <= 0 || height <= 0 {
		return "", errors.New("некорректные входные  данные TrainingInfo")
	}

	var calories float64
	switch typeTrain {
	case "Ходьба":
		calories, err = WalkingSpentCalories(steps, weight, height, trainDur)
	case "Бег":
		calories, err = RunningSpentCalories(steps, weight, height, trainDur)
	default:
		err = errors.New("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeTrain, trainDur.Hours(), distance(steps, height), meanSpeed(steps, height, trainDur), calories), nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные входные параметры RunningSpentCalories")
	}

	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / float64(minInH), nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("некорректные входные параметры WalkingSpentCalories")
	}
	meanSpeed := meanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return ((weight * meanSpeed * durationInMinutes) / float64(minInH)) * float64(walkingCaloriesCoefficient), nil

}
