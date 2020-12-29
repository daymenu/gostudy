package main

func main() {
	// use wire
	event := InitializeEvent()
	event.Start()

	// don't use wire
	// message := NewMessage()
	// greeter := NewGreeter(message)
	// event := NewEvent(greeter)
	// event.Start()
}
