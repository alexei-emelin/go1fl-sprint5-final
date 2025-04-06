package daysteps

import (
	"strings"
	"strconv"
	"time"
	"fmt"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps int
	Duration time.Duration
	personaldata.Personal
}


// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring,",")
	if len(parts) != 2 {
		return fmt.Errorf("некорректные исходные данные")
	}
	ds.Steps, err = strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return fmt.Errorf("некорректное число шагов: %w", err)
	}
	if ds.Steps <= 0 {
		return fmt.Errorf("некорректное количество шагов: %w", err)
	}
	ds.Duration, err = time.ParseDuration(strings.TrimSpace(parts[1]))
	if err != nil {
		return fmt.Errorf("ошибка парсинга длительности: %w", err)
	}
	if ds.Duration <= 0 {
		return fmt.Errorf("некорректное время длительности: %w", err)
	}
	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	dist := spentenergy.Distance(ds.Steps)
	calories := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, calories), nil
}
