package main

import (
	utils "bibleutils"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	biblev1 "bibleapi/src/codegen/pb/bible/v1"
	"bibleapi/src/codegen/pb/bible/v1/biblev1connect"

	"connectrpc.com/connect"
	// "github.com/a-h/templ"
)

type PageData struct {
	Books          []string
	Verses         []*biblev1.Verse
	Translation    string
	Book           string
	Chapter        int32
	HasPrevChapter bool
	HasNextChapter bool
}

func main() {
	fmt.Println("Listening on :8080 ...")

	client := biblev1connect.NewBibleServiceClient(http.DefaultClient, "http://localhost:8000", connect.WithGRPC())
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		translation := r.PostFormValue("translation")
		if translation == "" {
			translation = "BSB"
		}
		book := r.PostFormValue("book")
		if book == "" {
			book = "Genesis"
		}
		chapter := r.PostFormValue("chapter")
		if chapter == "" {
			chapter = "1"
		}
		prev := r.PostFormValue("prev")
		next := r.PostFormValue("next")

		chapterNum, err := strconv.Atoi(chapter)
		if err != nil {
			log.Println(err)
			return
		}

		if next == "true" && chapterNum < utils.GetTotalChapters()[book] {
			chapterNum += 1
		} else if prev == "true" && chapterNum > 1 {
			chapterNum -= 1
		}

		res, err := client.GetChapter(context.Background(), connect.NewRequest(&biblev1.GetChapterRequest{Translation: translation, Book: book, Chapter: int32(chapterNum)}))
		if err != nil {
			log.Println(err)
			return
		}

		data := PageData{
			Books:          utils.GetBookNames(),
			Verses:         res.Msg.Verses,
			Translation:    res.Msg.Verses[0].Translation,
			Book:           res.Msg.Verses[0].Book,
			Chapter:        res.Msg.Verses[0].Chapter,
			HasPrevChapter: chapterNum > 1,
			HasNextChapter: chapterNum < utils.GetTotalChapters()[book],
		}

		// component := Index(data)
		// templ.Handler(component)
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, data)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

/* TODOS:
 *	dockerize
 *	mysql
 *	accessibility
 *	improve importing
 *	book wrap
 *	infinite scroll
 *	unhardcode ports
 *	documentation
 */
