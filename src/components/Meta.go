package components

func Meta() string {
	return `
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
		<script src="//unpkg.com/alpinejs" defer></script>
		<script src="https://unpkg.com/htmx.org@1.9.2" integrity="sha384-L6OqL9pRWyyFU3+/bjdSri+iIphTN/bvYyM37tICVyOJkWZLpP2vGn6VUEXgzg6h" crossorigin="anonymous"></script>
		<link rel="stylesheet" href="/static/output.css">
		<script src="/static/app.js"></script>
		<script src="https://unpkg.com/hyperscript.org@0.9.9"></script>
	`
}