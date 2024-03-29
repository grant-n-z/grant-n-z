package users

import (
	"encoding/json"
	"net/http"

	"github.com/tomoyane/grant-n-z/gnz/entity"
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
	service      service.Service
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
	return GroupImpl{
		groupService: service.GetGroupServiceInstance(),
		service:      service.GetServiceInstance(),
	}
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
	secret := r.Context().Value(middleware.ScopeSecret)
	var groups []*entity.Group
	if secret == nil {
		data, err := gh.groupService.GetGroupByUser(jwt.UserUuid)
		if err != nil {
			model.WriteError(w, err.ToJson(), err.Code)
			return
		}
		groups = data
	} else {
		ser, err := gh.service.GetServiceBySecret(secret.(string))
		if err != nil {
			model.WriteError(w, err.ToJson(), err.Code)
			return
		}
		data, err := gh.groupService.GetGroupByServices(ser.Uuid.String())
		if err != nil {
			model.WriteError(w, err.ToJson(), err.Code)
			return
		}
		groups = data
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

	secret := r.Context().Value(middleware.ScopeSecret)
	if secret == nil {
		err := model.BadRequest("Require Client-Secret.")
		model.WriteError(w, err.ToJson(), err.Code)
	}

	jwt := r.Context().Value(middleware.ScopeJwt).(model.JwtPayload)
	group, err := gh.groupService.InsertGroupWithRelationalData(*groupEntity, jwt.UserUuid, secret.(string))
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
