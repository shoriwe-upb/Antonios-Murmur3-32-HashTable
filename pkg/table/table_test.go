package table

import (
	"log"
	"strconv"
	"testing"
)

const hashLength = 100000

func init() {
	log.Println("Testing Antonio's hash table")
}

func TestStringOnlyInitializationCheckCollision(t *testing.T) {
	log.Printf("Testing Antonio's hash table with %d\n entries", hashLength)
	stringOnlyHash := NewTable(hashLength)
	reference := map[string]int{}
	for i := 0; i < hashLength; i++ {
		key := strconv.Itoa(i)
		value := i
		stringOnlyHash.Set(key, value)
		reference[key] = value
	}
	collisions := 0
	for key, value := range reference {
		getValue, getError := stringOnlyHash.Get(key)
		if getError != nil {
			t.Error(getError)
			return
		}
		if getValue != value {
			collisions++
		}
	}
	if collisions == 0 {
		t.Logf("Test String - Integer: SUCCESS -> No collisions found")
	} else {
		t.Errorf("Test String - Integer: FAILED %d Collisions", collisions)
	}
}

type SampleStruct struct {
	Id   int
	Name string
}

const (
	alicePassword = "Alice'sPassword"
	bobPassword   = "Bob'sPassword"
	carlPassword  = "Carl'sPassword"
)

func TestStruct(t *testing.T) {
	hashTable := NewTable(hashLength)
	alice := &SampleStruct{
		1,
		"Alice",
	}
	bob := &SampleStruct{
		2,
		"Bob",
	}
	carl := &SampleStruct{
		3,
		"Carl",
	}

	hashTable.Set(alice, alicePassword)
	hashTable.Set(bob, bobPassword)
	hashTable.Set(carl, carlPassword)

	storedAlicePassword, _ := hashTable.Get(alice)
	if storedAlicePassword != alicePassword {
		t.Errorf("Test Structs: FAILED -> Colission detected, received \"%s\" but received \"%s\" for Alice's password", alicePassword, storedAlicePassword)
		return
	}
	storedBobPassword, _ := hashTable.Get(bob)
	if storedBobPassword != bobPassword {
		t.Errorf("Test Structs: FAILED -> Colission detected, received \"%s\" but received \"%s\" for Bob's password", bobPassword, storedBobPassword)
		return
	}
	storedCarlPassword, _ := hashTable.Get(carl)
	if storedCarlPassword != carlPassword {
		t.Errorf("Test Structs: FAILED -> Colission detected, received \"%s\" but received \"%s\" for Carl's password", carlPassword, storedCarlPassword)
		return
	}
	t.Logf("Test Structs: SUCCESS")
}
