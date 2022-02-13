package main

import (
	"github.com/kyokomi/emoji/v2"
)

func HelloMessage() string {
	helloMessage := emoji.Sprint("Hello :world_map:!")
	return helloMessage
}
