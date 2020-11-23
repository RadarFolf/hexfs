package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"strings"
)

type FilterStatus int

const (
	FilterPass FilterStatus = iota
	FilterFail
	FilterSanitize
)

// FilterCheck runs the file's mime type through a series of checks.
// FilterPass means the mime type is fine and no additional action should be taken.
// FilterFail means a reponse was already returned, and the caller should terminate its function.
// FilterSanitize means the file's Content-Type header returned to the client should be changed to text/plain.
func (b *BaseHandler)FilterCheck(ctx *fasthttp.RequestCtx, mimeType string) FilterStatus {
	if len(b.Config.Security.Filter.Blacklist) > 0 {
		for _, t := range b.Config.Security.Filter.Blacklist {
			if strings.HasPrefix(mimeType, t) && len(t) > 0 {
				SendTextResponse(ctx, "File type blacklisted.", fasthttp.StatusForbidden)
				return FilterFail
			}
		}
	}
	if len(b.Config.Security.Filter.Whitelist) > 0 {
		fmt.Println("Whitelist")
		for i, t := range b.Config.Security.Filter.Whitelist {
			if strings.HasPrefix(mimeType, t) && len(t) > 0 {
				break
			}
			if i + 1 == len(b.Config.Security.Filter.Whitelist) {
				SendTextResponse(ctx, "File type is not whitelisted.", fasthttp.StatusForbidden)
				return FilterFail
			}
		}
	}
	if len(b.Config.Security.Filter.Sanitize) > 0 {
		for _, t := range b.Config.Security.Filter.Sanitize {
			if strings.HasPrefix(mimeType, t) && len(t) > 0 {
				return FilterSanitize
			}
		}
	}
	return FilterPass
}