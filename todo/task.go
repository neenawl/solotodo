package todo

import "time"

type Task struct {
	Title  string
	Text   string
	IsDone bool

	CreatedAt time.Time
	DoneAt    *time.Time
}

func NewTask(title string, text string) Task {
	task := Task{
		Title: title,
		Text:  text,

		IsDone:    false,
		CreatedAt: time.Now(),
		DoneAt:    nil,
	}

	return task
}

func (t *Task) Done() {
	t.IsDone = true

	now := time.Now()
	t.DoneAt = &now
}
