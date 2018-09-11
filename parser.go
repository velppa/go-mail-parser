package main

import (
	"fmt"
	"net/mail"
	"os"

	"flag"
	"github.com/jhillyerd/go.enmime"
)

func main() {
	// Command line arg
	filePath := flag.String("file", "", "path to .eml file")
	// Flag parsing
	flag.Parse()
	// Checking file flag
	if *filePath == "" {
		panic("You should provide path to .eml file by using -file flag")
	}
	file, _ := os.Open(*filePath)
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

	// The HTML body is stored in mime.HTML
	fmt.Printf("HTML Body: %v chars\n", len(mime.HTML))

	// mime.Inlines is a slice of inlined attachments
	fmt.Printf("Inlines: %v\n", len(mime.Inlines))

	// mime.Attachments contains the non-inline attachments
	fmt.Printf("Attachments: %v\n", len(mime.Attachments))
}
