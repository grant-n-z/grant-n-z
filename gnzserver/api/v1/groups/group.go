package groups

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var ghInstance Group

type Group interface {
	// Http GET method
	Get(w http.ResponseWriter, r *http.Request)
}

type GroupImpl struct {
	GroupService service.GroupService
}

func GetGroupInstance() Group {
	if ghInstance == nil {
		ghInstance = NewGroup()
	}
	return ghInstance
}

func NewGroup() Group {
	log.Logger.Info("New `v1.groups.Group` instance")
	return GroupImpl{
		GroupService: service.GetGroupServiceInstance(),
	}
}

func (gh GroupImpl) Get(w http.ResponseWriter, r *http.Request) {
	id, err := middleware.ParamGroupId(r)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	group, err := gh.GroupService.GetGroupById(id)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(group)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
