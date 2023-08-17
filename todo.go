package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(number int) error {
	ls := *t
	if number <= 0 || number > len(ls) {
		return errors.New("invalid syntax")
	}

	ls[number-1].Done = true
	ls[number-1].CompletedAt = time.Now()

	return nil
}

func (t *Todos) Delete(number int) error {
	ls := *t
	if number <= 0 || number > len(ls) {
		return errors.New("invalid syntax")
	}
	*t = append(ls[:number-1], ls[number:]...)

	return nil
}

func (t *Todos) Load(fileName string) error {
	file, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) List() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignCenter, Text: "CreatedAt"},
			{Align: simpletable.AlignCenter, Text: "CompleteAt"},
		},
	}

	var cells [][]*simpletable.Cell
	for idx, element := range *t {
		idx++
		task := blue(element.Task)
		done := blue("no")
		if element.Done {
			task = green(fmt.Sprintf("\u2705 %s", element.Task))
			done = green("yes")
		}
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: task},
			{Text: done},
			{Text: element.CreatedAt.Format(time.RFC822)},
			{Text: element.CompletedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending todos", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todos) Store(fileName string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}

func (t *Todos) CountPending() int {
	total := 0
	for _, element := range *t {
		if !element.Done {
			total++
		}
	}

	return total
}
