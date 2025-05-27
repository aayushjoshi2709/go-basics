package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"aayushjoshi2709/notes/note"
	"aayushjoshi2709/notes/todo"
)


func main() {
	title, content := getNoteDate()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
	}

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Error saving the note: ", err)
	}
	fmt.Println("Saving the note succeeded")

	todoText := getTodoData()
	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
	}
	userTodo.Display()
	err = userTodo.Save()
	if err != nil {
		fmt.Println("Error saving todo: ", err)
	}
	fmt.Println("Saving todo succeed...")
}

func getTodoData() string {
	content := getUserInput("Enter todo text: ")
	return content
}

func getNoteDate() (string, string) {
	title := getUserInput("enter title for the note: ")
	content := getUserInput("enter note content for the note: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
