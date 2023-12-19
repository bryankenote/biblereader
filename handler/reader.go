package handler

import (
	"context"
	biblev1 "github.com/bryankenote/bibleapi/codegen/pb/bible/v1"
	"github.com/bryankenote/bibleapi/codegen/pb/bible/v1/biblev1connect"
	"github.com/bryankenote/biblereader/model"
	"github.com/bryankenote/biblereader/view/reader"
	utils "github.com/bryankenote/bibleutils"
	"log"
	"strconv"

	"connectrpc.com/connect"
	"github.com/labstack/echo/v4"
)

type ReaderHandler struct {
	BibleClient biblev1connect.BibleServiceClient
}

func (r ReaderHandler) HandleReaderShow(c echo.Context) error {

	translation := c.FormValue("translation")
	if translation == "" {
		translation = "BSB"
	}
	book := c.FormValue("book")
	if book == "" {
		book = "Genesis"
	}
	chapter := c.FormValue("chapter")
	if chapter == "" {
		chapter = "1"
	}
	prev := c.FormValue("prev")
	next := c.FormValue("next")

	chapterNum, err := strconv.Atoi(chapter)
	if err != nil {
		log.Println(err)
		return nil
	}

	if next == "true" && chapterNum < utils.GetTotalChapters()[book] {
		chapterNum += 1
	} else if prev == "true" && chapterNum > 1 {
		chapterNum -= 1
	}

	res, err := r.BibleClient.GetChapter(context.Background(), connect.NewRequest(&biblev1.GetChapterRequest{Translation: translation, Book: book, Chapter: int32(chapterNum)}))
	if err != nil {
		log.Println(err)
		return nil
	}

	data := model.Reader{
		Books:          utils.GetBookNames(),
		Verses:         res.Msg.Verses,
		Translation:    res.Msg.Verses[0].Translation,
		Book:           res.Msg.Verses[0].Book,
		Chapter:        res.Msg.Verses[0].Chapter,
		HasPrevChapter: chapterNum > 1,
		HasNextChapter: chapterNum < utils.GetTotalChapters()[book],
	}

	return render(c, reader.Show(data))
}
