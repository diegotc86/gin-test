package bd

type Task struct {
	Id           string `json:"id"`
	Body         string `json:"body"`
	Completed    bool   `json:"completed"`
	Created_date string `json:"created_date"`
}

var Tasks = []Task{
	{Id: "10", Body: "Task 1", Completed: false, Created_date: "2022-09-09T15:00:00+07:00"},
	{Id: "11", Body: "Task 2", Completed: true, Created_date: "2022-09-10T15:00:00+07:00"},
	{Id: "12", Body: "Task 3", Completed: true, Created_date: "2022-09-11T15:00:00+07:00"},
}
