package app

import (
	"log"
	"os"
	"time"

	"github.com/golangcollege/sessions"
	"gorm.io/gorm"
	"undrakh.net/summarizer/pkg/common/mailer"
	"undrakh.net/summarizer/pkg/easyOAuth2"
	"undrakh.net/summarizer/pkg/userman"
)

var (
	// Defaults
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	DB       *gorm.DB
	Config   = conf{}
	Location *time.Location
	Session  *sessions.Session

	// Services
	Mailer       *mailer.Mailer
	Users        *userman.Service
	GoogleOAuth2 *easyOAuth2.EasyOAuthClient
)

func Init() {
	InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// path string
	// loc, err := time.LoadLocation("Asia/Ulaanbaatar")
	// if err != nil {
	// 	panic(err)
	// }
	// Location = loc

	// // apputils.LoadConfig(&Config, path)

	// GoogleOAuth2 = &easyOAuth2.EasyOAuthClient{
	// 	Name: "google",
	// 	Config: &oauth2.Config{
	// 		RedirectURL:  Config.OAuth2.Google.RedirectURL,
	// 		ClientID:     Config.OAuth2.Google.ClientID,
	// 		ClientSecret: Config.OAuth2.Google.ClientSecret,
	// 		Scopes:       Config.OAuth2.Google.Scopes,
	// 		Endpoint:     google.Endpoint,
	// 	},
	// 	UserInfoEndpoint: Config.OAuth2.Google.UserInfoEndpoint,
	// }

	// DB = apputils.OpenDB(Config.DSN)

	// Users = userman.NewService(DB, InfoLog, ErrorLog)

	// Session = sessions.New([]byte(Config.SessionSecret))
	// Session.Lifetime = 7 * 24 * time.Hour
	// Session.Secure = true
	// Session.HttpOnly = false
	// Session.Path = "/"
}

func Close() {
}
