package mailer

import (
	"context"
	"fmt"
	"net/smtp"
)

type Mailer struct {
	auth smtp.Auth

	from string
	addr string
}

type MailerOpts struct {
	User     string
	Password string

	From string

	Host string
	Addr string
}

func NewMailer(opts MailerOpts) *Mailer {
	auth := smtp.PlainAuth("", opts.User, opts.Password, opts.Host)

	return &Mailer{
		auth: auth,
		from: opts.From,
		addr: opts.Addr,
	}
}

func (m Mailer) Send(ctx context.Context, to string, subject string, msg string) error {
	errChan := make(chan error)

	go func() {
		processedMsg := fmt.Sprintf("From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n\r\n"+
			"%s", m.from, to, subject, msg)

		errChan <- smtp.SendMail(m.addr, m.auth, m.from, []string{to}, []byte(processedMsg))
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return err
		}

		close(errChan)

		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
