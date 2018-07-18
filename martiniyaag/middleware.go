package martiniyaag

import (
	"net/http"

	"github.com/go-martini/martini"

	"github.com/sniperkit/yaag/middleware"
	"github.com/sniperkit/yaag/yaag"
	"github.com/sniperkit/yaag/yaag/models"
)

func Document(c martini.Context, w http.ResponseWriter, r *http.Request) {
	if !yaag.IsOn() {
		c.Next()
		return
	}
	apiCall := models.ApiCall{}
	writer := middleware.NewResponseRecorder(w)
	c.MapTo(writer, (*http.ResponseWriter)(nil))
	middleware.Before(&apiCall, r)
	c.Next()
	middleware.After(&apiCall, writer, r)
}
