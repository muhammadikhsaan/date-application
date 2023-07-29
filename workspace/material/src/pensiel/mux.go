package pensiel

import "github.com/go-chi/chi/v5"

type Mux struct {
	*chi.Mux
}

func NewRouter() *Mux {
	c := chi.NewRouter()

	return &Mux{c}
}

func (m *Mux) Route(pattern string, fn func(r Router)) {
	m.Mux.Route(pattern, func(r chi.Router) {
		fn(&router{r})
	})
}
