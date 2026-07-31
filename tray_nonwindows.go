//go:build !windows

package main

// trayController keeps non-Windows builds functional while the current tray
// implementation targets the Windows notification area used by this project.
type trayController struct{}

func newTrayController(app *App) *trayController {
	return &trayController{}
}

func (t *trayController) start() {}

func (t *trayController) stop() {}

func (t *trayController) isSupported() bool { return false }
