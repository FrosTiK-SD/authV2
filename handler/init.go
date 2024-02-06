package handler

import (
	"github.com/FrosTiK-SD/auth/controller"
	"github.com/FrosTiK-SD/auth/model"
	mongik "github.com/FrosTiK-SD/mongik/models"
	jsoniter "github.com/json-iterator/go"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type Mode string

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	HANDLER    Mode = "HANDLER"
	MIDDLEWARE Mode = "MIDDLEWARE"
)

type Session struct {
	Error   error
	Student *model.StudentPopulated
}

type Handler struct {
	MongikClient *mongik.Mongik
	JwkSet       *jwk.Set
	Session      *Session
	Config       Config
}

type Config struct {
	Mode Mode
}

type RoleCheckerHandler struct {
	Role string
}

func NewAuthClient(mongik *mongik.Mongik) *Handler {
	defaultJwkSet, _ := controller.GetJWKs(mongik.CacheClient, false)
	return &Handler{
		MongikClient: mongik,
		JwkSet:       defaultJwkSet,
		Config: Config{
			Mode: MIDDLEWARE,
		},
	}
}

func NewRoleCheckerClient(role *string) *RoleCheckerHandler {
	return &RoleCheckerHandler{
		Role: *role,
	}
}
