package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var ghInstance GroupHandler

type GroupHandlerImpl struct {
	RequestHandler RequestHandler
	GroupService   service.GroupService
	AuthService    service.AuthService
}

func GetGroupHandlerInstance() GroupHandler {
	if ghInstance == nil {
		ghInstance = NewGroupHandler()
	}
	return ghInstance
}

func NewGroupHandler() GroupHandler {
	log.Logger.Info("New `GroupHandler` instance")
	log.Logger.Info("Inject `RequestHandler`, `AuthService`, `GroupService` to `GroupHandler`")
	return GroupHandlerImpl{
		RequestHandler: GetRequestHandlerInstance(),
		GroupService:   service.GetGroupServiceInstance(),
		AuthService:    service.GetAuthServiceInstance(),
	}
}

func (gh GroupHandlerImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		gh.Get(w, r)
	default:
		err := model.MethodNotAllowed()
		http.Error(w, err.ToJson(), err.Code)
	}
}

func (gh GroupHandlerImpl) Get(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get(entity.GroupName.String())

	groupEntities, err := gh.GroupService.Get(name)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(groupEntities)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (gh GroupHandlerImpl) Post(w http.ResponseWriter, r *http.Request) {
	var groupEntity *entity.Group

	body, err := gh.RequestHandler.InterceptHttp(w, r)
	if err != nil {
		return
	}

	json.Unmarshal(body, &groupEntity)
	if err := gh.RequestHandler.ValidateHttpRequest(w, groupEntity); err != nil {
		return
	}

	groupEntity, err = gh.GroupService.InsertGroup(groupEntity)
	if err != nil {
		http.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(groupEntity)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (gh GroupHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
}
