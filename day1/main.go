package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func findMax(data map[int]int) (elf, calories int){ 
    for key, value := range data{
        if value> calories{
            elf = key + 1 // plus 1 since you don't say elf 0 
            calories = value 
        }
    }
    return elf,calories 
}



func main(){
    f, err := os.Open("data.txt")
    if err!= nil {
        log.Fatal(err)
    }
    defer f.Close()
    // initiating a new scanner 
    scanner := bufio.NewScanner(f)
    //initilising count varibale
    count := 0 
    calorieCount := make(map[int]int)
    for scanner.Scan(){
        strValue := scanner.Text()
        if strValue == ""{
            count++ 
            continue 
        }
        value, err := strconv.Atoi(strValue)
        if err!= nil {
            log.Fatal(err)
        }
        calorieCount[count] += value
    }
    if err := scanner.Err(); err!= nil {
        log.Fatal(err)
    }
    
    elf, calories := findMax(calorieCount)
    fmt.Println("Elf number: ",elf, ", has the highest calories: ",calories)
}

