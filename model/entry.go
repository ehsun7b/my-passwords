package model

// Entry password entry model class
type Entry struct {
	title    string
	username string
	password string
	url      string
	tags     []Tag
}

// getter methods

//GetTitle return title field of entry
func (e *Entry) GetTitle() string {
	return e.title
}

//GetUsername return username field of entry
func (e *Entry) GetUsername() string {
	return e.username
}

//GetPassword return password field of entry
func (e *Entry) GetPassword() string {
	return e.password
}

//GetURL return url field of entry
func (e *Entry) GetURL() string {
	return e.url
}

//GetTags return tags field of entry
func (e *Entry) GetTags() []Tag {
	return e.tags
}

// setter methids

//SetTitle set title field of entry
func (e *Entry) SetTitle(title string) {
	e.title = title
}

//SetUsername set username field of entry
func (e *Entry) SetUsername(username string) {
	e.username = username
}

//SetPassword set password field of entry
func (e *Entry) SetPassword(password string) {
	e.password = password
}

//SetURL set url field of entry
func (e *Entry) SetURL(url string) {
	e.url = url
}

//SetTags set tags field of entry
func (e *Entry) SetTags(tags []Tag) {
	e.tags = tags
}
