package main

import (
	"fmt"
	"sync"
	"time"
)

type LogMessage struct {
	Level   string
	Content string
}

// LogStats maintains the count of log messages by level
type LogStats struct {
	sync.Mutex
	Count map[string]int
}

func main() {
	messages := make(chan LogMessage, 10)
	done := make(chan bool)

	stats := &LogStats{Count: make(map[string]int)}

	// Simulate input of log messages
	go inputLogMessages(messages)

	// Process messages and update statistics concurrently
	go processLogMessages(messages, done)
	go updateStatistics(messages, stats)

	// Wait for processing to complete
	<-done

	// Display collected statistics
	displayStatistics(stats)
}

// inputLogMessages simulates the input stage of the log system
func inputLogMessages(messages chan<- LogMessage) {
	for i := 0; i < 100; i++ {
		message := LogMessage{Level: "INFO", Content: "This is a log message"}
		messages <- message
	}
	close(messages)
}

// processLogMessages processes incoming log messages
func processLogMessages(messages <-chan LogMessage, done chan<- bool) {
	for message := range messages {
		// Simulate message processing
		processMessage(message)
		// Here you would typically send the message to a logging output, not shown for simplicity
	}
	done <- true
}

// updateStatistics updates log message statistics
func updateStatistics(messages <-chan LogMessage, stats *LogStats) {
	for message := range messages {
		stats.Lock()
		if _, exists := stats.Count[message.Level]; exists {
			stats.Count[message.Level]++
		} else {
			stats.Count[message.Level] = 1
		}
		stats.Unlock()
	}
}

// displayStatistics prints out the log statistics
func displayStatistics(stats *LogStats) {
	stats.Lock()
	fmt.Println("Log Statistics:")
	for level, count := range stats.Count {
		fmt.Printf("Level %s: %d messages\n", level, count)
	}
	stats.Unlock()
}

// processMessage simulates a processing delay
func processMessage(message LogMessage) {
	time.Sleep(time.Millisecond * 10)
}
