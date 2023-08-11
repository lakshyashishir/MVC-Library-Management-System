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
	UserID   int      `json:"user_id"`
	Username string   `json:"username"`
	Hash     string   `json:"hash"`
	Salt     string   `json:"salt"`
	Role     UserRole `json:"role"`
}

type Book struct {
	BookID     int        `json:"book_id"`
	Title      string     `json:"title"`
	Author     string     `json:"author"`
	BookStatus BookStatus `json:"book_status"`
	Quantity   int        `json:"quantity"`
}

type BookUserView struct {
	Title      string        `json:"title"`
	BookStatus RequestStatus `json:"book_status"`
	RequestID  int           `json:"request_id"`
	UserID     int           `json:"user_id"`
	BookID     int           `json:"book_id"`
}

type Request struct {
	RequestID  int           `json:"request_id"`
	UserID     int           `json:"user_id"`
	BookID     int           `json:"book_id"`
	BookStatus RequestStatus `json:"book_status"`
}

type RequestAlt struct {
	RequestID  int           `json:"request_id"`
	UserID     int           `json:"user_id"`
	BookID     int           `json:"book_id"`
	BookStatus RequestStatus `json:"book_status"`
	Username   string        `json:"username"`
	Title      string        `json:"title"`
}

type RequestAdminView struct {
	RequestID     int           `json:"request_id"`
	UserID        int           `json:"user_id"`
	BookID        int           `json:"book_id"`
	BookStatus    BookStatus    `json:"book_status"`
	RequestStatus RequestStatus `json:"request_status"`
	Username      string        `json:"username"`
	Title         string        `json:"title"`
	Author        string        `json:"author"`
}

type Cookie struct {
	ID        int    `json:"id"`
	SessionID string `json:"session_id"`
	UserID    int    `json:"user_id"`
}
