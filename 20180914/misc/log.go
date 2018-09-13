package misc

import "github.com/fatih/color"

// SuccessMessage is Println message
func SuccessMessage() {
	c := color.New(color.FgCyan).Add(color.Underline).Add(color.Bold)
	c.Println("SUCCESS")
}

// ErrorMessage is Println message
func ErrorMessage() {
	c := color.New(color.FgRed).Add(color.Underline).Add(color.Bold)
	c.Println("ERROR")
}
