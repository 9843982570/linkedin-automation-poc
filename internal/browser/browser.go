package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func Start(headless bool) (*rod.Browser, error) {
	u := launcher.New().
		Headless(headless).
		NoSandbox(true).
		MustLaunch()

	browser := rod.New().ControlURL(u)
	err := browser.Connect()
	if err != nil {
		return nil, err
	}

	return browser, nil
}
