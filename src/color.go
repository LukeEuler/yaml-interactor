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

func yellow(content string) string {
	style := 1
	front := 33
	//back := 40
	return fmt.Sprintf("%c[%d;%dm%s%c[0m ", 0x1B, style, front, content, 0x1B)
}

func blue(content string) string {
	style := 1
	front := 34
	//back := 40
	return fmt.Sprintf("%c[%d;%dm%s%c[0m ", 0x1B, style, front, content, 0x1B)
}

func magenta(content string) string {
	style := 1
	front := 35
	//back := 40
	return fmt.Sprintf("%c[%d;%dm%s%c[0m ", 0x1B, style, front, content, 0x1B)
}

func cyan(content string) string {
	style := 1
	front := 36
	//back := 40
	return fmt.Sprintf("%c[%d;%dm%s%c[0m ", 0x1B, style, front, content, 0x1B)
}

func lightGray(content string) string {
	style := 1
	front := 37
	//back := 40
	return fmt.Sprintf("%c[%d;%dm%s%c[0m ", 0x1B, style, front, content, 0x1B)
}
