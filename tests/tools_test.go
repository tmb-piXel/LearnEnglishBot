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

func TestContains(t *testing.T) {
	tables := []struct {
		m   map[int64]bool
		e   int64
		out bool
	}{
		{map[int64]bool{21321: true, 123123: false}, 21321, true},
		{map[int64]bool{21321: true, 123123: false}, 321321, false},
		{map[int64]bool{21321: true, 123123: false}, 123123, true},
	}

	for _, table := range tables {
		total := telegram.Contains(table.m, table.e)
		if total != table.out {
			t.Errorf(`%#v+%d was incorrect got: %t, want: %t.`,
				table.m, table.e, total, table.out)
		}
	}
}
