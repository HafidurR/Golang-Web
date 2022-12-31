package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type M map[string]interface{}

func main() {

	var tmpl, err = template.ParseGlob("views/*")
	if err != nil {
		panic(err.Error())
		return
	}

	http.Handle("/public/",
		http.StripPrefix("/public/",
			http.FileServer(http.Dir("public"))))

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"text": "Lorem ipsum dolor sit amet consectetur, adipisicing elit. Rerum architecto, labore aliquam doloribus saepe distinctio quis placeat? Sunt corporis rem exercitationem incidunt, maxime id, dolores similique aliquid iste possimus vero!. Lorem ipsum dolor sit amet consectetur adipisicing elit. Nobis modi voluptas quidem ab, tenetur, atque nisi fugit, harum vel repellat blanditiis sint dignissimos magnam libero! Quam quos ab fuga obcaecati."}

		err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"blog": "Lorem ipsum dolor sit amteur delur empsnamet consectetur, adipisicing elit. Rerum architecto, labore aliquam doloribus saepe distinctio quis placeat? Sunt corporis rem exercitationem incidunt, maxime id, dolores similique aliquid iste possimus vero!.", "blog ": "Lorem ipsum dolor sit amet consectetur, adipisicing elit. Rerum architecto, labore aliquam doloribus saepe distinctio quis placeat? Sunt corporis rem exercitationem incidunt, maxime id, dolores similique aliquid iste possimus vero!."}
		err = tmpl.ExecuteTemplate(w, "blog", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
