package main

import (
	"bufio"
	"fmt"
	"go/types"
	"os"
	"strings"

	"aayushjoshi2709/notes2/note"
	"aayushjoshi2709/notes2/todo"
)

type saver interface{
	Save() error
}

// type displayer interface{
// 	Display()
// }


type outputable interface{
	saver // embedded interface
	Display()
}
// this is same as

// type outputable interface{
// 	Save()
// 	Display()
// }

// or

// type outputable interface{
// 	saver
// 	displayer
// }



func main() {
	title, content := getNoteDate()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
	}
	err = outputData(userNote)
	if err != nil{
		return 
	}

	todoText := getTodoData()
	userTodo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
	}
	err = outputData(userTodo)
	if err != nil{
		return
	}

	printSomething(1)
	printSomething(1.5)
	printSomething("Hello")
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

func saveData(data saver) error{
	err := data.Save()
	if err != nil {
		fmt.Println("Error saving todo: ", err)
		return err
	}
	fmt.Println("Saving todo succeed...")
	return nil
}

func outputData(data outputable) error{
	data.Display()
	return saveData(data)
}


// any value is allowed
func printSomething(value interface{}){
	// this is only preffered inside a switch statement

	switch value.(type){
	case int:
		fmt.Println(value)
	case float64:
		fmt.Printf("%.2f\n", value)
	case string:
		fmt.Println(value)
	default:
		fmt.Println("Invalid value....")
	}

	// methods to access
	// typedValue, ok := value.(int)
	// if ok{
	// 	typedValue += 1
	// }
}