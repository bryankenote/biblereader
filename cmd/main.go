package main

import (
	"net/http"

	"github.com/bryankenote/bibleapi/codegen/pb/bible/v1/biblev1connect"
	"github.com/bryankenote/biblereader/handler"

	"connectrpc.com/connect"
	"github.com/labstack/echo/v4"
)

func main() {
	bibleV1client := biblev1connect.NewBibleServiceClient(http.DefaultClient, "http://localhost:8000", connect.WithGRPC())

	app := echo.New()

	readerHandler := handler.ReaderHandler{BibleClient: bibleV1client}
	app.GET("/", readerHandler.HandleReaderShow)

	app.Start(":8080")
}
