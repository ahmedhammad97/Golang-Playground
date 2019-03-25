package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"log"
)

func main(){
	frequency := make(map[int]int)
	fmt.Println("Start typing numbers from 0 to 9...\n")
	reader := bufio.NewReader(os.Stdin)
	for  {
		location, _ := reader.ReadString('\n')
		location = location[:len(location)-1]
		NumLocation, err := strconv.Atoi(location)
		if err != nil {
		    log.Fatal(err)
		}
    ProcessInput(NumLocation, frequency)
  }
}

func ProcessInput(loc int, frequency map[int]int){
	frequency[loc]++
	if frequency[loc] == 3{
		fmt.Println("Current position is at", loc)

		// Clear Map
		for k := range frequency {
		    delete(frequency, k)
		}
	}
}
