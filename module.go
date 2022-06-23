package main

type ContactList struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Position  string
}

type TaskList struct {
	ID        int
	Name      string
	Status    string
	Priority  int
	CreatedAt string
	CreatedBy string
	DueDate   string
}
