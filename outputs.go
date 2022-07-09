package main

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

func token(token string) {
	c := color.New(color.FgYellow)
	c.Print("! ")
	c1 := color.New(color.FgWhite)
	c1.Print("chat token ")
	c2 := color.New(color.Faint).Add(color.Bold)
	c2.Print(token + "\n")
}

func bulletPoint(text string) {
	c := color.New(color.FgCyan)
	c.Print("- ")
	c1 := color.New(color.FgWhite)
	c1.Print(text + "\n")
}

func success(text string) {
	c := color.New(color.FgGreen)
	c.Print("🗸")
	c1 := color.New(color.FgWhite)
	c1.Print(text + "\n")
}

func otherChat(other string, text string) {
	c := color.New(color.Faint)
	c.Print("["+other+"] " + text + "\n")
}

func title(title string) {
	d := color.New(color.FgWhite, color.Bold)
	d.Println(title)
}

func askInput(question string) string {
	result := ""
	prompt := &survey.Input{
		Message: question,
	}
	survey.AskOne(prompt, &result)
	return result
}

func askOptions(question string, options []string) string {
	result := ""
	prompt := &survey.Select{
		Message: question,
		Options: options,
	}
	survey.AskOne(prompt, &result)
	return result
}