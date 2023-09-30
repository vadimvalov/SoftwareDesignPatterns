package main

import (
	"fmt"
	"time"
)

// Subject представляет субъект (наблюдаемый объект).
type Subject struct {
	observers []Observer
	state     string
}

// Observer определяет интерфейс для наблюдателя.
type Observer interface {
	Update(string)
}

// ConcreteObserver представляет конкретного наблюдателя.
type ConcreteObserver struct {
	name string
}

// Update обновляет состояние наблюдателя.
func (o *ConcreteObserver) Update(state string) {
	fmt.Printf("Студент %s получил новое уведомление: %s\n", o.name, state)
}

// Attach добавляет наблюдателя к субъекту.
func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

// Notify уведомляет всех наблюдателей об изменении состояния.
func (s *Subject) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.state)
	}
}

// SetState устанавливает новое состояние субъекта и уведомляет наблюдателей.
func (s *Subject) SetState(newState string) {
	s.state = newState
	s.Notify()
}

func main() {
	// Создаем субъект
	subject := &Subject{}

	// Добавляем наблюдателей
	observer1 := &ConcreteObserver{name: "Вадик Валов"}
	observer2 := &ConcreteObserver{name: "Рома рома роман мущина всей моей жизни"}
	subject.Attach(observer1)
	subject.Attach(observer2)

	// Устанавливаем новое состояние субъекта
	subject.SetState("новое домашнее задание")

	// Подождем немного, чтобы увидеть уведомления
	time.Sleep(time.Second)
}
