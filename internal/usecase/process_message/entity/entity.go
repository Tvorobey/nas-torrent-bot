package entity

type Command string

type SimpleMessageIn struct {
	UserID int64
	ChatID int64
}

type CommandMessageIn struct {
	UserID  int64
	ChatID  int64
	Command Command
	Args    string
}

type FileMessageIn struct {
	UserID   int64
	FileName string
	Url      string
}
