package nbi5

import (
	"database/sql"
	"html/template"
	"log/slog"
	"net/http"
)

// Is this enough for modules to reuse?
type Notebrew struct {
	DB        *sql.DB
	Dialect   string
	ErrorCode func(error) string
	Domain    string
}

// User represents a user in the users table.
type User struct {
	// UserID uniquely identifies a user. It cannot be changed.
	UserID ID `json:"userID"`

	// Username uniquely identifies a user. It can be changed.
	Username string `json:"username"`

	// Email uniquely identifies a user. It can be changed.
	Email string `json:"email"`

	// TimezoneOffsetSeconds represents a user's preferred timezone offset in
	// seconds.
	TimezoneOffsetSeconds int `json:"timezoneOffsetSeconds"`

	// Is not empty, DisableReason is the reason why the user's account is
	// marked as disabled.
	DisableReason string `json:"disableReason"`

	// SiteLimit is the limit on the number of sites the user can create.
	SiteLimit int64 `json:"siteLimit"`

	// StorageLimit is the limit on the amount of storage the user can use.
	StorageLimit int64 `json:"storageLimit"`

	// UserFlags are various properties on a user that may be enabled or
	// disabled e.g. UploadImages.
	UserFlags map[string]bool `json:"userFlags"`
}

type ContextData struct {
	Domain string `json:"domain"`

	URLPath string `json:"urlPath"`

	UserID ID `json:"userID"`

	Username string `json:"username"`

	DisableReason string `json:"disableReason"`

	UserFlags map[string]bool `json:"userFlags"`

	Referer string `json:"-"`

	PathTail string `json:"-"`

	DevMode bool `json:"-"`

	CSS template.CSS `json:"-"`

	JS template.JS `json:"-"`

	Logger *slog.Logger `json:"-"`

	ErrorHandler func(w http.ResponseWriter, r *http.Request, statusCode int, err error) `json:"-"`
}

func (nbrew *Notebrew) GetUser(r *http.Request) (User, error) {
	return User{}, nil
}

func (nbrew *Notebrew) ExecuteTemplate(w http.ResponseWriter, r *http.Request, tmpl *template.Template, data any) {
}

func (nbrew *Notebrew) GetSession() {
}

func (nbrew *Notebrew) SetSession() {
}
