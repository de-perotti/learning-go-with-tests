package tremetreme

func TremeTreme(intensidade int) (bum string) {
	for i := 0; i < intensidade; i += 1 {
		bum += "bum"
	}

	return bum
}
