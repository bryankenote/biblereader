package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	biblev1 "bibleapi/src/codegen/pb/bible/v1"
	"bibleapi/src/codegen/pb/bible/v1/biblev1connect"

	"connectrpc.com/connect"
)

type PageData struct {
	Verses         []*biblev1.Verse
	Translation    string
	Book           string
	Chapter        int32
	HasPrevChapter bool
	HasNextChapter bool
}

func main() {
	fmt.Println("Listening on :8080 ...")

	// TODO: gRPC
	client := biblev1connect.NewBibleServiceClient(http.DefaultClient, "http://localhost:8000")
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

		if next == "true" && chapterNum < getTotalChapters()[book] {
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
			Verses:         res.Msg.Verses,
			Translation:    res.Msg.Verses[0].Translation,
			Book:           res.Msg.Verses[0].Book,
			Chapter:        res.Msg.Verses[0].Chapter,
			HasPrevChapter: chapterNum > 1,
			HasNextChapter: chapterNum < getTotalChapters()[book],
		}
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, data)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// TODO: move to library
func getTotalChapters() map[string]int {
	chapters := make(map[string]int)

	chapters["Genesis"] = 50
	chapters["Exodus"] = 40
	chapters["Leviticus"] = 27
	chapters["Numbers"] = 36
	chapters["Deuteronomy"] = 34
	chapters["Joshua"] = 24
	chapters["Judges"] = 21
	chapters["Ruth"] = 4
	chapters["1 Samuel"] = 31
	chapters["2 Samuel"] = 24
	chapters["1 Kings"] = 22
	chapters["2 Kings"] = 25
	chapters["1 Chronicles"] = 29
	chapters["2 Chronicles"] = 36
	chapters["Ezra"] = 10
	chapters["Nehemiah"] = 13
	chapters["Esther"] = 10
	chapters["Job"] = 42
	chapters["Psalms"] = 150
	chapters["Proverbs"] = 31
	chapters["Ecclesiastes"] = 12
	chapters["Song of Solomon"] = 8
	chapters["Isaiah"] = 66
	chapters["Jeremiah"] = 52
	chapters["Lamentations"] = 5
	chapters["Ezekiel"] = 48
	chapters["Daniel"] = 12
	chapters["Hosea"] = 14
	chapters["Joel"] = 6
	chapters["Amos"] = 9
	chapters["Obadiah"] = 1
	chapters["Jonah"] = 4
	chapters["Micah"] = 7
	chapters["Nahum"] = 3
	chapters["Habakkuk"] = 3
	chapters["Zephaniah"] = 3
	chapters["Haggai"] = 2
	chapters["Zechariah"] = 14
	chapters["Malachi"] = 4

	chapters["Matthew"] = 28
	chapters["Mark"] = 16
	chapters["Luke"] = 24
	chapters["John"] = 21
	chapters["Acts"] = 28
	chapters["Romans"] = 16
	chapters["1 Corinthians"] = 16
	chapters["2 Corinthians"] = 13
	chapters["Galatians"] = 6
	chapters["Ephesians"] = 6
	chapters["Philippians"] = 4
	chapters["Colossians"] = 4
	chapters["1 Thessalonians"] = 5
	chapters["2 Thessalonians"] = 3
	chapters["1 Timothy"] = 6
	chapters["2 Timothy"] = 4
	chapters["Titus"] = 3
	chapters["Philemon"] = 1
	chapters["Hebrews"] = 13
	chapters["James"] = 5
	chapters["1 Peter"] = 5
	chapters["2 Peter"] = 3
	chapters["1 John"] = 5
	chapters["2 John"] = 1
	chapters["3 John"] = 1
	chapters["Jude"] = 1
	chapters["Revelation"] = 22

	return chapters
}

/* TODOS:
 *	1. dockerize
 *	2. mysql
 *	3. accessibility
 *	4. css
 *	5. improve importing
 *	6. book wrap
 *	7. book dropdown
 *	8. chapter dropdown
 *	9. infinite scroll
 */
