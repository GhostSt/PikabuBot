package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/charmap"
)

type Post struct {
	postId string
	author string
	title  string
	url    string
}

func main() {
	url := "https://pikabu.ru/profile/boss1w"

	setup()

	doc, err := goquery.NewDocument(url)

	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div.story").Each(func(i int, s *goquery.Selection) {
		storyId, _ := s.Attr("data-story-id")
		author := s.Find("div.story__author").Text()
		title := s.Find("div.story__header-title").Text()
		url, _ := s.Find("div.story__header-title a").Attr("href")

		post := Post{
			storyId,
			convertWin1251ToUtf8(author),
			convertWin1251ToUtf8(title),
			convertWin1251ToUtf8(url),
		}


		savePost(post)
	})

}

func convertWin1251ToUtf8(string string) string {
	returnedString, _ := charmap.Windows1251.NewDecoder().String(string)

	return returnedString
}

func savePost(post Post)  {
	db := reg.db

	transaction, err := db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	stmt, err := transaction.Prepare("INSERT INTO post(id, author, title, url) VALUES(?, ?, ?, ?)")

	if err != nil {
		panic(err)
		log.Fatal(err)
	}

	defer stmt.Close()

	/**
	_, err = stmt.Exec(post.postId, post.author, post.title, post.url)

	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	*/
	transaction.Commit()

}