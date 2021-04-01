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
	players1["players1_key1"] = 3
	delete(players1, "players1_key") /* 删除映射中的key */
	/* 打印结果 map[players_key:1] map[players1_key:2 players1_key1:3] map[one:1 two:2] map[players3_key:3]*/
	fmt.Println(players, players1, players2, players3)
}
