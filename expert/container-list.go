package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Raisa")
	data.PushBack("Supriatna")
	data.PushFront("Gozenx")
	data.PushBack("Zenklot")

	for i := data.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
