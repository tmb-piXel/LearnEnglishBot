package models

type User struct {
	chatID     int64
	language   string
	topic      string
	isToRu     bool
	original   *[]string
	translated *[]string
	iterWord   int
}

func NewUser(chatID int64) *User {
	u := &User{
		chatID: chatID,
		isToRu: true,
	}
	return u
}

func (u *User) GetChatID() int64   { return u.chatID }
func (u *User) SetChatID(id int64) { u.chatID = id }

func (u *User) GetLanguage() string  { return u.language }
func (u *User) SetLanguage(l string) { u.language = l }

func (u *User) GetTopic() string  { return u.topic }
func (u *User) SetTopic(t string) { u.topic = t }

func (u *User) GetIsToRu() bool       { return u.isToRu }
func (u *User) SetIsToRu(isToRu bool) { u.isToRu = isToRu }

func (u *User) GetOriginal() *[]string  { return u.original }
func (u *User) SetOriginal(o *[]string) { u.original = o }

func (u *User) GetTransleted() *[]string  { return u.translated }
func (u *User) SetTransleted(t *[]string) { u.translated = t }

func (u *User) GetIterWord() int  { return u.iterWord }
func (u *User) SetIterWord(i int) { u.iterWord = i }
