package src

import "fmt"

func red(content string) string {
	style := 1
	front := 31
	//back := 40
	return fmt.Sprintf("%c[%d;%dm%s%c[0m ", 0x1B, style, front, content, 0x1B)
}

func green(content string) string {
	style := 1
	front := 32
	//back := 40
	return fmt.Sprintf("%c[%d;%dm%s%c[0m ", 0x1B, style, front, content, 0x1B)
}

func blue(content string) string {
	style := 1
	front := 34
	//back := 40
	return fmt.Sprintf("%c[%d;%dm%s%c[0m ", 0x1B, style, front, content, 0x1B)
}
