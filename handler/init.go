package handler

import (
	"github.com/FrosTiK-SD/authV2/controller"
	"github.com/FrosTiK-SD/authV2/model"
	mongik "github.com/FrosTiK-SD/mongik/models"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

type Mode string

const (
	HANDLER    Mode = "HANDLER"
	MIDDLEWARE Mode = "MIDDLEWARE"
)

type Session struct {
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
