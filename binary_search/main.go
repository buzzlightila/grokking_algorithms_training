package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func generateOrderedRandomList(size int) []int {
	randomList := make([]int, size)

	for i := 0; i < size; i++ {
		// Generate a random number between 1 and 100 (you can adjust the range as needed)
		randomNumber := rand.Intn(1000) + 1
		randomList[i] = randomNumber
	}

	sort.Ints(randomList)
	return randomList
}

func getRandomNumber(list []int) int {
	// Generate a random index within the range of the list
	randomIndex := rand.Intn(len(list))

	// Return the element at the random index
	return list[randomIndex]
}

func simpleSearch(item int, list []int) int {
	var attemptsToFind int

	for index, v := range list {
		if v == item {
			attemptsToFind = index
		}
	}
	return attemptsToFind
}

func binarySearch(item int, list []int) int {
	listLen := len(list)
	lowNumber := 0
	highNumber := listLen - 1
	attemptsToFind := 0

	for lowNumber <= highNumber {
		midNumber := (lowNumber + highNumber) / 2
		likelyNumber := list[midNumber]

		if likelyNumber == item {
			return attemptsToFind
		} else if likelyNumber > item {
			highNumber = midNumber - 1
		} else {
			lowNumber = midNumber + 1
		}
		attemptsToFind += 1
	}
	return attemptsToFind
}

/*
A ideia é comparar uma procura simples e uma pesquisa binária em uma lista de números ordenados.
*/
func main() {
	numberList := generateOrderedRandomList(10000000)
	randomNumber := getRandomNumber(numberList)
	fmt.Printf("Random number: %d\n", randomNumber)

	// Simple Search
	fmt.Println("Executing a simple search...")
	attemptsTime := simpleSearch(randomNumber, numberList)
	fmt.Printf("Attempts time to find number: %d\n", attemptsTime)

	//	Binary Search
	fmt.Println("Executing a binary search...")
	attemptsTime = binarySearch(randomNumber, numberList)
	fmt.Printf("Attempts time to find number: %d\n", attemptsTime)
}
