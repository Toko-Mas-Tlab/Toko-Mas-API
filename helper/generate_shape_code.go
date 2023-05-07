package helper

import "strings"

func GenerateShapeCode(name string) string {
	code := ""

	code += string(name[0])
	code += string(name[1])
	code += string(name[2])

	inisial := strings.ToUpper(code)

	return inisial
}
