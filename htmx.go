// Package htmx provides utilities for working with HTMX-enabled web apps.
package htmx

import "net/http"

type Headers struct {
	boosted                 string
	currentURL              string
	hxRequest               string
	isHistoryRestoreRequest string
	prompt                  string
	target                  string
	triggerName             string
	trigger                 string
}

func ParseHeaders(r *http.Request) Headers {
	return Headers{
		boosted:                 r.Header.Get("HX-Boosted"),
		currentURL:              r.Header.Get("HX-Current-URL"),
		hxRequest:               r.Header.Get("HX-Request"),
		isHistoryRestoreRequest: r.Header.Get("HX-History-Restore-Request"),
		prompt:                  r.Header.Get("HX-Prompt"),
		target:                  r.Header.Get("HX-Target"),
		triggerName:             r.Header.Get("HX-Trigger-Name"),
		trigger:                 r.Header.Get("HX-Trigger"),
	}
}
