package main

import (
	//	"errors"
	"fmt"
	//	"os"
	//	"slices"
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
	Verse_s string
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
	huh.NewSelect[string]().
		Options(huh.NewOptions(Bookid(bibleargs.Bible_s)...)...).
		Title(fmt.Sprintf("Choose a Book from the %s", bibleargs.Bible_s)).
		Value(&bibleargs.Book_s).
		Run()

	// Choose a Chapter of the Bible
	huh.NewSelect[string]().
		Options(huh.NewOptions(Chapid(bibleargs.Bible_s, bibleargs.Book_s)...)...).
		Title(fmt.Sprintf("Choose a Chapter from the Book of %s", bibleargs.Book_s)).
		Value(&bibleargs.Chap_s).
		Run()

	// Choose a Verse from the Chapter of the Bible Selected
	huh.NewSelect[string]().
		Options(huh.NewOptions(Verseid(bibleargs.Bible_s, bibleargs.Book_s, bibleargs.Chap_s)...)...).
		Title(fmt.Sprintf("Choose a Chapter from %s %s", bibleargs.Book_s, bibleargs.Chap_s)).
		Value(&bibleargs.Verse_s).
		Run()

	fmt.Printf("%s", Biblecontent(bibleargs.Bible_s, bibleargs.Book_s, bibleargs.Chap_s, bibleargs.Verse_s))
}
