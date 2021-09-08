package rolezao

import "testing"

func TestPrintRolezao(t *testing.T) {
	t.Run("role fala openenglish tbm", func(t *testing.T) {
		onde := "a la playa"
		lingua := "openenglish"
		rolezao := PrintRolezao(onde, lingua)
		want := "dudes do be meeting at a la playa"

		if rolezao != want {
			t.Errorf("ma deu um ruim da porra. eu queria %q, mas ganhei %q", want, rolezao)
		}
	})

	t.Run("rolezinho no openenglish", func(t *testing.T) {
		onde := ""
		lingua := "openenglish"
		rolezao := PrintRolezao(onde, lingua)
		want := "dudes do be doo"

		if rolezao != want {
			t.Errorf("ma deu um ruim da porra. eu queria %q, mas ganhei %q", want, rolezao)
		}
	})

	for _, tc := range []struct {
		onde string
		want string
	}{
		{" ", "so um rolezinho"},
		{"\t", "so um rolezinho"},
		{"\r", "rolezao \r"},
	} {
		t.Run("da para passar string com whitespace e dar um rolezinho", func(t *testing.T) {
			rolezao := PrintRolezao(tc.onde, "aramaico")

			if rolezao != tc.want {
				t.Errorf("ma deu um ruim da porra. eu queria %q, mas ganhei %q", tc.want, rolezao)
			}
		})
	}

	t.Run("da para passar string vazia e dar um rolezinho", func(t *testing.T) {
		rolezao := PrintRolezao("", "aramaico")
		want := "so um rolezinho"

		if rolezao != want {
			t.Errorf("ma deu um ruim da porra. eu queria %q, mas ganhei %q", want, rolezao)
		}
	})

	t.Run("da para passar string nao vazia e dar um rolezao", func(t *testing.T) {
		rolezao := PrintRolezao("no shop iguatemi", "aramaico")
		want := "rolezao no shop iguatemi"

		if rolezao != want {
			t.Errorf("ma deu um ruim da porra. eu queria %q, mas ganhei %q", want, rolezao)
		}
	})
}
