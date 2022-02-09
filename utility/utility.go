package utility

import (
	"bufio"
	"fmt"
	"go-vt100/terminal"
	"go-vt100/vt100"
	"os"
)

func DebugPrintf(y int, format string, content ...interface{}) {
	<-terminal.ControlSignal
	formatContent := fmt.Sprintf(format, content...)
	vt100.MoveCursorToAndPrint(2, y, formatContent)
	<-terminal.ControlSignal
	vt100.ClearLine()
}

// ReadFileLineOneByOne 逐行读取文件内容，执行函数返回 true 则继续读取，返回 false 则结束读取
func ReadFileLineOneByOne(filename string, f func(string) bool) error {
	file, openError := os.Open(filename)
	if openError != nil {
		return openError
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if !f(scanner.Text()) {
			break
		}
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	return nil
}
