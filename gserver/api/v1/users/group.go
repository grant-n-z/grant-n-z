package users

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gserver/api"
	"github.com/tomoyane/grant-n-z/gserver/common/property"
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/log"
	"github.com/tomoyane/grant-n-z/gserver/model"
	"github.com/tomoyane/grant-n-z/gserver/service"
)

var ghInstance Group

type Group interface {
	// Implement group api
	Api(w http.ResponseWriter, r *http.Request)

	// Http GET method
	get(w http.ResponseWriter, r *http.Request)

	// Http POST method
	post(w http.ResponseWriter, r *http.Request, body []byte)

	// Http DELETE method
	delete(w http.ResponseWriter, r *http.Request)
}

// Group api struct
type GroupImpl struct {
	request      api.Request
	groupService service.GroupService
}

// Get Policy instance.
// If use singleton pattern, call this instance method
func GetGroupInstance() Group {
	if ghInstance == nil {
		ghInstance = NewGroup()
	}
	return ghInstance
}

// Constructor
func NewGroup() Group {
	log.Logger.Info("New `Group` instance")
	log.Logger.Info("Inject `request`, `AuthService`, `GroupService` to `Group`")
	return GroupImpl{
		request:      api.GetRequestInstance(),
		groupService: service.GetGroupServiceInstance(),
	}
}

func (gh GroupImpl) Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := gh.request.Intercept(w, r, property.AuthUser)
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
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (gh GroupImpl) get(w http.ResponseWriter, r *http.Request) {
	groups, err := gh.groupService.GetGroupOfUser()
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(groups)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (gh GroupImpl) post(w http.ResponseWriter, r *http.Request, body []byte) {
	var request *entity.Group
	json.Unmarshal(body, &request)
	if err := gh.request.ValidateBody(w, request); err != nil {
		return
	}

	group, err := gh.groupService.InsertGroup(request)
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(group)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (gh GroupImpl) delete(w http.ResponseWriter, r *http.Request) {
}
