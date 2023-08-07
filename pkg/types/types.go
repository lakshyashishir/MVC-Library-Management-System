package types

type UserRole string

const (
	Admin          UserRole = "admin"
	Userrole       UserRole = "user"
	AdminRequested UserRole = "admin requested"
)

type BookStatus string

const (
	Available    BookStatus = "available"
	NotAvailable BookStatus = "not available"
	Requested    BookStatus = "requested"
)

type RequestStatus string

const (
	Pending  RequestStatus = "pending"
	Approved RequestStatus = "approved"
	Rejected RequestStatus = "rejected"
)

type User struct {
	UserID   int
	Username string
	Hash     string
	Salt     string
	Role     UserRole
}

type Book struct {
	BookID     int
	Title      string
	Author     string
	BookStatus BookStatus
	Quantity   int
}

type Request struct {
	RequestID  int
	UserID     int
	BookID     int
	BookStatus RequestStatus
}

type Cookie struct {
	ID        int
	SessionID string
	UserID    int
}
