//go:build windows

package main

import (
	_ "embed"
	"fmt"
	"strings"
	"sync"

	"SnippetsDome/internal/domain"

	"github.com/getlantern/systray"
)

//go:embed build/windows/icon.ico
var trayIcon []byte

type trayController struct {
	app  *App
	once sync.Once
}

func newTrayController(app *App) *trayController {
	return &trayController{app: app}
}

func (t *trayController) start() {
	t.once.Do(func() {
		go systray.Run(t.onReady, func() {})
	})
}

func (t *trayController) stop() {
	systray.Quit()
}

func (t *trayController) isSupported() bool { return true }

func (t *trayController) onReady() {
	systray.SetIcon(trayIcon)
	systray.SetTooltip("SnippetsDome")

	showItem := systray.AddMenuItem("Open SnippetsDome", "Show the application window")
	go func() {
		for range showItem.ClickedCh {
			t.app.showWindow()
		}
	}()

	systray.AddSeparator()
	for _, snippet := range t.traySnippets() {
		snippet := snippet
		item := systray.AddMenuItem("Copy: "+traySnippetTitle(snippet.Title), "Copy this snippet to the clipboard")
		go func() {
			for range item.ClickedCh {
				t.app.copySnippetToClipboard(snippet.Code)
			}
		}()
	}

	systray.AddSeparator()
	quitItem := systray.AddMenuItem("Quit SnippetsDome", "Close the application completely")
	go func() {
		for range quitItem.ClickedCh {
			t.app.quitFromTray()
		}
	}()
}

func (t *trayController) traySnippets() []domain.Snippet {
	snippets, err := t.app.snippets.List()
	if err != nil {
		return nil
	}

	limit := t.app.snippets.TraySnippetLimit()
	if len(snippets) > limit {
		snippets = snippets[:limit]
	}
	return append([]domain.Snippet(nil), snippets...)
}

func traySnippetTitle(title string) string {
	trimmed := strings.TrimSpace(title)
	if trimmed == "" {
		return "Untitled snippet"
	}
	const maximumLength = 48
	if len(trimmed) > maximumLength {
		return fmt.Sprintf("%s…", trimmed[:maximumLength-1])
	}
	return trimmed
}
