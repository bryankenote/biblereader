package model

import biblev1 "github.com/bryankenote/bibleapi/codegen/pb/bible/v1"

type Reader struct {
	Books          []string
	Verses         []*biblev1.Verse
	Translation    string
	Book           string
	Chapter        int32
	HasPrevChapter bool
	HasNextChapter bool
}
