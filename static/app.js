
function qs(selector) {
	return document.querySelector(selector)
}

function qsa(selector) {
	return document.querySelectorAll(selector)
}

function hideAndShow(elementToHide, elementToShow) {
	elementToHide.classList.add("hidden")
	elementToShow.classList.remove('hidden')
}

function showBigLoader() {
	let bigLoader = qs("#big-loader")
	bigLoader.classList.toggle('hidden')
	let loadingOverlay = qs('#big-loader-overlay')
	loadingOverlay.classList.toggle('hidden')
}

function hideBigLoader() {
	let bigLoader = qs("#big-loader")
	let loadingOverlay = qs('#big-loader-overlay')
	if (bigLoader != null) {
		bigLoader.classList.add('hidden')
	}
	if (loadingOverlay != null) {
		loadingOverlay.classList.add('hidden')
	}
}

function openNav() {
	document.querySelector("#nav-hamburger").classList.add("hidden")
	document.querySelector("#nav-menu").classList.remove("hidden")
	document.querySelector("#nav-x").classList.remove("hidden")
	document.querySelector("#nav-overlay").classList.remove('hidden')
}

function closeNav() {
	let hamburger = qs('#nav-hamburger')
	let menu = qs('#nav-menu')
	let x = qs('#nav-x')
	let overlay = qs('#nav-overlay')
	if (hamburger != null) {
		hamburger.classList.remove("hidden")
	}
	if (menu != null) {
		menu.classList.add('hidden')
	}
	if (x != null) {
		x.classList.add("hidden")
	}
	if (overlay != null) {
		overlay.classList.add('hidden')
	}

}












