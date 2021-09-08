// Package rolezao Braveza dos sete mares no aprendizado de go lek
package rolezao

import "strings"

// PrintRolezao Essa funçao aqui é braba
func PrintRolezao(onde, proverbialidade string) string {
	isEmpty := strings.TrimSpace(onde) == ""
	switch proverbialidade {
	case "openenglish":
		{
			if isEmpty {
				return "dudes do be doo"
			}

			return "dudes do be meeting at " + onde
		}
	default:
		{
			if isEmpty {
				return "so um rolezinho"
			}
			return "rolezao " + onde
		}
	}
}
