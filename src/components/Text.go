package components

import (
	"html"
)

func HeaderLg(text string) string {
    safeText := html.EscapeString(text)
    return `<h1 class="text-lg font-serif p-4">` + safeText + `</h1>`
}
