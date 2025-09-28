package main

// EmailAddress ...
type EmailAddress struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

// EmailAttachment ...
type EmailAttachment struct {
	Filename    string `json:"filename"`
	ContentType string `json:"content_type"`
	Data        string `json:"data"`
}

// EmailEmbeddedFile ...
type EmailEmbeddedFile struct {
	CID         string `json:"cid"`
	ContentType string `json:"content_type"`
	Data        string `json:"data"`
}

// EmailMessage ...
type EmailMessage struct {
	References []string `json:"references,omitempty"`
	SPFResult  string   `json:"spf,omitempty"`

	ID      string `json:"id,omitempty"`
	Date    string `json:"date,omitempty"`
	Subject string `json:"subject,omitempty"`

	ResentDate string `json:"resent_date,omitempty"`
	ResentID   string `json:"resent_id,omitempty"`

	Body struct {
		Text string `json:"text,omitempty"`
		HTML string `json:"html,omitempty"`
	} `json:"body"`

	Addresses struct {
		From      *EmailAddress   `json:"from"`
		To        *EmailAddress   `json:"to"`
		ReplyTo   []*EmailAddress `json:"reply_to,omitempty"`
		Cc        []*EmailAddress `json:"cc,omitempty"`
		Bcc       []*EmailAddress `json:"bcc,omitempty"`
		InReplyTo []string        `json:"in_reply_to,omitempty"`

		ResentFrom *EmailAddress   `json:"resent_from,omitempty"`
		ResentTo   []*EmailAddress `json:"resent_to,omitempty"`
		ResentCc   []*EmailAddress `json:"resent_cc,omitempty"`
		ResentBcc  []*EmailAddress `json:"resent_bcc,omitempty"`
	} `json:"addresses"`

	Attachments   []*EmailAttachment   `json:"attachments,omitempty"`
	EmbeddedFiles []*EmailEmbeddedFile `json:"embedded_files,omitempty"`
}

// PostmarkMessage represents the Postmark webhook format
type PostmarkMessage struct {
	MessageID     string                   `json:"MessageID"`
	Date          string                   `json:"Date"`
	Subject       string                   `json:"Subject"`
	From          string                   `json:"From"`
	FromName      string                   `json:"FromName"`
	FromFull      PostmarkAddress          `json:"FromFull"`
	To            string                   `json:"To"`
	ToFull        []PostmarkAddress        `json:"ToFull"`
	Cc            string                   `json:"Cc,omitempty"`
	CcFull        []PostmarkAddress        `json:"CcFull,omitempty"`
	Bcc           string                   `json:"Bcc,omitempty"`
	BccFull       []PostmarkAddress        `json:"BccFull,omitempty"`
	ReplyTo       string                   `json:"ReplyTo,omitempty"`
	TextBody      string                   `json:"TextBody,omitempty"`
	HtmlBody      string                   `json:"HtmlBody,omitempty"`
	StrippedTextReply string               `json:"StrippedTextReply,omitempty"`
	Tag           string                   `json:"Tag,omitempty"`
	Headers       []PostmarkHeader         `json:"Headers,omitempty"`
	Attachments   []PostmarkAttachment     `json:"Attachments,omitempty"`
}

// PostmarkAddress represents an email address in Postmark format
type PostmarkAddress struct {
	Email string `json:"Email"`
	Name  string `json:"Name"`
}

// PostmarkHeader represents email headers in Postmark format
type PostmarkHeader struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

// PostmarkAttachment represents an attachment in Postmark format
type PostmarkAttachment struct {
	Name        string `json:"Name"`
	Content     string `json:"Content"`
	ContentType string `json:"ContentType"`
	ContentLength int  `json:"ContentLength"`
}
