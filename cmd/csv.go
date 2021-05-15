package main

import (
	"fmt"
	"github.com/shoriwe-upb/Antonios-Murmur3-32-HashTable/pkg/table"
	"os"
	"strconv"
)

var hashTableSize int

var intReference = map[int]int{}
var stringReference = map[string]int{}
var hashTable *table.Table

func init() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintf(os.Stderr, "Usage: %s HASH_TABLE_SIZE\n", os.Args[0])
		os.Exit(0)
	}
	number, transformationError := strconv.Atoi(os.Args[1])
	if transformationError != nil {
		_, _ = fmt.Fprintln(os.Stderr, "HASH_TABLE_SIZE must be a integer number")
		os.Exit(1)
	}
	hashTableSize = number
	hashTable = table.NewTable(uint32(hashTableSize))
	for i := 0; i < (hashTableSize); i++ {
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
		stringI := strconv.Itoa(i)
		antonioIntValue, getError := hashTable.Get(i)
		if getError != nil {
			panic(getError)
		}
		var antonioStringValue interface{}
		antonioStringValue, getError = hashTable.Get(stringI)

		goIntValue := intReference[i]
		goStringValue := stringReference[stringI]
		fmt.Printf("%d, %d, %s, %d, %d, %d, %s, %d, %t\n", i, antonioIntValue, stringI, antonioStringValue, i, goIntValue, stringI, goStringValue, goIntValue != antonioIntValue || goStringValue != antonioStringValue)
	}
}
