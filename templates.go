package main

func index() string {
	var index = `<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.}}</title>
	</head>
	<body>
		<p>Välkommen till Kamratposten</p>
    </body>
</html>`
	return index
}

func post() string {
	var post = ``
	return post
}

func login() string {
	var login = ``
	return login
}

func submitForm() string {
	var submitForm = `{{if .Success}}
	<h1>Message posted OK.</h1>
{{else}}
	<h1>Submit</h1>
	<form method="POST">
		<label>URL:</label><br />
		<input type="text" name="URL"><br />
		<label>Title:</label><br />
		<input type="text" name="Title"><br />
		<label>Text:</label><br />
		<textarea name="Text"></textarea><br />
		<input type="submit">
	</form>
{{end}}`
	return submitForm
}
