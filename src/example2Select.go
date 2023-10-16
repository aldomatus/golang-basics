package main

import (
	"fmt"
	"time"
)

func trafficLight(color string, duration time.Duration, lightChan chan<- string, stopChan <-chan bool) {
	for {
		select {
		case <-stopChan:
			close(lightChan)
			return
		default:
			lightChan <- color
			time.Sleep(duration)
		}
	}
}

func main() {
	redLight := make(chan string)
	yellowLight := make(chan string)
	greenLight := make(chan string)
	stopSignal := make(chan bool)

	go trafficLight("Red", 5*time.Second, redLight, stopSignal)
	go trafficLight("Yellow", 3*time.Second, yellowLight, stopSignal)
	go trafficLight("Green", 6*time.Second, greenLight, stopSignal)

	for {
		select {
		case red := <-redLight:
			fmt.Println("Traffic Light:", red)
		case yellow := <-yellowLight:
			fmt.Println("Traffic Light:", yellow)
		case green := <-greenLight:
			fmt.Println("Traffic Light:", green)
		case <-stopSignal:
			fmt.Println("Traffic Light System Stopped.")
			return
		}
	}
}
