package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data, err := store.Fetch(ctx)

		if err != nil {
			return
		}

		_, _ = fmt.Fprint(w, data)
	}
}
