package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type Saver interface {
	Save() error
}

// type Displayer interface {
// 	Display()
// }

type Outputable interface {
	Saver
	Display()
}

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Print(err)
		return
	}

	todoText := getUserInput("Todo list")
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Print(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		fmt.Print(err)
		return
	}

	err = outputData(userNote)

}

func outputData(data Outputable) error {
	data.Display()
	return saveData(data)

}

func saveData(data Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving  failed")
		return err
	}
	fmt.Println("Saving succeded")
	return nil
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
