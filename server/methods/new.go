package methods

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gobin/database/model"
	"gobin/public"
	"gobin/public/templates"
	"log"
	"net/http"
	"time"
)

func (m Methods) New(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	content := r.Form.Get("content")
	if len(content) < 30 {
		templates.Error.Execute(w, "Your text is too small!")

		return
	}

	buf := make([]byte, 8)
	rand.Read(buf)
	key := hex.EncodeToString(buf)

	bin := model.Bin{
		Created: time.Now().Format(time.RFC822),
		Content: content,
	}

	data, err := json.Marshal(bin)
	if err != nil {
		templates.Error.Execute(w, public.Internal)
		log.Println(fmt.Errorf("new bin: marshal: %w", err))

		return
	}

	// 1. Write to memcached
	err = m.db.Memcached.NewBin(key, data)
	if err != nil {
		templates.Error.Execute(w, public.Internal)
		log.Println(fmt.Errorf("new bin: memcached: %w", err))

		return
	}

	// 2. Write to postgres
	err = m.db.Postgres.NewBin(key, data)
	if err != nil {
		log.Println(fmt.Errorf("new bin: postgres: %w", err))
	}

	http.Redirect(w, r, "/"+key, http.StatusPermanentRedirect)
}
