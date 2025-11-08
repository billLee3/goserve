# GoServe
---
### Description
---
GoServe is a web server built from scratch using Golang. Primarily focusing on the implementation of the net/http package.
The application can perform crud operations on users and chirps (similar to tweets). Users are authenticated via JWTs and the JWTs are validated for access tokens to the database operations on resources they created. Refresh tokens are also implemented in the application for long lasting revocable authentication.

### Motivation
---
My motivation for this project was to get really adept at the fundamentals of web servers before going to use any form of framework like Gin for Golang or something else. I've used other frameworks in different languages like ASP.NET and FAST API but always felt like I was lacking some fundamental knowledge regarding how the DOM, API, and JWTs worked together. This project did a great job of soaring up the fundamentals I was weak in while reinforcing my understanding on other aspects of the HTTP request and response model.

### Quick Start
---
#### Setup Required to Run GoServe on your machine
(Linux) Install postgres on your local machine using the following code:
```
sudo apt update
sudo apt install postgresql postgresql-contrib
```
If on linux then you will need to update your password, likeso: sudo passwd postgres (Linux) Install Golang on your machine using the following code:
```
sudo apt install golang-go
```
#### Install GoServe
---
Now that you have Golang install on your machine use the following golang command from the terminal
`go install https://github.com/billLee3/goserve@latest

#### Create a .env file
---
Create a .env file with the following variables:
- DB_URL="<postgres connection string>"
- PLATFORM=["dev"]
- SECRET="<user generated secret>"
- POLKA_KEY=["f271c81ff7084ee5b99a5091b42d486e"]

Platform is required to use dev in order to work effectively. Polka Key can really be any string but the string above is suggested. The secret variable can be any string the user wants to use. 

#### Migrate the DB
---
From the goserve/sql/schema directory run: `goose postgres [connection string] up`. This will get the database ready for CRUD operations.

### Usage
---
The primary way to use the application is through a client like Postman or through curl requests. 
To run the application run the following command from the root folder goserve/:
`go build -o out && ./out`
This command builds the executable and then runs the executable.

Below are the endpoints available to the user:
- POST /api/users -> allows the user to create a new user. Supplying the request with a json structure like so: 
`{"email": "testemail@gmail.com", "password": "****"}`
- POST /api/login -> logins in the user and creates a JWT for authentication. The request body would be structured as:
`{"email": "testemail@gmail.com", "password": "****"}`
- PUT /api/users -> allows the user to update their password. The handler behind the endpoint validates using the JWT to authenticate the user and uses that authentication to update the user's password from the request body following this structure: 
`{"email": "testemail@gmail.com", "password": "****"}`
- POST /api/chirps -> allows the user to create a chirp (essentially a tweet) using the following resquest body structure: 
`{"id": "uuid", "created_at": timestampValue, "updated_at": timestampValue, "user_id": "uuid", "body": "text of the chirp"}`
- POST /api/chirps -> allows the user to create a chirp (essentially a tweet) using the following resquest body structure: 
`{"id": "uuid", "created_at": timestampValue, "updated_at": timestampValue, "user_id": "uuid", "body": "text of the chirp"}`
- GET /api/chirps -> gets all of the chirps from the database and filters them based on url query parameters. Below are the query options available:
    - sort: asc or desc
    - author_id: uuid value of the author's uuid.
- GET /api/chirps/{chirp_id} -> gets a single chirp based on the UUID passed into the url. 

There are additional endpoints to hit but those are the main pieces of functionality for a given end user to plug in. The remaining endpoints deal with admin metrics and resets, as well as refresh token management and webhook execution validation. 

## Contributing
If you'd like to contribute fork the repository and open a pull request to the main branch.

