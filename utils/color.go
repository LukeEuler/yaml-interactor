package utils

import "fmt"

func raw(style, front int, content string) string {
	return fmt.Sprintf("%c[%d;%dm%s%c[0m ", 0x1B, style, front, content, 0x1B)
}

func normal(front int, content string) string {
	return raw(0, front, content)
}

func bold(front int, content string) string {
	return raw(1, front, content)
}

func dim(front int, content string) string {
	return raw(2, front, content)
}

func italic(front int, content string) string {
	return raw(3, front, content)
}

func underlined(front int, content string) string {
	return raw(4, front, content)
}

// Transparent XXX
func Transparent(content string) string {
	return bold(30, content)
}

// Red XXX
func Red(content string) string {
	return bold(31, content)
}

// Green XXX
func Green(content string) string {
	return bold(32, content)
}

// Yellow XXX
func Yellow(content string) string {
	return bold(33, content)
}

// Blue XXX
func Blue(content string) string {
	return bold(34, content)
}

// Magenta XXX
func Magenta(content string) string {
	return bold(35, content)
}

// Cyan XXX
func Cyan(content string) string {
	return bold(36, content)
}

// LightGray XXX
func LightGray(content string) string {
	return bold(37, content)
}
