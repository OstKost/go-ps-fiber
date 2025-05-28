package sitemap

import (
	"bytes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sabloger/sitemap-generator/smg"
)

type SitemapHandler struct {
	router fiber.Router
}

func NewSitemapHandler(router fiber.Router) {
	h := &SitemapHandler{
		router: router,
	}
	h.router.Get("/sitemap.xml", h.sitemap)
}

func (h SitemapHandler) sitemap(ctx *fiber.Ctx) error {
	sm := smg.NewSitemap(false)
	sm.SetHostname("https://ostkost.github.io/go-ps-fiber/")
	now := time.Now().UTC()
	sm.SetLastMod(&now)
	sm.SetCompress(false)
	sm.Add(&smg.SitemapLoc{
		Loc:        "/",
		LastMod:    &now,
		ChangeFreq: smg.Daily,
		Priority:   0.8,
	})
	sm.Add(&smg.SitemapLoc{
		Loc:        "/login",
		LastMod:    &now,
		ChangeFreq: smg.Weekly,
		Priority:   0.6,
	})
	sm.Add(&smg.SitemapLoc{
		Loc:        "/register",
		LastMod:    &now,
		ChangeFreq: smg.Weekly,
		Priority:   0.5,
	})
	sm.Add(&smg.SitemapLoc{
		Loc:        "/404",
		LastMod:    &now,
		ChangeFreq: smg.Monthly,
		Priority:   0.2,
	})
	sm.Finalize()
	var buf bytes.Buffer
	if _, err := sm.WriteTo(&buf); err != nil {
		return err
	}
	ctx.Set("Content-Type", "application/xml")
	return ctx.Send(buf.Bytes())
}
