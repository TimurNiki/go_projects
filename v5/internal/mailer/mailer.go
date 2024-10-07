package mailer

import "embed"

//go:embed "templates"
var FS embed.FS

const (
	FromName            = "GopherSocial"
	maxRetires          = 3
	UserWelcomeTemplate = "user_invitation.tmpl"
)
type Client interface{
	Send(templateFile,username, email string,date any, isSandbox bool) (int,error)
}