package main

import (
  "dates"
  "fmt"
)

func main() {
  days := 3
  fmt.Println("Your appointment is is", days, "days")
  fmt.Println("with a follow-up in", days+dates.DaysInWeek, "days")
}