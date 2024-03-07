package blog

import "time"

type Entry struct {
	Id         string    `firestore:"id"`
	Date       time.Time `firestore:"date"`
	Author     string    `firestore:"author"`
	Title      string    `firestore:"title"`
	Lead       string    `firestore:"lead"`
	Tags       string    `firestore:"tags"`
	DocumentId string    `firestore:"documentId"`
}
