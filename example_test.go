package dlgs

import (
	"time"
)

func ExampleColor() {
	_, _, err := Color("Pick color", "#BEBEBE")
	if err != nil {
		panic(err)
	}
}

func ExampleDate() {
	_, _, err := Date("Date", "Date of traveling:", time.Now())
	if err != nil {
		panic(err)
	}
}

func ExampleEntry() {
	_, _, err := Entry("Entry", "Enter something here, anything:", "default text")
	if err != nil {
		panic(err)
	}
}

func ExampleFile() {
	_, _, err := File("Select file", "", false)
	if err != nil {
		panic(err)
	}
}

func ExampleFileMulti() {
	_, _, err := FileMulti("Select files", "")
	if err != nil {
		panic(err)
	}
}

func ExampleList() {
	_, _, err := List("List", "Select item from list:", []string{"Bug", "New Feature", "Improvement"})
	if err != nil {
		panic(err)
	}
}

func ExampleListMulti() {
	_, _, err := ListMulti("ListMulti", "Select languages from list:", []string{"PHP", "Go", "Python", "Bash"})
	if err != nil {
		panic(err)
	}
}

func ExampleInfo() {
	_, err := Info("Info", "Lorem ipsum dolor sit amet.")
	if err != nil {
		panic(err)
	}
}

func ExampleWarning() {
	_, err := Warning("Warning", "Incomplete information!")
	if err != nil {
		panic(err)
	}
}

func ExampleError() {
	_, err := Error("Error", "Cannot divide by zero.")
	if err != nil {
		panic(err)
	}
}

func ExampleQuestion() {
	_, err := Question("Question", "Are you sure you want to format this media?", true)
	if err != nil {
		panic(err)
	}
}

func ExamplePassword() {
	_, _, err := Password("Password", "Enter your API key:")
	if err != nil {
		panic(err)
	}
}
