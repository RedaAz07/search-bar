package helpers
func PageDeleted() string {



return 	`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Status Page</title>
    <link rel="stylesheet" href="/static/status_Page.css">
    <link rel="shortcut icon" href="/static/images/error logo.png" type="image/x-icon">
</head>
<body>
    <div class="container">
        <div class="error-content">
            <h1>500</h1>
            <h2>oops!</h2>
            <p>Something went wrong on our end. We're working on fixing it—please try again later!</p>
            <div class="buttons">
                <a href="/" class="btn">Go home</a>
                <a href="" class="btn">Contact us</a>
            </div>
        </div>
        <div class="image">
            <img src="/static/images/errrrrrrrrr.png" alt="404 illustration">
        </div>
    </div>
</body>
</html>
	`
}