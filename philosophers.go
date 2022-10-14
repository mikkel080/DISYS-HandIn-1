package main

import (
	"fmt"
	"time"
)

func ph(name string, rightHand chan bool, leftHand chan bool) {
	fmt.Println(name + " has been seated.") //Philosopher gorutine succesfully initiated
	var hunger = 3                          //Set amount of times philosopher has to eat before leaving the table

	for hunger > 0 {
		time.Sleep(6000)  //Think for 6 seconds before eating
		rightHand <- true //Ask if right fork is available
		<-rightHand       //Wait until right fork is available
		leftHand <- true  //Ask if left fork is available
		<-leftHand        //Wait until left fork is available

		fmt.Println(name + " has eaten.") //Philosopher has succesfully eaten with two forks
		hunger--                          //Lower amount of times the philosopher has to eat by 1

		rightHand <- true //Put down right fork
		leftHand <- true  //put down left fork
	}

	fmt.Println(name + " has left the table") //Philosopher has succesfully eaten 3 times and leaves the table
}

func fork(right chan bool, left chan bool) {

	for {
		select {
		case <-right: //Fork waits to be picked up
			right <- true //When fork is picked up it sends a message that it is available
			<-right       //Wait for message that the fork has been put down
		case <-left: //Fork waits to be picked up
			left <- true //When fork is picked up it sends a message that it is available
			<-left       //Wait for message that the fork has been put down
		}
	}
}

func main() {
	fork1right := make(chan bool)
	fork1left := make(chan bool)
	go fork(fork1right, fork1left) //First Fork

	fork2right := make(chan bool)
	fork2left := make(chan bool)
	go fork(fork2right, fork2left) //Second Fork

	fork3right := make(chan bool)
	fork3left := make(chan bool)
	go fork(fork3right, fork3left) //Third Fork

	fork4right := make(chan bool)
	fork4left := make(chan bool)
	go fork(fork4right, fork4left) //Fourth Fork

	fork5right := make(chan bool)
	fork5left := make(chan bool)
	go fork(fork5right, fork5left) //Fifth Fork

	go ph("Philosopher 1", fork1left, fork2right) //First Philosopher
	go ph("Philosopher 2", fork2left, fork3right) //Second Philosopher
	go ph("Philosopher 3", fork3left, fork4right) //Third Philosopher
	go ph("Philosopher 4", fork4left, fork5right) //Fourth Philosopher
	go ph("Philosopher 5", fork1right, fork5left) //Fifth Philosopher sits with "hands crossed" to avoid a deadlock.
	time.Sleep(10 * time.Second)                  //give the program some time to actually run.
}
