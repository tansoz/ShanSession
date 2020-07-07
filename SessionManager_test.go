package Session_test

import (
	"fmt"
	"net/http"
	"testing"
)
import "github.com/tansoz/ShanSession"

func TestNewSessionManager(t *testing.T) {

	handlers := &Handlers{SessionManager: Session.NewSessionManager("SESSID", 1440)}

	http.HandleFunc("/index.html", handlers.TestHandler)
	http.ListenAndServe("0.0.0.0:5656", nil)

}

type Handlers struct {
	SessionManager Session.SessionManager
}

func (this *Handlers) TestHandler(rw http.ResponseWriter, r *http.Request) {

	session := this.SessionManager.HttpSession(rw, r)

	count := session.Data().Get("count").Int() + 1
	rw.Write([]byte(fmt.Sprintf("<title>第%d次访问</title><h1>第%d次访问</h1>", count, count)))
	session.Data().Set("count", count)

}
