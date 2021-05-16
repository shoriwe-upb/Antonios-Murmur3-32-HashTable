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
	log.Printf("Testing Antonio's hash table with %d entries\n", hashLength)
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

func TestValueUpdate(t *testing.T) {
	hashTable := NewTable(hashLength)
	for i := 0; i < hashLength; i++ {
		hashTable.Set(i, i)
	}
	reference := map[int]int{}
	for i := 0; i < hashLength; i++ {
		hashTable.Set(i, -i)
		reference[i] = -i
	}
	for i := 0; i < hashLength; i++ {
		value, getError := hashTable.Get(i)
		if getError != nil {
			t.Error(getError)
			return
		}
		if value != reference[i] {
			t.Errorf("Key: %d; Received: %d, Expecting %d", i, value, reference[i])
			return
		}
	}
}

func TestDeletionSmall(t *testing.T) {
	deletionHash := NewTable(hashLength)
	deletionHash.Set(1, 1)
	deletionError := deletionHash.Delete(1)
	if deletionError != nil {
		t.Error(deletionError)
		return
	}
	value, getError := deletionHash.Get(1)
	if getError == nil {
		t.Errorf("Key %d was not deleted, returned %v", 1, value)
		return
	}
	// Finally re set the value
	deletionHash.Set(1, 1)
	value, getError = deletionHash.Get(1)
	if getError != nil {
		t.Error(getError)
		return
	}
	if value != 1 {
		t.Errorf("Received %v Expecting %d", value, 1)
		return
	}

}

func TestDeletionBig(t *testing.T) {
	deletionHash := NewTable(hashLength)
	for i := 0; i < hashLength; i++ {
		deletionHash.Set(i, i)
	}
	for i := 0; i < hashLength; i++ {
		deletionError := deletionHash.Delete(i)
		if deletionError != nil {
			t.Error(deletionError)
			return
		}
	}
	for i := 0; i < hashLength; i++ {
		value, getError := deletionHash.Get(i)
		if getError == nil {
			t.Errorf("Key %d was not deleted, returned %v", i, value)
			return
		}
	}
}
