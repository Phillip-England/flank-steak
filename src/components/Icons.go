package components

func HamburgerIcon() string {
	return `	
		<div id="nav-hamburger" class="text-black self-center m-4">
			<svg class="w-6 h-6" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" onclick="openNav()">
				<path fill="currentColor" d="M4 6h16v2H4zm0 5h16v2H4zm0 5h16v2H4z"/>
			</svg>
		</div>
	`
}

func CloseIcon() string {
	return `
		<div id="nav-x" class="text-black self-center m-4 hidden">
			<svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24" onclick="closeNav()">
				<path d="M19.07 16.24l-4.83-4.83 4.83-4.83-1.41-1.41-4.83 4.83-4.83-4.83-1.41 1.41 4.83 4.83-4.83 4.83 1.41 1.41 4.83-4.83 4.83 4.83 1.41-1.41z"/>
			</svg>
		</div>
	`
}