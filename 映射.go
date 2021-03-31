package main

import "fmt"

var players = make(map[string]int, 10)
var players1 = map[string]int{}

func main() {
	players2 := map[string]int{"one": 1, "two": 2}
	players3 := map[string]int{}
	players3["players3_key"] = 3
	players["players_key"] = 1
	players1["players1_key"] = 2
	fmt.Println(players, players1, players2, players3)
}
