package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/calc", calcPage)
	http.HandleFunc("/calc/result", calcResult)

	fmt.Println("Сервер запущено на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home.html", nil)
}

func calcPage(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "calc.html", nil)
}

func round(val float64, decimals int) float64 {
	format := "%." + strconv.Itoa(decimals) + "f"
	v, _ := strconv.ParseFloat(fmt.Sprintf(format, val), 64)
	return v
}

func calcResult(w http.ResponseWriter, r *http.Request) {

	pc, _ := strconv.ParseFloat(r.FormValue("pc"), 64)
	vartist, _ := strconv.ParseFloat(r.FormValue("vartist"), 64)

	delta1 := 0.2
	delta2 := 0.68

	// σ1
	W1 := pc * 24 * delta1
	P1 := W1 * vartist * 1000
	W2 := pc * 24 * (1 - delta1)
	S1 := W2 * vartist * 1000

	// σ2
	W3 := pc * 24 * delta2
	P2 := W3 * vartist * 1000
	W4 := pc * 24 * (1 - delta2)
	S2 := W4 * vartist * 1000

	// прибуток
	profit := (P2 - S2) / 1000

	result := fmt.Sprintf(`
Розрахунок для σ1:

W1 = %.3f
P1 = %.3f
W2 = %.3f
S1 = %.3f


Розрахунок для σ2:

W3 = %.3f
P2 = %.3f
W4 = %.3f
S2 = %.3f


Прибуток:

%.3f тис.
`,
		round(W1, 3), round(P1, 3), round(W2, 3), round(S1, 3),
		round(W3, 3), round(P2, 3), round(W4, 3), round(S2, 3),
		round(profit, 3))

	templates.ExecuteTemplate(w, "result.html", result)
}
