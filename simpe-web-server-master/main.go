package main

import (
        "net/http"
	"time"
	"html/template"
	"strconv"
	"regexp"
	"log"
	"github.com/kabukky/httpscerts"
	"fmt"
)

//like hash table
var dispatch map[string]func(http.ResponseWriter, *http.Request)

type myHandler struct{}
type Add_result struct {
	Add   float64
	Error  string
}

type text_pag struct{
	Title  string
	Enc    string
	Text   string
	Button string
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	var path string
	path = r.URL.Path[1:]

	//dispatch[path] => pointer for function, status(nil, true)
	if h, ok := dispatch[path]; ok{
		h(w, r)
		return
	}

	//if not in dispatch

	//time, used to cache
	current_time := time.Now().Local()
	w.Header().Set("Date", current_time.Format(time.RFC1123))

	//name server and content-type
	w.Header().Set("Server", "Server-Netpro")
	w.Header().Set("Content-type", "text/html")
	
	//status-line
	w.WriteHeader(http.StatusOK)

	t, _ := template.ParseFiles("static/index.html")

	t.Execute(w, nil)
}

func tambah(w http.ResponseWriter, r *http.Request){
	var add  float64 = 0.0
	var error string  = " " 
	r.ParseForm()
	
	//time
	current_time := time.Now().Local()
	w.Header().Set("Date", current_time.Format(time.RFC1123))

	w.Header().Set("Server", "Server-Netpro")
	w.Header().Set("Content-type", "text/html")
	w.WriteHeader(http.StatusOK)
	
	t, _ := template.ParseFiles("static/adder.html")
	
	if r.Method == "GET"{
				
		num1s := r.Form["num1"]
		num2s := r.Form["num2"]
		if len(num1s)==1 && len(num2s) ==1 {
			num1, err1 := strconv.ParseFloat(num1s[0], 64)
			num2, err2 := strconv.ParseFloat(num2s[0], 64)

			if err1 == nil && err2 == nil{
				add = num1 + num2
			} else {
				error = "Argument Invalid"
			}
		}
	}
	
	a := Add_result{Add: add, Error: error}
	t.Execute(w, a)
}

func pindah(w http.ResponseWriter, r *http.Request){
	
	current_time := time.Now().Local()
	w.Header().Set("Date", current_time.Format(time.RFC1123))

	w.Header().Set("Server", "Server-Netpro")
	w.Header().Set("Content-type", "text/html")

	//location moved
	w.Header().Set("Location", "https://github.com/mazayazd")

	w.WriteHeader(http.StatusMovedPermanently)
}

func teks(w http.ResponseWriter, r *http.Request){
	current_time := time.Now().Local()
	w.Header().Set("Date", current_time.Format(time.RFC1123))

	//name server and content-type
	w.Header().Set("Server", "Server-Netpro")
	w.Header().Set("Content-type", "text/html")
	
	w.WriteHeader(http.StatusOK)

	t, _ := template.ParseFiles("static/text.html")

	indonesia      := text_pag{Title: "Teks", Enc: "id-ID", Text: " Selamat Datang", Button: "kembali"}
	english     := text_pag{Title: "Text",  Enc: "en-US", Text: "Hello Selamat Datang di Netpro",      Button: "Back"}
	
	accept_lang := r.Header.Get("Accept-Language")
	re          := regexp.MustCompile("[a-z]*-[A-Z]*")
	lang        := re.FindAllString(accept_lang, -1)

	if len(lang) >= 1{
		language := lang[0]
		if language == "id-ID"{
			t.Execute(w, indonesia)
		}  else {
			t.Execute(w, english)
		}
	}
}

/* func StartNonTLSServer() {

    mux := new(http.ServeMux)
    mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("Redirecting to https://localhost/")
        http.Redirect(w, r, "https://localhost/", http.StatusTemporaryRedirect)
    }))

    http.ListenAndServe(":8383", mux)
} */

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi Network Programming!")
}

func main() {

	server := http.Server{
		Addr:    ":443",
		Handler:  &myHandler{},
}

	dispatch = make(map[string]func(http.ResponseWriter, *http.Request))
	
	dispatch["Adder"]  = tambah
	dispatch["moved"]  = pindah
	dispatch["text"]   = teks
	
	err := server.ListenAndServe()

	if err != nil {
        err = httpscerts.Generate("cert.pem", "key.pem", "192.168.66.129:8383")
        if err != nil {
            log.Fatal("Error: Couldn't create https certs.")
        }
    }
	http.HandleFunc("/", handler)
    	http.ListenAndServeTLS(":8383", "cert.pem", "key.pem", nil)
}

//show files
//http.ServeFile(w, r, "static/index.html") 
