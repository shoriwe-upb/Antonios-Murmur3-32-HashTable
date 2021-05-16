# Antonio's Hash Table

Source code for my implementation of a Hash Table based in the 32bits version of the **`Murmur3 Hash`**  function as
object hashing mechanism.

Apparently it doesn't have collisions but is highly probable that I used a very small test input, the only thing a I
could conclude is that by the time it starts getting collisions, became more slower, since it end navigating very large
collision lists.

## Usage

### Test

First of all clone the repository

```shell
git clone github.com/shoriwe-upb/Antonios-Murmur3-32-HashTable
```

#### Behavior test

A quick test can be executed by just running... In the root of the cloned repository.

```shell
go test -count=1 ./...
```

##### Output

```csv
?       DataStructures/HashTable/cmd    [no test files]
ok      DataStructures/HashTable/pkg/hash       0.091s
ok      DataStructures/HashTable/pkg/table      0.278s
```

#### CSV table test

This test will generate a CSV file in case you want to manually check it the collisions.

You can execute it by:

1. Changing your folder to `cmd` of the repository:

```shell
cd cmd
```

2. Executing:

```shell
go run csv.go HASH_TABLE_SIZE
```

##### Example

```shell
go run csv.go 6
```

##### Output

| Antonio's Int HashTable Key | Antonio's  Int HastTable Value | Antonio's String HashTable Key | Antonio's  String HastTable Value | Go's Int HashTable Key | Go's Int HashTable Value | Go's String HashTable Key | Go's String HashTable Value | Collision |
| --------------------------- | ------------------------------ | ------------------------------ | --------------------------------- | ---------------------- | ------------------------ | ------------------------- | --------------------------- | --------- |
| 0                           | -1                             | 0                              | -1                                | 0                      | -1                       | 0                         | -1                          | false     |
| 1                           | 0                              | 1                              | 0                                 | 1                      | 0                        | 1                         | 0                           | false     |
| 2                           | 1                              | 2                              | 1                                 | 2                      | 1                        | 2                         | 1                           | false     |
| 3                           | 2                              | 3                              | 2                                 | 3                      | 2                        | 3                         | 2                           | false     |
| 4                           | 3                              | 4                              | 3                                 | 4                      | 3                        | 4                         | 3                           | false     |
| 5                           | 4                              | 5                              | 4                                 | 5                      | 4                        | 5                         | 4                           | false     |

### As a library

From your program/module root simply execute

```shell
go get github.com/shoriwe-upb/Antonios-Murmur3-32-HashTable
```

Your **`go.mod`** and **`go.sum`** should now contain the library.

Use the **`pkg`** package since there is the implementation.

```go
package main

import (
	"github.com/shoriwe-upb/Antonios-Murmur3-32-HashTable/table"
)

func main() {
	hashTable := table.NewTable()
	hashTable.Set("Hello", "World")
	value, getError := hashTable.Get("Hello")
	if getError != nil {
		panic(getError)
	}
	fmt.Println(value)
}
```

## Behavior

### Overview

The base behavior consist in approaching Go's string formatting of object to generate string keys independent of the
type. From the resulting string, then is calculated a number (of size 32 bits) thanks to
the **`Murmur3 Hash algorithm`** which is mod divided by the length of the hash table array to determine the relative
index. After that the element in the index is a list of entries which prevents keys from colliding between them.

#### In resume:

1. Format the object in a string with **`"%+v-%T"`**, which means, format every detail of the object ( **`%+v`**) and
   format its type too (**`%T`**).
2. Of the resulting string, calculate its **`Murmur3`** 32bits hash (which will be always a **uint32** number).
3. To get the index just mod divide the Hash result between the length of the base array of the hash
   table. **`Murmur3_32Bits(generatedKey) % tableArrayLength`**.
4. If the key already existing in the indexed list, update its value, if not append a new entry relating to the new key.

### Details

More detailed explanation of the indexing algorithm

#### Key creation

Every object received as key is formatted in **`"%+v-%T"`**

- **`%+v`** Formatting means, give me every detail of the object.
- **`%T`** Formatting means, give the type of the object.

#### Maximum size

Thanks to the `Murmur3` 32bits version hash function, the maximum size that the hash table will always work
is **`4294967296`** elements, or at least is what my knowledge let me understand.

#### Collision safe

Collisions are handled by using in the array indexes lists which contains the raw-key and its value, this way when two
different keys has the same computed index, there is less chance of getting collision since the keys are added
independently to this list.

## Important

Don't use this in production software, since it's performance is not great enough.

## References

| Title      | Source                                   |
| ---------- | ---------------------------------------- |
| MurmurHash | https://en.wikipedia.org/wiki/MurmurHash |

