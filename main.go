package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/Knetic/govaluate"
)

func main() {
	a := app.New()
	w := a.NewWindow("cursed calc")
	w.SetFixedSize(true)

	formula := ""

	formulaDisplay := widget.NewEntry()
	formulaDisplay.SetPlaceHolder("formula")
	formulaDisplay.OnChanged = func(a string) { formula = a }

	valueDisplay := widget.NewLabel("result")

	buttonPad := fyne.NewContainerWithLayout(layout.NewGridLayout(5))

	for i := 1; i < 10; i++ {
		buttonPad.AddObject(changeButton(fmt.Sprint(i), &formula, formulaDisplay))
	}
	buttonPad.AddObject(changeButton("0", &formula, formulaDisplay))

	bonus := []string{"+", "-", "/", "*", "%", "(", ")", "^", "<", ">"}

	for i := range bonus {
		buttonPad.AddObject(changeButton(bonus[i], &formula, formulaDisplay))
	}

	enter := widget.NewButton("enter", func() {
		//this is wrong way to do it but eh
		//someone else did the hard work
		expression, err := govaluate.NewEvaluableExpression(formula)
		if err != nil {
			valueDisplay.SetText("ree")
			return
		}

		value, _ := expression.Evaluate(nil)

		valueDisplay.SetText(fmt.Sprint(value))
	})

	clear := widget.NewButton("clear", func() {
		formula = ""
		formulaDisplay.SetText(formula)
		valueDisplay.SetText("")
	})

	submitPad := fyne.NewContainerWithLayout(layout.NewGridLayout(2))
	submitPad.AddObject(clear)
	submitPad.AddObject(enter)

	w.SetContent(widget.NewVBox(
		formulaDisplay,
		valueDisplay,
		buttonPad,
		submitPad,
	))

	w.ShowAndRun()
}

func changeButton(mod string, formula *string, formulaDisplay *widget.Entry) *widget.Button {
	button := widget.NewButton(mod, func() {
		*formula += mod
		formulaDisplay.SetText(*formula)

	})
	return button
}
