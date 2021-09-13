# go-crud-api
An API for borrow books at the library.  

## Features
There are the following endpoints
 
- **User**
 
  `POST /register` Register a user
  
  `POST /login` Login a user
  
- **Book**

  `GET /book` Show all books
  
  `POST /book` Input new book
  
  `GET /book/:id` Show specify book
  
  `PATCH /book/:id` Update specify book
  
  `DELETE /book/:id` Delete specify book
  
## Getting Started
Before you begin, don't forget to clone this repo first and put it locally wherever you want.

### Prerequisites
Things you need to install so that you can run this project is

- Go (https://golang.org/dl/)

- PostgreSQL (https://www.postgresql.org/download/)

- An .database.env file, it such a configuration thing.

### Installing
Open terminal in the root of this project where you saved it. And do the following command.

```shell
go build
```

That command will build the project so you can run it after that. 
If it's already built then you just run this command. 

```shell
go run main.go
```

Make sure you already set up your database on your own, if not? You can go following these [tutorial](https://www.postgresqltutorial.com/). 

## License
developed with 💕 by **Ridho Saputra** - *[Instagram](https://instagram.com/mridhosap)*
