package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/charmap"
	"database/sql"
	"fmt"
	"regexp"
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

	posts := []Post{}

	doc.Find("div.story").Each(func(i int, s *goquery.Selection) {
		storyId, _ := s.Attr("data-story-id")
		author := s.Find("a.story__author").Text()
		url, _ := s.Find("div.story__header-title a").Attr("href")

		title := s.Find("div.story__header-title").Text()
		clearTitlePattern := regexp.MustCompile(`[\[\]\(\)\n\t]*`)

		title = clearTitlePattern.ReplaceAllString(title, "")

		var post = Post{
			storyId,
			convertWin1251ToUtf8(author),
			convertWin1251ToUtf8(title),
			convertWin1251ToUtf8(url),
		}

		posts = append(posts, post)
		/**
		postIDPattern := regexp.MustCompile(`^\d+$`)

		fmt.Println(post.postId)
		if isPostExists(registry.db, post.postId) {
			return
		}

		fmt.Println(post)
		if postIDPattern.MatchString(post.postId) {
			savePost(post, registry)

			addedPosts++
		}
		*/
	})

	post := posts[0]
	savePost(post)
	fmt.Println(post)

	message := fmt.Sprintf("%s: [%s](%s)", post.author, post.title, post.url)
	fmt.Println(message)

	res, err := sendMessage(message)

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
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
	}

	defer stmt.Close()

	/**
	_, err = stmt.Exec(post.postId, post.author, post.title, post.url)

	if err != nil {
		panic(err)
	}
	*/
	transaction.Commit()
}

func isPostExists(db *sql.DB, postId string) bool {
	rows, err := db.Query("SELECT * FROM post WHERE id = ?", postId);
	if err != nil {
		panic(err)
	}

	var post Post

	for rows.Next() {
		rows.Scan(post.postId, post.author, post.title, post.url)

		return true
	}

	return false
}