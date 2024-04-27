package main

import (
	//	"errors"
	//	"fmt"
	//	"os"
	"slices"
	//	"strconv"
	//	"strings"
	//	"time"

	"github.com/charmbracelet/huh"
	// "github.com/charmbracelet/huh/spinner"
	// "github.com/charmbracelet/lipgloss"
	// xstrings "github.com/charmbracelet/x/exp/strings"
)

type BibleArgs struct {
	Lang_s  string
	Bible_s string
	Book_s  string
	Chap_s  string
}

func main() {
	var bibleargs BibleArgs

	// Choose a Language
	huh.NewSelect[string]().
		Options(huh.NewOptions(Languages()...)...).
		Title("Choose a Bible Language").
		Value(&bibleargs.Lang_s).
		Run()

	// Choose a Bible Version
	huh.NewSelect[string]().
		Options(huh.NewOptions(Bibleid(bibleargs.Lang_s)...)...).
		Title("Choose a Bible Version").
		Value(&bibleargs.Bible_s).
		Run()

	// Choose a Book of the Bible
	bookid := Bookid(bibleargs.Bible_s)

	var bookids []string

	for i := 0; i < len(bookid); i++ {
		bookids = slices.Insert(bookids, len(bookids), bookid[i].Name)
	}

	huh.NewSelect[string]().
		Options(huh.NewOptions(bookids...)...).
		Title("Choose a Book").
		Value(&bibleargs.Book_s).
		Run()

}
