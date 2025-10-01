package daysteps

import (
	"errors"
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
	dataSlice := strings.Split(data, ",")
	if len(dataSlice) != 2 {
		err := errors.New("неверный формат входных данных parsePackage")
		log.Println(err)
		return 0, 0, err
	}

	steps, err := strconv.Atoi(dataSlice[0])
	if err != nil || steps <= 0 {
		err := errors.New("ошибка преобразования шагов parsePackage")
		log.Println(err)
		return 0, 0, err
	}

	trainDur, err := time.ParseDuration(dataSlice[1])
	if err != nil || trainDur.Seconds() <= 0 {
		err := errors.New("ошибка преобразования длительности тренировки parsePackage")
		log.Println(err)
		return 0, 0, err
	}
	return steps, trainDur, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, trainDur, err := parsePackage(data)
	if err != nil || steps <= 0 || trainDur.Seconds() <= 0 {
		log.Println(err)
		return ""
	}
	distance := (stepLength * float64(steps)) / float64(mInKm)
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, trainDur)
	if err != nil {
		log.Println(err)
		return ""
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, calories)

}
