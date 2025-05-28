package sitemap

import (
	"bytes"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sabloger/sitemap-generator/smg"
)

type SitemapHandler struct {
	router fiber.Router
	logger *slog.Logger
}

func NewSitemapHandler(router fiber.Router, logger *slog.Logger) {
	h := &SitemapHandler{
		router: router,
		logger: logger,
	}
	h.router.Get("/sitemap.xml", h.sitemap)
}

func (h SitemapHandler) sitemap(ctx *fiber.Ctx) error {
	now := time.Now().UTC()

	sm := smg.NewSitemap(true) // The argument is PrettyPrint which must be set on initializing
	sm.SetName("sitemap")      // Optional
	sm.SetHostname("https://ostkost.github.io/go-ps-fiber/")
	sm.SetOutputPath("./some/path")
	sm.SetLastMod(&now)
	sm.SetCompress(false)     // Default is true
	sm.SetMaxURLsCount(25000) // Default maximum number of URLs in each file is 50,000 to break

	sm.Add(&smg.SitemapLoc{
		Loc:        "/",
		LastMod:    &now,
		ChangeFreq: smg.Daily,
		Priority:   0.9,
	})
	sm.Add(&smg.SitemapLoc{
		Loc:        "/categories",
		LastMod:    &now,
		ChangeFreq: smg.Daily,
		Priority:   0.8,
	})
	sm.Add(&smg.SitemapLoc{
		Loc:        "/login",
		LastMod:    &now,
		ChangeFreq: smg.Weekly,
		Priority:   0.7,
	})
	sm.Add(&smg.SitemapLoc{
		Loc:        "/register",
		LastMod:    &now,
		ChangeFreq: smg.Weekly,
		Priority:   0.5,
	})
	sm.Add(&smg.SitemapLoc{
		Loc:        "/news-create",
		LastMod:    &now,
		ChangeFreq: smg.Daily,
		Priority:   0.6,
	})
	sm.Add(&smg.SitemapLoc{
		Loc:        "/404",
		LastMod:    &now,
		ChangeFreq: smg.Monthly,
		Priority:   0.1,
	})

	sm.Finalize()
	var buf bytes.Buffer
	if _, err := sm.WriteTo(&buf); err != nil {
		return err
	}
	ctx.Set("Content-Type", "application/xml")
	return ctx.Send(buf.Bytes())
}
