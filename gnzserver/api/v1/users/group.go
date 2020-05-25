package users

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/middleware"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
	"github.com/tomoyane/grant-n-z/gnzserver/service"
)

var ghInstance Group

type Group interface {
	// Implement group api
	// Endpoint is `/api/v1/users/group`
	Api(w http.ResponseWriter, r *http.Request)

	// Http GET method
	get(w http.ResponseWriter, r *http.Request)

	// Http POST method
	post(w http.ResponseWriter, r *http.Request)

	// Http DELETE method
	delete(w http.ResponseWriter, r *http.Request)
}

// Group api struct
type GroupImpl struct {
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
	log.Logger.Info("New `v1.users.Group` instance")
	return GroupImpl{groupService: service.GetGroupServiceInstance()}
}

func (gh GroupImpl) Api(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		gh.get(w, r)
	case http.MethodPost:
		gh.post(w, r)
	default:
		err := model.MethodNotAllowed()
		model.WriteError(w, err.ToJson(), err.Code)
	}
}

func (gh GroupImpl) get(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value(middleware.ScopeJwt).(model.JwtPayload)
	groups, err := gh.groupService.GetGroupByUser(jwt.UserUuid.String())
	if err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return
	}

	res, _ := json.Marshal(groups)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (gh GroupImpl) post(w http.ResponseWriter, r *http.Request) {
	var groupEntity *entity.Group
	if err := middleware.BindBody(w, r, &groupEntity); err != nil {
		return
	}

	if err := middleware.ValidateBody(w, groupEntity); err != nil {
		return
	}

	jwt := r.Context().Value(middleware.ScopeJwt).(model.JwtPayload)
	group, err := gh.groupService.InsertGroupWithRelationalData(*groupEntity, jwt.UserUuid.String(), r.Context().Value(middleware.ScopeSecret).(string))
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
