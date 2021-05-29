module github.com/dantas/gocoindesktop

go 1.23

require (
	fyne.io/fyne/v2 v2.0.3
	github.com/getlantern/systray v1.1.0
	github.com/gocolly/colly/v2 v2.1.0
)

replace (
	fyne.io/fyne/v2 v2.0.3 => github.com/dantas/fyne/v2 v2.0.4-0.20210529211031-9bfde19a0d95
	github.com/getlantern/systray v1.1.0 => github.com/dantas/systray v1.1.1-0.20210526043731-4b434817eab3
)
