package tests

import (
	"testing"

	"github.com/tmb-piXel/LearnEnglishBot/pkg/telegram"
)

func TestCheckAnswer(t *testing.T) {
	tables := []struct {
		correct string
		answer  string
		result  bool
	}{
		{"обильный", " Обильный ", true},
		{" примерно ", "ПриМерно ", true},
		{" выполнить / завершить ", " Завершить", true},
		{" красивый ", " некрасивый", false},
		{" красивый ", " красив", false},
	}

	for _, table := range tables {
		total := telegram.CheckAnswer(table.correct, table.answer)
		if total != table.result {
			t.Errorf(`%s+%s was incorrect got: %t, want: %t.`,
				table.correct, table.answer, total, table.result)
		}
	}
}
