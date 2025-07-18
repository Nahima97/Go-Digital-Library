# Welcome to Go Digital's Digital Library!

# How to Navigate The Digital Library

KEY:

API Endpoints
Method	   Endpoint	    Description
POST	   /register	Register a user
POST	   /login	    Authenticate + get JWT Token

Books (protected routes)
Method	     Endpoint	             Description
POST	     /books	                 Add a new book
GET	         /books/borrow/{id}	     Borrow a book by ID
PUT	         /books/return/{id}	     Return a book
DELETE	     /books/delete/{id}	     Delete a book



# Deployment Instructions
User Routes:
In order to Register a user please follow the below link:
http/localhost:8080/register

For user login follow this link: 
http/localhost:8080/login

Book Routes:
TO Search for books within the library follow this:
http/localhost:8080/books

To Add a book follow the below:
http/localhost:8080/books/add

To Borrow a book follow the below:
http/localhost:8080/books/borrow/{id}

To Return a book follow the below:
http/localhost:8080/books/return/{id}

To Delete a book follow the below:
http/localhost:8080/books/delete/{id}


# Thank you Oreva
To Oreva:
From all of us, Thank you for guiding us through the world of Go. What felt daunting at first quickly became more manageable, thanks to your support.

You turned our panic() into fmt.Println("I got this!")
and transformed our "nil" knowledge into real skills.

We couldn’t have asked for a better teacher —
we appreciate you more than a perfectly bug-free assignment! 😄

- From Team Go Digital - 

