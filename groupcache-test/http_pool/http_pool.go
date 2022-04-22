package http_pool

import (
	"fmt"
	"groupcache-test/local_utils"
	"groupcache-test/models"
	"net/http"
	"strings"
)

const (
	defaultBasePath = "/_chainmakeroracle_"
)

type HttpPool struct {
	Self     string
	BasePath string
	GroupMap *models.Group
}

func NewHttpPool(self, groupName string) *HttpPool {
	return &HttpPool{
		Self:     self,
		BasePath: defaultBasePath,
		GroupMap: &models.Group{
			Name: groupName,
		},
	}
}

func (thisP *HttpPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, thisP.BasePath) {
		panic("HttpPool serving unexpected path: " + r.URL.Path)
	}
	local_utils.Logf(fmt.Sprintf("%s  %s", r.Method, r.URL.Path))
	// /<basepath>/<groupname>/<key> required
	parts := strings.SplitN(r.URL.Path[len(thisP.BasePath):], "/", 2)
	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	groupName := parts[0]
	if groupName != local_utils.GCFG.Server.Id {
		http.Error(w, "no such group: "+groupName, http.StatusNotFound)
	}
	key := parts[1]
	thisP.GroupMap.Get(key)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write([]byte(key))
}
