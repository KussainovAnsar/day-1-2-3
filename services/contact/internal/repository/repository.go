package repository

type ContactRepository interface {
	CreateContact() error
	ReadContact() error
	UpdateContact() error
	DeleteContact() error
}

type GroupRepository interface {
	CreateGroup() error
	ReadGroup() error
	AddContactToGroup() error
}
