package main

// import "fmt"

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

func (c ContactList) Create() {

}

func main() {
	x := ContactList{
		1, "Rasul", "Rustamov", "+998991770767", "rustamov.abdurasul@inbox.ru", "QA",
	}

	// fmt.Println(x.Create())
}
