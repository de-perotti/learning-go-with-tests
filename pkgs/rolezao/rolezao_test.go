package rolezao

import "testing"

func TestPrintRolezao(t *testing.T) {
	for _, tc := range []struct {
		onde string
		want string
	}{
		{" ", "so um rolezinho"},
		{"\t", "so um rolezinho"},
		{"\r", "rolezao \r"},
	} {
		t.Run("da para passar string com whitespace e dar um rolezinho", func(t *testing.T) {
			rolezao := PrintRolezao(tc.onde)

			if rolezao != tc.want {
				t.Errorf("ma deu um ruim da porra. eu queria %q, mas ganhei %q", tc.want, rolezao)
			}
		})
	}

	t.Run("da para passar string vazia e dar um rolezinho", func(t *testing.T) {
		rolezao := PrintRolezao("")
		want := "so um rolezinho"

		if rolezao != want {
			t.Errorf("ma deu um ruim da porra. eu queria %q, mas ganhei %q", want, rolezao)
		}
	})

	t.Run("da para passar string nao vazia e dar um rolezao", func(t *testing.T) {
		rolezao := PrintRolezao("no shop iguatemi")
		want := "rolezao no shop iguatemi"

		if rolezao != want {
			t.Errorf("ma deu um ruim da porra. eu queria %q, mas ganhei %q", want, rolezao)
		}
	})

}
