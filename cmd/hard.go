package main

import (
	"DataStructures/HashTable/pkg/table"
	"fmt"
	"strconv"
)

const hashTableSize = 1000000

var intReference = map[int]int{}
var stringReference = map[string]int{}
var hashTable = table.NewTable(hashTableSize * 2)

func init() {
	for i := 0; i < hashTableSize; i++ {
		value := i - 1
		stringI := strconv.Itoa(i)
		intReference[i] = value
		stringReference[stringI] = value

		hashTable.Set(i, value)
		hashTable.Set(stringI, value)
	}
}

func main() {
	fmt.Println("Antonio's Int HashTable Key, Antonio's  Int HastTable Value, Antonio's String HashTable Key, Antonio's  String HastTable Value, Go's Int HashTable Key, Go's Int HashTable Value, Go's String HashTable Key, Go's String HashTable Value, Collision")
	for i := 0; i < hashTableSize; i++ {
		stringI :=  strconv.Itoa(i)
		antonioIntValue, getError := hashTable.Get(i)
		if getError != nil {
			panic(getError)
		}
		var  antonioStringValue interface{}
		antonioStringValue, getError = hashTable.Get(stringI)

		goIntValue := intReference[i]
		goStringValue := stringReference[stringI]
		fmt.Printf("%d, %d, %s, %d, %d, %d, %s, %d, %t\n", i, antonioIntValue, stringI, antonioStringValue, i, goIntValue, stringI, goStringValue, goIntValue != antonioIntValue || goStringValue != antonioStringValue)
	}
}
