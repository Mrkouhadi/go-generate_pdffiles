package main

// tut: https://www.youtube.com/watch?v=jwOy4JgleTU

import (
	"log"
	"os"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func Generate_pdf_file() {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20) // left,top,right   20 units, 10 units, 20 units
	BuildHeader(m)
	buildFruitList(m)
	err := m.OutputFileAndClose("./files/test.pdf")
	HandleError("could not save pdf file", err)

	log.Println("Pdf file saved successfully !")
}

// build a header of the pdf file
func BuildHeader(m pdf.Maroto) { // this header has 2 rows, and 1st row has an image, 2nd has 3 columns of txt
	m.RegisterHeader(func() {
		m.Row(50, func() {
			m.Col(12, func() {
				err := m.FileImage("./assets/images/img.png", props.Rect{
					Center:  true,
					Percent: 75,
				})
				HandleError("could not add image: ", err)
			})
		})
	})
	m.Row(10, func() { // 10 refers to units
		m.Col(12, func() { // 12 refers to number of spaces
			m.Text("Brought to you by Mr.kouhadi", props.Text{
				Color: getDarkPurpleColor(),
				Style: consts.Bold,
				Align: consts.Center,
				Top:   3,
				Size:  13,
			})
			m.Text("I am on the left side", props.Text{
				Color: getDarkPurpleColor(),
				Style: consts.Italic,
				Top:   3,
			})
			m.Text("I am on the right side", props.Text{
				Color: getDarkPurpleColor(),
				Style: consts.Italic,
				Align: consts.Right,
				Top:   3,
			})
		})
	})
}

// build content
func buildFruitList(m pdf.Maroto) {
	m.SetBackgroundColor(getTealColor())
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Products", props.Text{
				Top:    2,
				Size:   13,
				Color:  color.NewWhite(),
				Family: consts.Courier,
				Style:  consts.Bold,
				Align:  consts.Center,
			})
		})
	})
}

// ///////////// helpers
func getDarkPurpleColor() color.Color {
	return color.Color{
		Red:   88,
		Green: 80,
		Blue:  99,
	}
}
func getTealColor() color.Color {
	return color.Color{
		Red:   3,
		Green: 166,
		Blue:  166,
	}
}
func HandleError(msg string, err any) {
	if err != nil {
		log.Println(msg, err)
		os.Exit(1)
	}
}
