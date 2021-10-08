package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/youthlin/logs"
	"github.com/youthlin/pub/models"
	"github.com/youthlin/t"
	"golang.org/x/text/language"
)

var log = logs.GetLogger()

// I18n is a middleware which can process translation for different request
func I18n(c *gin.Context) {
	ctx := GetCtx(c)
	log := log.Ctx(ctx)
	langs := t.Locales()
	var supported []language.Tag
	for _, lang := range langs {
		supported = append(supported, language.Make(lang))
	}
	matcher := language.NewMatcher(supported)
	accept := c.GetHeader(models.HeaderAcceptLanguage)
	userAccept, q, err := language.ParseAcceptLanguage(accept)
	log.Debug("ParseAcceptLanguage|userAccept=%v, q=%v, err=%+v", userAccept, q, err)
	matchedTag, index, confidence := matcher.Match(userAccept...)
	userLang := langs[index]
	t := t.L(userLang)
	log.Debug("userLangMatch|matched=%v, index=%v, confidence=%v | select=%v, use=%v",
		matchedTag, index, confidence, userLang, t.UsedLocale())
	c.Set(models.GinKeyT, t)
	c.Next()
}

// GetTranslations return the Translations instance of this request
func GetTranslations(c *gin.Context) *t.Translations {
	ts, ok := c.Get(models.GinKeyT)
	if !ok {
		ts = t.Global()
	}
	return ts.(*t.Translations)
}
