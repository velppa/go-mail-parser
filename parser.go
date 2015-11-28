package main

import (
	"fmt"
	"net/mail"
	"os"

	"github.com/jhillyerd/go.enmime"
)

func main() {
	file, _ := os.Open("/Users/rbd/Temp/7ccaa5be-acf7-45c7-a289-c6bd2cb0492a.eml")
	msg, _ := mail.ReadMessage(file)     // Read email using Go's net/mail
	mime, _ := enmime.ParseMIMEBody(msg) // Parse message body with enmime

	// Headers are in the net/mail Message
	fmt.Printf("From: %v\n", msg.Header.Get("From"))

	fmt.Printf("From: %v\n", mime.GetHeader("From"))

	// enmime can decode quoted-printable headers
	fmt.Printf("Subject: %v\n", mime.GetHeader("Subject"))

	// The plain text body is available as mime.Text
	fmt.Printf("Text Body: %v chars\n", len(mime.Text))

	fmt.Printf("Text Body: %s\n", mime.Text)

	// The HTML body is stored in mime.Html
	fmt.Printf("HTML Body: %v chars\n", len(mime.Html))

	// mime.Inlines is a slice of inlined attacments
	fmt.Printf("Inlines: %v\n", len(mime.Inlines))

	// mime.Attachments contains the non-inline attachments
	fmt.Printf("Attachments: %v\n", len(mime.Attachments))
}
