package methods

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gobin/public"
	"gobin/public/templates"
	"html/template"
	"log"
	"net/http"
)

func (m Methods) Get(w http.ResponseWriter, r *http.Request) {
	key, ok := mux.Vars(r)["key"]

	if !ok || len(key) != 16 {
		templates.Error.Execute(w, "Wrong key format")

		return
	}

	// 1. Seek in memcached
	bin, err := m.db.Memcached.GetBin(key)
	if err != nil { // only internal
		templates.Error.Execute(w, public.Internal)
		log.Println(fmt.Errorf("get bin: memcached: get (%s): %w", key, err))

		return
	}

	// If found -> execute
	if bin != nil {
		// Execute template
		templates.View.Execute(w, view{
			&bin.Created,
			template.HTML(bin.Content),
		})

		return
	}

	var data []byte

	// 2. Seek in postgres
	data, err = m.db.Postgres.GetBin(key)
	if err != nil { // only internal
		templates.Error.Execute(w, public.Internal)
		log.Println(fmt.Errorf("get bin: postgres: get (%s): %w", key, err))

		return
	}

	// If not found -> error
	if data == nil {
		templates.Error.Execute(w, "bin not found")

		return
	}

	// Parse data
	err = json.Unmarshal(data, &bin)
	if err != nil {
		templates.Error.Execute(w, public.Internal)
		log.Println(fmt.Errorf("get bin: postgres: get (%s): unmarshalling: %w", key, err))

		return
	}

	// Save copy in memcached
	err = m.db.Memcached.NewBin(key, data)
	if err != nil {
		templates.Error.Execute(w, public.Internal)
		log.Println(fmt.Errorf("get bin: memcached: new bin (%s): %w", key, err))

		return
	}

	// Execute template
	templates.View.Execute(w, view{
		&bin.Created,
		template.HTML(bin.Content),
	})
}

type view struct {
	Created *string
	Content template.HTML
}
