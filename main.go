package main

import (
	"time"
	// extra empty line, gofmt cannot solve
	"fmt"       // BadImportOrdering: "fmt" should be grouped with other standard library imports
	"io/ioutil" // Deprecated api: io/ioutil has been deprecated since Go 1.19
	"math/rand"
	"os"
)

// BadFormatting: Vars often declared in the behind const vars.
var globalData int //  GlobalVariables: Using global variables can lead to unpredictable behavior and hard-to-track bugs

// BadNaming: const var should be max capped.
// BadFormatting: const vars can be groupped together.
const a = 1
const b = 2

type user struct {
	id        int       `json: "id"` // BadSyntax: should be `json:"id"`(no extra space)
	nameStr   string    //  BadNaming: Repetitive name, can be `name`
	data      *userData //  InappropriatePointerUse: Unnecessary use of pointers can increase complexity
	LinkedUrl string    //  BadNaming: URL should be URL or url.
}

type userData struct {
	description string
	detailsID   int
}

// LongFunction: Functions should be short and focused on a single task
// BadFunctionName: Function names should be descriptive and in MixedCaps.
func processdata(u *user, params ...string) { // 7. LongParameterList: Functions should avoid long parameter lists
	if len(params) > 10 { // MagicNumbers: Avoid using magic numbers without explanation
		fmt.Println("Too many parameters")
		return
	}

	// NonClosedFile: Not deferring file closure can lead to resource leaks
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err) // WrongLogging: Should use a logging package instead of fmt for errors
		return
	}
	defer file.Close() // Fix for NonClosedFile

	content, _ := ioutil.ReadAll(file) // IgnoredError: Ignoring errors can lead to silent failures
	fmt.Println("File content:", string(content))

	// InefficientResourceUse: Creating a new rand.Source every time is inefficient
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	fmt.Println("Random number:", rnd.Intn(100))

	// 13. UnhandledConcurrency: Launching a goroutine without handling its execution or errors
	go func() {
		fmt.Println("Asynchronous operation")
		// Assume more complex logic...
	}()

	// BadFormatting: Inconsistent spacing and indentation
	if u.id > 100 {
		fmt.Println("ID is high")
	} else {
		fmt.Println("ID is normal")
	}

	// PoorlyDesignedStruct: userData struct is not well-designed for its intended use
	u.data = &userData{description: "", detailsID: 1} // BadInitialization: description has zero value ""

	dataSlice := []int{} // BadInitialization & BadNaming(Repetitive), should be `var data []int`
	for i := 0; i < 100; i++ {
		dataSlice = append(dataSlice, i)
	}

	// BadChannelUse: Incorrect use of channels can lead to deadlocks or inefficient data processing
	ch := make(chan int, 10)
	ch <- 1 // Sending to channel without receiving
	// Channel not closed or received from

	// RedundantCode: Redundant or unnecessary code that could be simplified
	if u.nameStr == "" {
		fmt.Println("Name is empty")
	} else {
		fmt.Println("Name is not empty")
	}
	unreachableCode()

	// InconsistentErrorHandling: Mixing error handling styles leads to inconsistent code
	if err := doSomething(1, 2, "", 3, "", 4, "", 5); err != nil {
		panic(err) // Using panic instead of structured error handling
	}
}

func unreachableCode() {
	return
	fmt.Println("This will never be called")
}

func uncalledFunc() {
	return
}

// Long Function Parameter list.
func doSomething(p1 int, p2 int, p3 string, p4 int, p5 string, p6 int, p7 string, p8 int) error {
	return nil
}

func main() {
	u := &user{id: 1, nameStr: "John Doe"}
	processdata(u, "param1", "param2", "param3")
}
