package main

import (
	"database/sql"
	//	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	//	"io"
	"log"
	//	"net/http"
	"os"
	"sort"
	//"strings"
	// "time"
)

type BibleId struct {
	Id   string
	Name string
}

type Lang struct {
	id              string
	name            string
	nameLocal       string
	script          string
	scriptDirection string
}

type BookId = BibleId

type ChapId = BibleId

var (
	language_code string
	language      string
	name          string
	bibleapi      bool   = false
	dbName        string = fmt.Sprintf("file:%s/bin/test.db", os.Getenv("GOPATH"))
)

func removeDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func Languages() []string {
	langIds := []string{}

	db, err := sql.Open("libsql", dbName)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select language from translations")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&language)
		if err != nil {
			log.Fatal(err)
		}
		langIds = append(langIds, language)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dupremlangs := removeDuplicate(langIds)

	sort.Slice(dupremlangs, func(i, j int) bool {
		return dupremlangs[i] < dupremlangs[j]
	})
	return dupremlangs
}

func Bibleid(langid string) []string {
	bibleIds := []string{}

	db, err := sql.Open("libsql", dbName)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select name from translations where language = ?", langid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		bibleIds = append(bibleIds, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return bibleIds
}

func Bookid(biblId string) []string {
	bookIds := []string{}

	db, err := sql.Open("libsql", dbName)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select distinct book from verses s where s.translation_id = (select t.id from translations t where t.name = ?)", biblId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		bookIds = append(bookIds, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return bookIds
}

func Chapid(biblId string, bookId string) []string {
	chapIds := []string{}

	db, err := sql.Open("libsql", dbName)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select distinct chapter from verses s where s.translation_id = (select t.id from translations t where t.name = ?) and s.book = ?", biblId, bookId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		chapIds = append(chapIds, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return chapIds
}

func Verseid(biblId string, bookId string, chapId string) []string {
	verseIds := []string{}

	db, err := sql.Open("libsql", dbName)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("select verse from verses s where s.translation_id = (select t.id from translations t where t.name = ?) and s.book = ? and s.chapter = ?", biblId, bookId, chapId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	verseIds = append(verseIds, "Whole Chapter")
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		verseIds = append(verseIds, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return verseIds
}

func Biblecontent(bibleId string, bookId string, chapId string, verseId string) string {
	biblecontent := ""

	db, err := sql.Open("libsql", dbName)
	if err != nil {
		log.Fatal(err)
		//os.Exit(1)
	}
	rows, err := db.Query("select verse,text from verses s where s.translation_id = (select t.id from translations t where t.name = ?) and s.book = ? and s.chapter = ? and s.verse = ?", bibleId, bookId, chapId, verseId)
	if verseId == "Whole Chapter" {
		rows, err = db.Query("select verse,text from verses s where s.translation_id = (select t.id from translations t where t.name = ?) and s.book = ? and s.chapter = ?", bibleId, bookId, chapId)
	}

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&language_code, &language)
		if err != nil {
			log.Fatal(err)
		}
		biblecontent = fmt.Sprintf("%s %s %s", biblecontent, language_code, language)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return biblecontent

}

func shiftfirsttoend[T any](s []T) []T {
	if len(s) == 0 {
		return s
	}
	z := append(s, s[0])
	return z[1:]
}
