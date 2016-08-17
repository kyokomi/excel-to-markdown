package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/tealeg/xlsx"
	"github.com/urfave/cli"
)

func readWriteSheet(inputFilePath, outputDirPath string) error {
	xlFile, err := xlsx.OpenFile(inputFilePath)
	if err != nil {
		return err
	}

	for _, sheet := range xlFile.Sheets {
		// sheet単位でfile生成
		fmt.Printf("Start %s ...\n", sheet.Name)

		writeFilePath := strings.Join([]string{outputDirPath, sheet.Name}, "/") + ".md"
		f, err := os.Create(writeFilePath)
		if err != nil {
			log.Fatal(err)
		}

		hyou := false
		// rowはまとめて1行にする
		for rowIdx, row := range sheet.Rows {

			if rowIdx == 0 {
				// #見出し
				f.WriteString("# ")
			}

			text := ""
			for _, cell := range row.Cells {
				text += cell.Value
			}

			if len(text) == 0 {
				hyou = false
				f.WriteString("\n")
				f.WriteString("## ")
				continue
			}

			if len(row.Cells) >= 2 && len(row.Cells[0].Value) == 0 {
				f.WriteString("- ")
				idx := 1
				f.WriteString(row.Cells[idx].Value)

			} else if len(row.Cells) >= 2 {

				// 表
				for _, cell := range row.Cells {
					f.WriteString("|")
					f.WriteString(cell.Value)
				}
				f.WriteString("|")

				if !hyou {
					f.WriteString("\n")
					f.WriteString(strings.Repeat("| --- ", len(row.Cells)))
					f.WriteString("|")
					hyou = true
				}

			} else if strings.HasPrefix(row.Cells[0].Value, "http") {
				f.WriteString(fmt.Sprintf("![%s](%s)", row.Cells[0].Value, row.Cells[0].Value))
				f.WriteString("\n")
			} else {
				// その他
				f.WriteString(text)
				f.WriteString("\n")
			}
			f.WriteString("\n")
		}
		fmt.Printf("End %s => %s\n", sheet.Name, writeFilePath)
		f.Close()
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "excel-to-markdown"
	app.Version = Version
	app.Usage = ""
	app.Author = "kyokomi"
	app.Email = "kyoko1220adword@gmail.com"
	app.Action = doMain
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input-dir,i",
			Value: "",
			Usage: "convert target directory path",
		},
		cli.StringFlag{
			Name:  "output-dir,o",
			Value: "",
			Usage: "dist directory after convert path",
		},
	}
	app.Run(os.Args)
}

func doMain(c *cli.Context) error {
	inputDirPath := c.String("input-dir")
	outputDirPath := c.String("output-dir")

	if inputDirPath == "" || outputDirPath == "" {
		cli.ShowAppHelp(c)
		return nil
	}

	d, err := ioutil.ReadDir(inputDirPath)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, file := range d {
		if file.IsDir() {
			continue
		}
		inputFilePath := strings.Join([]string{inputDirPath, file.Name()}, "/")
		if !strings.HasSuffix(inputFilePath, ".xlsx") {
			fmt.Println("error don't xlsx file.")
			continue
		}

		outputDirPath := strings.Join([]string{outputDirPath, file.Name()}, "/")
		outputDirPath = strings.TrimSuffix(outputDirPath, ".xlsx")
		if _, err := ioutil.ReadDir(outputDirPath); err != nil {
			err := os.Mkdir(outputDirPath, 0755)
			if err != nil {
				return err
			}
		}

		wg.Add(1)
		go func(inputFilePath string) {
			readWriteSheet(inputFilePath, outputDirPath)
			wg.Done()
		}(inputFilePath)
	}
	wg.Wait()

	return nil
}
