package daos

// Varables to stroe data in-memmory as long as server is running

var (
	UsersDB          IUsers
	DocumentsDB      IDocuments
	DocumentAccessDB IDocumentAccesses
)

func InitDB() {
	UsersDB = NewUsers()
	DocumentsDB = NewDocuments()
	DocumentAccessDB = NewDocumentAccess()
}
