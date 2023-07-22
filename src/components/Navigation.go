package components

import "html/template"


func Navbar(banner string) template.HTML {
	navbar := `
		<div class="flex fixed h-16 top-0 w-full flex-row justify-between z-30 bg-gray-100 shadow-md text-gray-700">
			`+ HeaderLg(banner) +`
			`+ HamburgerIcon() +`
			`+ CloseIcon() +`
		</div>
		`+ Spacer() +`
		`+ NavOverlay() +`
	`
	return template.HTML(navbar)
}

func NavOverlay() string {
	return `
		<div id="nav-overlay" class="fixed top-0 z-20 w-full hidden h-full opacity-50 bg-black" onclick="closeNav()"></div>
	`
}