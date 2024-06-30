package i18n

import (
	"embed"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pkg/errors"
	"golang.org/x/text/language"
)

const GinLocalizerKey = "localizer"

var bundle *i18n.Bundle

//go:embed messages.*.toml
var LocaleFS embed.FS

func Init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	registryLocale("en")
	registryLocale("zh-CN")
}

func registryLocale(locale string) {
	_, err := bundle.LoadMessageFileFS(LocaleFS, fmt.Sprintf("messages.%s.toml", locale))
	if err != nil {
		panic(errors.Wrap(err, "register locale error"))
	}
}

func NewLocalizer(accept string) *i18n.Localizer {
	lang := language.Make(accept)
	return i18n.NewLocalizer(bundle, lang.String())
}

func Translate(c *gin.Context, id string, args ...interface{}) string {
	localizer := c.MustGet(GinLocalizerKey).(*i18n.Localizer)
	return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: id, TemplateData: args})
}
