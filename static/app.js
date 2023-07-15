
function qs(selector) {
	return document.querySelector(selector)
}

function qsa(selector) {
	return document.querySelectorAll(selector)
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
	let hamburger = qs('#nav-hamburger')
	let menu = qs('#nav-menu')
	let x = qs('#nav-x')
	let overlay = qs('#nav-overlay')
	links = menu.querySelectorAll('li')
	links.forEach(link => {
		let href = link.children[0].getAttribute('href')
		if (href == window.location.pathname) {
			link.classList.add('bg-gray-200')
			link.classList.add('border-gray-400')
			link.children[0].classList.add('bg-gray-200')
		}
	});
	if (hamburger != null) {
		hamburger.classList.add("hidden")
	}
	if (menu != null) {
		menu.classList.remove("hidden")
	}
	if (x != null) {
		x.classList.remove("hidden")
	}
	if (overlay != null) {
		overlay.classList.remove('hidden')
	}
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












