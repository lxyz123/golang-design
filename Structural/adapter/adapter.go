package main

import "fmt"

type NewPlayer interface {
	play(fileType, filename string)
}

type OldPlayer struct {
}

func (*OldPlayer) playMp3(filename string) {
	fmt.Println("play mp3: ", filename)
}

func (*OldPlayer) playMp4(filename string) {
	fmt.Println("play mp4: ", filename)
}

type PlayAdapter struct {
	oldPlayer OldPlayer
}

func (p PlayAdapter) play(fileType, filename string) {
	switch fileType {
	case "mp3":
		p.oldPlayer.playMp3(filename)
	case "mp4":
		p.oldPlayer.playMp4(filename)
	default:
		fmt.Println("unsupported file type")
	}
}

func main() {
	player := PlayAdapter{}
	player.play("mp3", "hundred miles")
	player.play("mp4", "beautiful in white")
	player.play("mp5", "monster")
}
