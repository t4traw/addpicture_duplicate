package main

import (
	_ "image/png"

	"github.com/xuri/excelize/v2"
)

func main() {
	xlsx, _ := excelize.OpenFile("./Book1.xlsx")

	templateSheetName := "Sheet2"
	templateSheetIndex, _ := xlsx.GetSheetIndex(templateSheetName)

	Copied1, _ := xlsx.NewSheet("Copied1")
	_ = xlsx.CopySheet(templateSheetIndex, Copied1)
	Copied2, _ := xlsx.NewSheet("Copied2")
	_ = xlsx.CopySheet(templateSheetIndex, Copied2)

	photoPotisions, _ := xlsx.SearchSheet(templateSheetName, "photo")

	xlsx.AddPicture("Copied1", photoPotisions[0], "./a.png", &excelize.GraphicOptions{})
	xlsx.AddPicture("Copied2", photoPotisions[0], "./b.png", &excelize.GraphicOptions{})

	xlsx.SaveAs("./output.xlsx")
}
