package queue

import (
	"errors"
	"fmt"
)

type CircularQueue struct {
	values []int
	front  int
	rear   int
	size   int
}

// NewCircularQueue Новая очередь
func NewCircularQueue(size int) (CircularQueue, error) {
	var q = new(CircularQueue)
	if size <= 0 {
		return *q, errors.New("size must be greater than zero")
	}
	q.values = make([]int, size)
	q.front = -1
	q.rear = -1
	q.size = size
	return *q, nil
}

// Push Добавление элемента
func (q *CircularQueue) Push(i int) error {
	if q.Full() {
		return errors.New("queue is full")
	}
	// Инициализируем front если массив пустой
	if q.front == -1 {
		q.front = 0
	}
	// Задаем rear на позицию добавляемого элемента
	q.rear = (q.rear + 1) % q.size
	q.values[q.rear] = i

	return nil
}

// Pop Удаление элемента
func (q *CircularQueue) Pop() error {
	if q.Empty() {
		return errors.New("queue is empty")
	}
	if q.front == q.rear {
		// Если в очереди всего один элемент,
		// то мы сбрасываем очередь
		q.front = -1
		q.rear = -1
	} else {
		q.values[q.front] = 0
		q.front = (q.front + 1) % q.size
	}
	return nil
}

// Empty проверяем пуста ли очередь
func (q *CircularQueue) Empty() bool {
	// Если пустая
	if q.front == -1 {
		return true
	}
	// Если непустая
	return false
}

// Full проверяем заполнена ли очередь
func (q *CircularQueue) Full() bool {
	// Если у нас Front указывает на место куда бы мы
	// хотели положить новый элемент, то вернем true
	// [F][X][X][X][X][X][X][X][X][R]
	//  0                          9
	// 0 == (9+1)%10 → 0 == 0 → true
	if q.front == (q.rear+1)%q.size {
		return true
	}
	return false
}

func (q *CircularQueue) PrintQueue() {
	fmt.Println(q.values)
}
