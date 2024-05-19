package api

import (
	"ccvalidator/luhn"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type WorkerPool struct {
	tasks       chan Task
	workerCount int
	wg          *sync.WaitGroup
}

type Task struct {
	CardNumber string
	Response   chan bool
}

type Request struct {
	CardNumber string `json:"card_number"`
}

var Wp *WorkerPool

func InitWorkerPool(workerCount int) {
	Wp = NewWorkerPool(workerCount)
	go Wp.Run()
}

func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		tasks:       make(chan Task, 100),
		workerCount: workerCount,
		wg:          &sync.WaitGroup{},
	}
}

func (wp *WorkerPool) Run() {
	for i := 0; i < wp.workerCount; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
	wp.wg.Wait()
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for task := range wp.tasks {
		fmt.Printf("Worker %d processing task\n", id)
		task.Response <- luhn.Validate(task.CardNumber)
	}
}

func (wp *WorkerPool) AddTask(task Task) {
	wp.tasks <- task
}

func (wp *WorkerPool) Close() {
	close(wp.tasks)
}

func HandleRequestWithWorkerPool(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	if req.CardNumber == "" {
		http.Error(w, "card_number is required", http.StatusBadRequest)
		return
	}

	responseChan := make(chan bool)
	task := Task{
		CardNumber: req.CardNumber,
		Response:   responseChan,
	}

	Wp.AddTask(task)
	isValid := <-responseChan

	result := "invalid"
	if isValid {
		result = "valid"
	}

	fmt.Fprintf(w, "Card number %s is %s", req.CardNumber, result)
}
