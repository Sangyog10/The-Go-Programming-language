package main

import (
	"bufio"
	"example.com/note/note"
	"fmt"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Print(err)
		return
	}

	userNote.DisplayNote()
	err = userNote.SaveNote()
	if err != nil {
		fmt.Println("Saving note failed")
	}
	fmt.Println("Saving the note succeded")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title")
	content := getUserInput("Note content")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	// var value string
	// fmt.Scanln(&value) //it wont work when we want muliple input at same time
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n') //means stop reading at line break
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
