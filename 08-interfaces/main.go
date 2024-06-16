package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/wahbifadhillah/learn-go/tree/main/08-interfaces/note"
	"github.com/wahbifadhillah/learn-go/tree/main/08-interfaces/todo"
)

type saver interface {
	Save() error
}

// Embedded interface
type outputtable interface {
	saver
	Display()
}

func main() {
	// Any type and value switch
	print(1)
	print(1.5)
	print("Heloo")

	// Main
	title, content := getNoteData()
	todoText := getTodoData()

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}

}

// ANY type
// Type Switch
func print(value any) {
	// Approach 2
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float:", floatVal)
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String:", stringVal)
	}

	// Approach 1
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println("String:", value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed!")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getTodoData() string {
	return getUserInput("Todo text:")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
