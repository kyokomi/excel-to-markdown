package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tealeg/xlsx/v3"
	"github.com/urfave/cli/v3"
	"golang.org/x/sync/errgroup"
)

func readWriteSheet(inputFilePath, outputDirPath string) error {
	xlFile, err := xlsx.OpenFile(inputFilePath)
	if err != nil {
		return err
	}

	for _, sheet := range xlFile.Sheets {
		// sheet単位でfile生成
		fmt.Printf("Start %s ...\n", sheet.Name)

		writeFilePath := filepath.Join(outputDirPath, sheet.Name+".md")
		if err := writeSheetMarkdown(sheet, writeFilePath); err != nil {
			return err
		}
		sheet.Close()

		fmt.Printf("End %s => %s\n", sheet.Name, writeFilePath)
	}

	return nil
}

func writeSheetMarkdown(sheet *xlsx.Sheet, writeFilePath string) error {
	f, err := os.Create(writeFilePath)
	if err != nil {
		return err
	}

	// bufio.Writerのエラーは内部に保持されFlushで回収できるため、
	// 各WriteStringの戻り値は個別にチェックしない
	w := bufio.NewWriter(f)
	if err := writeSheetRows(sheet, w); err != nil {
		f.Close()
		return err
	}
	if err := w.Flush(); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

func writeSheetRows(sheet *xlsx.Sheet, w *bufio.Writer) error {
	hyou := false
	rowIdx := 0
	// rowはまとめて1行にする
	return sheet.ForEachRow(func(row *xlsx.Row) error {
		defer func() { rowIdx++ }()

		var cells []string
		if err := row.ForEachCell(func(cell *xlsx.Cell) error {
			cells = append(cells, cell.Value)
			return nil
		}); err != nil {
			return err
		}
		// xlsx v3はシートの最大列数まで空セルを補完するため、
		// 旧版(v1)の「行に存在するセルまで」の挙動に合わせて末尾の空セルを除去する
		for len(cells) > 0 && cells[len(cells)-1] == "" {
			cells = cells[:len(cells)-1]
		}

		if rowIdx == 0 {
			// #見出し
			w.WriteString("# ")
		}

		text := strings.Join(cells, "")

		if len(text) == 0 {
			hyou = false
			w.WriteString("\n")
			w.WriteString("## ")
			return nil
		}

		if len(cells) >= 2 && len(cells[0]) == 0 {
			w.WriteString("- ")
			w.WriteString(cells[1])

		} else if len(cells) >= 2 {

			// 表
			for _, cell := range cells {
				w.WriteString("|")
				w.WriteString(cell)
			}
			w.WriteString("|")

			if !hyou {
				w.WriteString("\n")
				w.WriteString(strings.Repeat("| --- ", len(cells)))
				w.WriteString("|")
				hyou = true
			}

		} else if strings.HasPrefix(cells[0], "http") {
			w.WriteString(fmt.Sprintf("![%s](%s)", cells[0], cells[0]))
			w.WriteString("\n")
		} else {
			// その他
			w.WriteString(text)
			w.WriteString("\n")
		}
		w.WriteString("\n")
		return nil
	})
}

func main() {
	cmd := &cli.Command{
		Name:    "excel-to-markdown",
		Version: Version,
		Usage:   "convert Excel (.xlsx) files to GitHub-Flavored Markdown",
		Authors: []any{"kyokomi <kyoko1220adword@gmail.com>"},
		Action:  doMain,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "input-dir",
				Aliases: []string{"i"},
				Usage:   "convert target directory path",
			},
			&cli.StringFlag{
				Name:    "output-dir",
				Aliases: []string{"o"},
				Usage:   "dist directory after convert path",
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func doMain(ctx context.Context, cmd *cli.Command) error {
	inputDirPath := cmd.String("input-dir")
	outputDirPath := cmd.String("output-dir")

	if inputDirPath == "" || outputDirPath == "" {
		return cli.ShowRootCommandHelp(cmd)
	}

	entries, err := os.ReadDir(inputDirPath)
	if err != nil {
		return err
	}

	g, _ := errgroup.WithContext(ctx)
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if !strings.HasSuffix(entry.Name(), ".xlsx") {
			fmt.Printf("skip %s: not a .xlsx file\n", entry.Name())
			continue
		}

		inputFilePath := filepath.Join(inputDirPath, entry.Name())
		fileOutputDirPath := filepath.Join(outputDirPath, strings.TrimSuffix(entry.Name(), ".xlsx"))
		if err := os.MkdirAll(fileOutputDirPath, 0755); err != nil {
			return err
		}

		g.Go(func() error {
			return readWriteSheet(inputFilePath, fileOutputDirPath)
		})
	}
	return g.Wait()
}
