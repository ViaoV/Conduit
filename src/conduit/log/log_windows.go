// +build windows

package log

import (
	"fmt"
	"strings"
)

func Status(label, value string, success bool) {
	var successStr string
	if success {
		successStr = "+"
	} else {
		successStr = "-"
	}
	str := fmt.Sprintf("%s  [%s] - %s", successStr, value, label)
	write("", str, noStyle)
}

func write(tag, text string, style func(string) string) {
	if tag == "ALERT" {
		fmt.Println(fmt.Sprintf("[%s]", text))
	} else if tag == "ERROR" {
		fmt.Println(fmt.Sprintf("!! %s", text))
	} else if tag == "DEBUG" {
		fmt.Println(fmt.Sprintf("-- %s", text))
	} else if tag == "FATAL" {
		fmt.Println(fmt.Sprintf("(!!) %s", text))
	} else if tag == "WARN" {
		fmt.Println(fmt.Sprintf("! %s", text))
	} else {
		fmt.Println(fmt.Sprintf("%s", text))
	}
	if LogFile == true {
		writeFile(tag, text)
	}
}

func Stats(name string, value interface{}) {

	nameStr := fmt.Sprintf("%s:", name)
	padding := strings.Repeat(" ", 20-len(nameStr))
	var valueStr = ""
	switch value.(type) {
	case int64:
		valueStr = fmt.Sprintf("%d", value)
	case string:
		valueStr = value.(string)
	}
	str := fmt.Sprintf("   %s %s %s", nameStr, padding, valueStr)
	write("", str, noStyle)
}