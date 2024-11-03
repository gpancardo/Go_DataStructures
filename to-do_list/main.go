package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	name          string
	content       string
	current_state bool
}

type TaskList struct {
	T_List []Task
}

func (list *TaskList) add_task(new_task Task) {
	list.T_List = append(list.T_List, new_task)
}

func (list *TaskList) complete_task(index int) {
	list.T_List[index].current_state = true
}

func (list *TaskList) remove_task(index int) {
	list.T_List = append(list.T_List[:index], list.T_List[index+1:]...)
}

func (list *TaskList) edit_task(index int, t Task) {
	list.T_List[index] = t
}

func main() {
	list_instance := TaskList{}
	var status string

	fmt.Println("---TO-DO LIST--\n\n", "     -----\n", "    /    / ^\n", "   /    /  l\n", "  /----/   8")

	getChoice := bufio.NewReader(os.Stdin)
	for {
		var choice int
		fmt.Println("Pick an action:\n",
			"1. New task\n",
			"2. Mark task as completed\n",
			"3. Edit task\n",
			"4. Remove task\n",
			"5. Show all\n",
			"6. Leave the app")
		fmt.Print(">  ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var t Task
			fmt.Print("Task:")
			t.name, _ = getChoice.ReadString('\n')
			fmt.Print("Content:")
			t.content, _ = getChoice.ReadString('\n')
			list_instance.add_task(t)
			fmt.Println("Added succesfully!")
		case 2:
			var index int
			fmt.Print("Which task have you completed? ")
			fmt.Scanln(&index)
			list_instance.complete_task(index)
			fmt.Println("Congrats!")
		case 4:
			var index int
			fmt.Print("Which task do you want to remove? ")
			fmt.Scanln(&index)
			list_instance.remove_task(index)
			fmt.Println("Fine!")
		case 3:
			var index int
			var t Task
			var changeTitle string
			var changeContent string

			fmt.Print("Which task do you want to edit? ")
			fmt.Scanln(&index)

			fmt.Print("Do you want to edit the name? ")
			fmt.Scanln(&changeTitle)
			fmt.Print("Do you want to edit the content? ")
			fmt.Scanln(&changeContent)

			if changeTitle == "yes" {
				fmt.Print("New task name: ")
				t.name, _ = getChoice.ReadString('\n')
			}
			if changeContent == "yes" {
				fmt.Print("New content: ")
				t.content, _ = getChoice.ReadString('\n')
			}
			list_instance.edit_task(index, t)
			fmt.Println("Changes are up! Good luck")
		case 5:
			for i := 0; i < len(list_instance.T_List); i++ {
				if list_instance.T_List[i].current_state {
					status = "Completed"
				} else {
					status = "Not completed"
				}
				fmt.Println(list_instance.T_List[i].name, " - ", status)
			}
		case 6:
			return
		}
	}
}
