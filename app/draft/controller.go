package draft

import (
	"database/sql"
	"errors"

	"github.com/tecposter/tec-node-go/lib/ws"
)

const (
	cmdSave  = "draft.save"
	cmdFetch = "draft.fetch"
	cmdList  = "draft.list"
)

var (
	errCmdNotFound    = errors.New("Command not found in draft module")
	errRequirePostID  = errors.New("Require post id")
	errRequireContent = errors.New("Require content")
)

// Controller in draft
type Controller struct {
	serv *service
}

// NewCtrl returns Controller instance
func NewCtrl(db *sql.DB) *Controller {
	return &Controller{
		serv: newServ(db),
	}
}

// Handle handle ws response and request
func (ctrl *Controller) Handle(res ws.IResponse, req ws.IRequest) {
	switch req.CMD() {
	case cmdSave:
		ctrl.save(res, req)
	default:
		res.SetErr(errCmdNotFound)
	}
}

func (ctrl *Controller) save(res ws.IResponse, req ws.IRequest) {
	postIDBase58, ok := req.Param("postID")
	if !ok {
		res.SetErr(errRequirePostID)
		return
	}
	content, ok := req.Param("content")
	if !ok {
		res.SetErr(errRequireContent)
		return
	}

	err := ctrl.serv.save(postIDBase58.(string), content.(string))
	if err != nil {
		res.SetErr(err)
	}
}
