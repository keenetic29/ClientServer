package console

import (
    "bufio"
    "fmt"
    "os"
    "client/internal/services"
)

type Console struct {
    analyzer *services.FileAnalyzer
    scanner  *bufio.Scanner
}

func New(a *services.FileAnalyzer) *Console {
    return &Console{
        analyzer: a,
        scanner:  bufio.NewScanner(os.Stdin),
    }
}

func (c *Console) Run() {
    for {
        fmt.Print("\nMenu:\n1. Set file\n2. Analyze\n3. Exit\nChoose: ")
        c.scanner.Scan()

        switch c.scanner.Text() {
        case "1":
            c.setFile()
        case "2":
            c.analyze()
        case "3":
            return
        default:
            fmt.Println("Invalid choice")
        }
    }
}

func (c *Console) setFile() {
    fmt.Print("Enter file path: ")
    c.scanner.Scan()
    if err := c.analyzer.SetPath(c.scanner.Text()); err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Println("File set successfully")
    }
}

func (c *Console) analyze() {
    result, err := c.analyzer.Analyze()
    if err != nil {
        fmt.Printf("Analysis error: %v\n", err)
        return
    }

    fmt.Printf("\nResults:\nFile: %s\nLines: %d\nWords: %d\nChars: %d\n",
        result.FileName, result.Lines, result.Words, result.Chars)
}