package v1

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/api"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var ghInstance Group

type Group interface {
	Api(w http.ResponseWriter, r *http.Request)

	get(w http.ResponseWriter, r *http.Request)

	post(w http.ResponseWriter, r *http.Request, body []byte)

	delete(w http.ResponseWriter, r *http.Request)
}

type GroupImpl struct {
	Request      api.Request
	GroupService service.GroupService
}

func GetGroupInstance() Group {
	if ghInstance == nil {
		ghInstance = NewGroup()
	}
	return ghInstance
}

func NewGroup() Group {
	log.Logger.Info("New `Group` instance")
	log.Logger.Info("Inject `Request`, `AuthService`, `GroupService` to `Group`")
	return GroupImpl{
		Request:      api.GetRequestInstance(),
		GroupService: service.GetGroupServiceInstance(),
	}
}

func (gh GroupImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, _, err := gh.Request.Intercept(w, r, "")
	if err != nil {
		return
	}

	switch r.Method {
	case http.MethodGet:
		gh.get(w, r)
	case http.MethodPost:
		gh.post(w, r, body)
	default:
		err := model.MethodNotAllowed()
		model.Error(w, err.ToJson(), err.Code)
	}
}

func (gh GroupImpl) get(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get(entity.GroupName.String())

	groupEntities, err := gh.GroupService.Get(name)
	if err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(groupEntities)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (gh GroupImpl) post(w http.ResponseWriter, r *http.Request, body []byte) {
	var groupEntity *entity.Group

	json.Unmarshal(body, &groupEntity)
	if err := gh.Request.ValidateBody(w, groupEntity); err != nil {
		return
	}

	group, err := gh.GroupService.InsertGroup(groupEntity)
	if err != nil {
		model.Error(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(group)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (gh GroupImpl) delete(w http.ResponseWriter, r *http.Request) {
}
