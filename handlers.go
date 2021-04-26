package main

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

func DefaultEndpointHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		err = templ.Execute(w, nil)
	}
}

func CreationEndpointHandler(w http.ResponseWriter, r *http.Request) {
	_, err := url.ParseRequestURI(r.FormValue("urlInputField"))
	if err != nil {
		log.Printf("You've entered invalid URL.")
	} else {
		err = r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Printf("ParseForm Error: ", err)
		} else {
			uniqueId := CreateId()
			myUrl := MyUrl{
				UniqueId: uniqueId,
				LongUrl:  r.FormValue("urlInputField"),
				ShortUrl: "http://localhost" + port + "/" + uniqueId,
			}
			AddToStorage(myUrl)
			w.Write([]byte(myUrl.ShortUrl))
			http.Redirect(w, r, "/create", http.StatusSeeOther)
		}
	}
}

func RedirectEndpointHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var u MyUrl

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("URLs"))
		v := b.Get([]byte(vars["id"]))
		if v == nil {
			fmt.Println("key does not exist")
			t, _ := template.ParseFiles("notfound.html")
			t.Execute(w, nil)
		}
		json.Unmarshal(v, &u)

		return nil
	})

	if err != nil {
		fmt.Printf(err.Error())
	}

	http.Redirect(w, r, u.LongUrl, http.StatusSeeOther)
}
