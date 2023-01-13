package main

//* see proper project structure from this appication 

/*
CMD -> main.go								-> create server and define location of our router

	-> config 		=> app.go				-> helps us connect with our database
	-> controllers 	=> book_controller.go	-> process data, after recieving and before sending 
PKG -> models		=> book.go				-> contains structs and models to be used by ourdatabase. It also contain operations to be performed with the database to be used by the controller
	-> routes		=> bookstore-routes.go	-> contains routes
	-> utils		=> utils.go				-> marsheling and unmarsheling json data


Controller Functions & Routes
	getbooks		<-		/book/			<-	GET
	createbook		<-		/book/			<-	POST
	get_book_by_id 	<-		/book{bookId}	<- 	GET
	updatebook		<-		/book{bookId}	<-	PUT
	deletebook		<-		/book{bookId}	<-	DELETE

*/
// GORN is a ORM pakage
// create routes first -> then config -> utils -> models -> controllers
// data flow: routes -> controllers -> models

import(
	// Path: 1st part of import is the path of github (FROM ). 
	// Path: 2nd part is the filename
	// Absolute Paths
	// github.com/chowdhuryrahulc/mysql-crud-api-using-json
	"log"
	"net/http"
	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/chowdhuryrahulc/mysql-crud-api-using-json/pkg/routes"
)
// Problem: Our app is creating a table named books in the database-named vocab. But where is the books comming from?
// Sol: model->book.go->db.AutoMigrate(&Book{})

// Problem2: Our name, author, publication names are set as null. Why??
// Sol: solved, no idea how

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)	// function in bookstore-routes.go file
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))	// if err != nil not needed, already added in this line
}