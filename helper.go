package main

import (
	"net/mail"
	"strings"
)

func extractEmails(addr []*mail.Address, _ ...error) []string {
	ret := []string{}

	for _, e := range addr {
		ret = append(ret, e.Address)
	}

	return ret
}

func transformStdAddressToEmailAddress(addr []*mail.Address) []*EmailAddress {
	ret := []*EmailAddress{}

	for _, e := range addr {
		ret = append(ret, &EmailAddress{
			Address: e.Address,
			Name:    e.Name,
		})
	}

	return ret
}

func transformStdAddressToPostmarkAddress(addr []*mail.Address) []PostmarkAddress {
	ret := []PostmarkAddress{}

	for _, e := range addr {
		ret = append(ret, PostmarkAddress{
			Email: e.Address,
			Name:  e.Name,
		})
	}

	return ret
}

func emailAddressesToPostmarkAddresses(addrs []*EmailAddress) []PostmarkAddress {
	ret := []PostmarkAddress{}

	for _, e := range addrs {
		ret = append(ret, PostmarkAddress{
			Email: e.Address,
			Name:  e.Name,
		})
	}

	return ret
}

func emailAddressesToCommaSeparated(addrs []*EmailAddress) string {
	emails := []string{}
	for _, addr := range addrs {
		if addr.Name != "" {
			emails = append(emails, addr.Name+" <"+addr.Address+">")
		} else {
			emails = append(emails, addr.Address)
		}
	}
	return strings.Join(emails, ", ")
}

func convertToPostmarkFormat(msg *EmailMessage) *PostmarkMessage {
	postmarkMsg := &PostmarkMessage{
		MessageID: msg.ID,
		Date:      msg.Date,
		Subject:   msg.Subject,
		TextBody:  msg.Body.Text,
		HtmlBody:  msg.Body.HTML,
	}

	if msg.Addresses.From != nil {
		postmarkMsg.From = msg.Addresses.From.Address
		postmarkMsg.FromName = msg.Addresses.From.Name
		postmarkMsg.FromFull = PostmarkAddress{
			Email: msg.Addresses.From.Address,
			Name:  msg.Addresses.From.Name,
		}
	}

	if msg.Addresses.To != nil {
		postmarkMsg.To = msg.Addresses.To.Address
		postmarkMsg.ToFull = []PostmarkAddress{{
			Email: msg.Addresses.To.Address,
			Name:  msg.Addresses.To.Name,
		}}
	}

	if len(msg.Addresses.Cc) > 0 {
		postmarkMsg.Cc = emailAddressesToCommaSeparated(msg.Addresses.Cc)
		postmarkMsg.CcFull = emailAddressesToPostmarkAddresses(msg.Addresses.Cc)
	}

	if len(msg.Addresses.Bcc) > 0 {
		postmarkMsg.Bcc = emailAddressesToCommaSeparated(msg.Addresses.Bcc)
		postmarkMsg.BccFull = emailAddressesToPostmarkAddresses(msg.Addresses.Bcc)
	}

	if len(msg.Addresses.ReplyTo) > 0 {
		postmarkMsg.ReplyTo = emailAddressesToCommaSeparated(msg.Addresses.ReplyTo)
	}

	for _, att := range msg.Attachments {
		contentLength := len(att.Data) * 3 / 4
		postmarkMsg.Attachments = append(postmarkMsg.Attachments, PostmarkAttachment{
			Name:          att.Filename,
			Content:       att.Data,
			ContentType:   att.ContentType,
			ContentLength: contentLength,
		})
	}

	return postmarkMsg
}

// func smtpsrvMesssage2EmailMessage(msg *smtpsrv.Context)
