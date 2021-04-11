package main

import (
  "fmt"
  "keyboard"
  "log"
)

func main() {
  fmt.Print("Enter a temperature in Fahrenheit:")
  fahrenheit, err := keyboard.GetFloat()
  if err != nil {
    log.Fatal(err)
  }
  celsius := (fahrenheit - 32) * 5 / 9
  fmt.Println("%0.2f degress Celsius \n", celsius)
}
