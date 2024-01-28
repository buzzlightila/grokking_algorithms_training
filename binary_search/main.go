package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func generateOrderedRandomList(size int) []int {
	randomList := make([]int, size)

	for i := 0; i < size; i++ {
		randomNumber := rand.Intn(1000) + 1
		randomList[i] = randomNumber
	}

	sort.Ints(randomList)
	return randomList
}

func getRandomNumber(list []int) int {
	randomIndex := rand.Intn(len(list))

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
	// Esta função pode ser encontrada no livro "Entendendo algoritmos". - Capítulo 1, página 27
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
A ideia é utilizar dois algoritmos um de procura simples e um de pesquisa binária.
Com isso, comparar o número de tentativas para procurar um número aleatório.
*/
func main() {
	numberList := generateOrderedRandomList(10000000)
	randomNumber := getRandomNumber(numberList)
	fmt.Printf("Número aleatório: %d\n", randomNumber)

	// Simple Search
	fmt.Println("Executando uma pesquisa simples...")
	attemptsTime := simpleSearch(randomNumber, numberList)
	fmt.Printf("Número de tentativas para buscar um número: %d\n", attemptsTime)

	//	Binary Search
	fmt.Println("Executando pesquisa binária...")
	attemptsTime = binarySearch(randomNumber, numberList)
	fmt.Printf("Número de tentativas para buscar um número: %d\n", attemptsTime)
}
