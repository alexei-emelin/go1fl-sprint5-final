package trainings

import (
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
	"fmt"
	"strings"
	"strconv"
	"time"
)

// создайте структуру Training
type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal
}


// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {

	parts := strings.Split(datastring,",")
	if len(parts) != 3 {
		return fmt.Errorf("некорректные исходные данные")
	}
	steps, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return fmt.Errorf("некорректное число шагов: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("не верное количество шагов: %w", err)
	}
	t.Steps = steps


	activity := strings.TrimSpace(parts[1])
	if activity != "Бег" && activity != "Ходьба" {
		return fmt.Errorf("неизвестный тип тренировки")
	}
	t.TrainingType = activity


	duration, err := time.ParseDuration(strings.TrimSpace(parts[2]))
	if err != nil {
		return fmt.Errorf("ошибка парсинга длительности: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("некорректное время длительности: %w", err)
	}
	t.Duration = duration
	return nil

}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	if t.Duration <= 0 {
		return "", fmt.Errorf("продолжительность тренировки должна быть больше 0")
	}

	switch t.TrainingType {
	case "Ходьба":
		dist := spentenergy.Distance(t.Steps)
		if dist <= 0 {
			return "", fmt.Errorf("дистанция должна быть больше 0")
		}
		speed := spentenergy.MeanSpeed(t.Steps, t.Duration)
		calories := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		return fmt.Sprintf(
			"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
			t.TrainingType, t.Duration.Hours(), dist, speed, calories,
		), nil
	case "Бег":
		dist := spentenergy.Distance(t.Steps)
		if dist <= 0 {
			return "", fmt.Errorf("Дистанция должна быть больше 0")
		}
		speed := spentenergy.MeanSpeed(t.Steps, t.Duration)
		calories := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
		return fmt.Sprintf(
			"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
			t.TrainingType, t.Duration.Hours(), dist, speed, calories,
		), nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}