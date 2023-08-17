# Library Management System

The Library Management System is a web application designed to efficiently manage the operations of a library.

## Features

- Secure authentication and session management to ensure the privacy and security of user data.
- Admin Dashboard: Allows administrators to add books, manage book requests, and handle new admin requests.
- User Dashboard: Enables users to search for books, request books, return books, and manage their requests.

## Tech Stack

**Client:** HTML, CSS

**Server:** Golang, MySQL

## Installation

1. Clone the repository: 
`git clone https://github.com/lakshyashishir/Library-Management-System.git`.

2. Run `make all`
 It will ask for configuration details and setup your MySQL database. It will also build, test and run the website at `http://localhost:8000`

## Usage

1. Open your browser and navigate to `http://localhost:8000`.
2. Create an account and log in to the system.
3. Once an admin has signed up, all other admins must request admin access from existing admins.
4. Users and admins have separate dashboards with different functionalities.
5. Users can search for books, request books, and return books. They can also view the status of their requests.
6. Admins can approve or reject book requests from users and handle admin signup requests. They can also view the list of issued books and add new books to the system.

7. You can use the following commands :

`make build` to get dependencies and build  
`make test` to test  
`make run` to run

## Feedback

If you have any feedback or suggestions, please reach out to me at lakshyashishir1@gmail.com. 