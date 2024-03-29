package middleware

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"gopkg.in/go-playground/validator.v9"

	"github.com/gorilla/mux"
	"github.com/tomoyane/grant-n-z/gnz/common"
	"github.com/tomoyane/grant-n-z/gnz/log"
	"github.com/tomoyane/grant-n-z/gnzserver/model"
)

// Http Header, Request scope const
const (
	Authorization             = "Authorization"
	ClientSecret              = "Client-Secret"
	ContentType               = "Content-Type"
	AccessControlAllowOrigin  = "Access-Control-Allow-Origin"
	AccessControlAllowHeaders = "Access-Control-Allow-Headers"
	ScopeSecret               = "secret"
	ScopeJwt                  = "jwt"
)

var iInstance Interceptor

type Interceptor interface {
	// Intercept Http request and Client-Secret header
	Intercept(next http.HandlerFunc) http.HandlerFunc

	// InterceptSecret
	// Intercept only http header
	InterceptSecret(next http.HandlerFunc) http.HandlerFunc

	// InterceptAuthenticateUser
	// Intercept Http request and Client-Secret header with user authentication
	InterceptAuthenticateUser(next http.HandlerFunc) http.HandlerFunc

	// InterceptAuthenticateGroupAdmin
	// Intercept Http request and Client-Secret header with user and group admin role authentication
	InterceptAuthenticateGroupAdmin(next http.HandlerFunc) http.HandlerFunc

	// InterceptAuthenticateGroupUser
	// Intercept Http request and Client-Secret header with user and group user role authentication
	InterceptAuthenticateGroupUser(next http.HandlerFunc) http.HandlerFunc

	// InterceptAuthenticateOperator
	// Intercept Http request and Client-Secret header with operator authentication
	InterceptAuthenticateOperator(next http.HandlerFunc) http.HandlerFunc
}

type InterceptorImpl struct {
	tokenProcessor TokenProcessor
}

func GetInterceptorInstance() Interceptor {
	if iInstance == nil {
		iInstance = NewInterceptor()
	}
	return iInstance
}

func NewInterceptor() Interceptor {
	return InterceptorImpl{
		tokenProcessor: GetTokenProcessorInstance(),
	}
}

func (i InterceptorImpl) Intercept(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Logger.Trace(rec)
				err := model.InternalServerError("Unexpected occurred")
				model.WriteError(w, err.ToJson(), err.Code)
			}
		}()

		if err := interceptHeader(w, r); err != nil {
			return
		}

		secret, secretErr := interceptClientSecret(r)
		if secretErr == nil {
			r = r.WithContext(context.WithValue(r.Context(), ScopeSecret, *secret))
		} else {
			r = r.WithContext(context.WithValue(r.Context(), ScopeSecret, nil))
		}

		next.ServeHTTP(w, r)
	}
}

func (i InterceptorImpl) InterceptSecret(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Logger.Trace(rec)
				err := model.InternalServerError("Unexpected occurred")
				model.WriteError(w, err.ToJson(), err.Code)
			}
		}()

		if err := interceptHeader(w, r); err != nil {
			return
		}

		secret, secretErr := interceptClientSecret(r)
		if secretErr != nil {
			err := model.Unauthorized(secretErr.Error())
			model.WriteError(w, err.ToJson(), err.Code)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), ScopeSecret, *secret))
		next.ServeHTTP(w, r)
	}
}

func (i InterceptorImpl) InterceptAuthenticateUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Logger.Trace(rec)
				err := model.InternalServerError("Unexpected occurred")
				model.WriteError(w, err.ToJson(), err.Code)
			}
		}()

		if err := interceptHeader(w, r); err != nil {
			return
		}

		secret, secretErr := interceptClientSecret(r)
		if secretErr == nil {
			r = r.WithContext(context.WithValue(r.Context(), ScopeSecret, *secret))
		} else {
			r = r.WithContext(context.WithValue(r.Context(), ScopeSecret, nil))
		}

		token := r.Header.Get(Authorization)
		jwtPayload, err := i.tokenProcessor.VerifyUserToken(token, "", "", "")
		if err != nil {
			model.WriteError(w, err.ToJson(), err.Code)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), ScopeJwt, *jwtPayload))
		next.ServeHTTP(w, r)
	}
}

func (i InterceptorImpl) InterceptAuthenticateGroupAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Logger.Trace(rec)
				err := model.InternalServerError("Unexpected occurred")
				model.WriteError(w, err.ToJson(), err.Code)
			}
		}()

		if err := interceptHeader(w, r); err != nil {
			return
		}

		secret, secretErr := interceptClientSecret(r)
		if secretErr != nil {
			err := model.BadRequest("Require Client-Secret.")
			model.WriteError(w, err.ToJson(), err.Code)
			return
		}

		token := r.Header.Get(Authorization)
		groupId := ParamGroupUuid(r)
		jwtPayload, err := i.tokenProcessor.VerifyUserToken(token, common.AdminRole, "", groupId)
		if err != nil {
			model.WriteError(w, err.ToJson(), err.Code)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), ScopeSecret, *secret))
		r = r.WithContext(context.WithValue(r.Context(), ScopeJwt, *jwtPayload))
		next.ServeHTTP(w, r)
	}
}

func (i InterceptorImpl) InterceptAuthenticateGroupUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Logger.Trace(rec)
				err := model.InternalServerError("Unexpected occurred")
				model.WriteError(w, err.ToJson(), err.Code)
			}
		}()

		if err := interceptHeader(w, r); err != nil {
			return
		}

		secret, secretErr := interceptClientSecret(r)
		if secretErr != nil {
			err := model.BadRequest("Require Client-Secret.")
			model.WriteError(w, err.ToJson(), err.Code)
			return
		}

		token := r.Header.Get(Authorization)
		groupId := ParamGroupUuid(r)
		jwtPayload, err := i.tokenProcessor.VerifyUserToken(token, fmt.Sprintf("%s,%s", common.AdminRole, common.UserRole), "", groupId)
		if err != nil {
			model.WriteError(w, err.ToJson(), err.Code)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), ScopeSecret, *secret))
		r = r.WithContext(context.WithValue(r.Context(), ScopeJwt, *jwtPayload))
		next.ServeHTTP(w, r)
	}
}

func (i InterceptorImpl) InterceptAuthenticateOperator(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Logger.Trace(rec)
				err := model.InternalServerError("Unexpected occurred")
				model.WriteError(w, err.ToJson(), err.Code)
			}
		}()

		if err := interceptHeader(w, r); err != nil {
			return
		}

		token := r.Header.Get(Authorization)
		jwtPayload, err := i.tokenProcessor.VerifyOperatorToken(token)
		if err != nil {
			model.WriteError(w, err.ToJson(), err.Code)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), ScopeJwt, *jwtPayload))
		next.ServeHTTP(w, r)
	}
}

// Intercept http request header
func interceptHeader(w http.ResponseWriter, r *http.Request) *model.ErrorResBody {
	w.Header().Set(ContentType, "application/json")
	w.Header().Set(AccessControlAllowOrigin, "*")
	w.Header().Set(AccessControlAllowHeaders, "*")
	if err := validateHeader(r); err != nil {
		model.WriteError(w, err.ToJson(), err.Code)
		return err
	}
	return nil
}

// Intercept Client-Secret header
func interceptClientSecret(r *http.Request) (*string, error) {
	clientSecret := r.Header.Get(ClientSecret)
	if strings.EqualFold(clientSecret, "") {
		return nil, errors.New("Required Client-Secret")
	}
	return &clientSecret, nil
}

// Validate http request header
func validateHeader(r *http.Request) *model.ErrorResBody {
	if r.Method == http.MethodOptions {
		return model.Options()
	}
	if r.Method != http.MethodGet && r.Header.Get(ContentType) != "application/json" {
		log.Logger.Info("Not allowed content-type")
		return model.BadRequest("Need to content type is only json.")
	}
	return nil
}

// Bind request body what http request converts to interface
func BindBody(w http.ResponseWriter, r *http.Request, i interface{}) *model.ErrorResBody {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if len(body) == 0 {
		err := model.BadRequest("Request is empty.")
		model.WriteError(w, err.ToJson(), err.Code)
		return err
	}

	if err := json.Unmarshal(body, i); err != nil {
		err := model.BadRequest("Request is not json.")
		model.WriteError(w, err.ToJson(), err.Code)
		return err
	}

	return nil
}

// Validate request body
func ValidateBody(w http.ResponseWriter, i interface{}) *model.ErrorResBody {
	if err := validator.New().Struct(i); err != nil {
		log.Logger.Info(err.Error())
		err := model.BadRequest("Invalid request.")
		model.WriteError(w, err.ToJson(), err.Code)
		return err
	}
	return nil
}

// Validate request body
func ValidateTokenRequest(w http.ResponseWriter, tokenRequest *model.TokenRequest) *model.ErrorResBody {
	if tokenRequest.GrantType == "" {
		tokenRequest.GrantType = "password"
	}
	switch tokenRequest.GrantType {
	case model.GrantPassword.String():
		if tokenRequest.Email == "" || tokenRequest.Password == "" || len(tokenRequest.Password) <= 7 {
			err := model.BadRequest("Invalid request.")
			model.WriteError(w, err.ToJson(), err.Code)
			return err
		}
	case model.GrantRefreshToken.String():
		if tokenRequest.RefreshToken == "" {
			err := model.BadRequest("Invalid request.")
			model.WriteError(w, err.ToJson(), err.Code)
			return err
		}
	default:
		err := model.BadRequest("Not support grant type.")
		model.WriteError(w, err.ToJson(), err.Code)
		return err
	}

	return nil
}

// Parse request group_uuid of path parameter
func ParamGroupUuid(r *http.Request) string {
	return mux.Vars(r)["group_uuid"]
}
