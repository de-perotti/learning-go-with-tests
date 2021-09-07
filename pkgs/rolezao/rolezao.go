// Package rolezao Braveza dos sete mares no aprendizado de go lek
package rolezao

import "strings"

// PrintRolezao Essa funçao aqui é braba
func PrintRolezao(onde string) string {
	if strings.Trim(onde, "\t ") == "" {
		return "so um rolezinho"
	}

	return "rolezao " + onde
}
