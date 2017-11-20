package misc

import "net/http"

func main() {
	// http.Get :: String -> JSON
	var renderPage = curry(func(makes, models) { /* render page */  })
	// return <div>html with makes and models ULs</div>
	Task.Of(renderPage).Ap(http.Get("/makes")).Ap(http.Get("/models"))
}