package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

//"github.com/gofiber/fiber/v2"
//"gorm.io/gorm"
//"github.com/joho/godotenv"

type Book struct{   //in golang you have to tell how the json will look like 
	Author     string    `json:"author"`
	Title      string    `json:"title"`
	Publisher   string    `json:"publisher"`
}

type Repository struct{
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *fiber.Ctx)  error{
	book := Book()

	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message":"request failed"})
			return err
	}

	r.DB.Create(&book).Error
	if err != nil{
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message":"could not create book"})
			return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"book has been added"})
	return nil
}

func (r *repository)DeleteBook(context, *fiber.Ctx) error{
	bookModel := models.Books{}
	id := context.Params("id")
	if id == ""{
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message":"id cannnot be empty"
		}),
		return nil
	}

	err := r.DB.Delete(bookModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message":"could not delete book"
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"book deleted successfully"
	})
	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error{
	bookModels := &[]models.Books{}
	
	err := r.DB.Find(booksModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message":"could not get the books"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"books fetched successfull",
		"data": bookModels,
	})
	return nil
}


func (r  *Repository) GetBookByID(context *fiber.Ctx) error {
	id := context.Params("id")
	bookModel := &models.Books{}
	if id =""{
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message":"id cannot be empty"
		})
		return nil
	}

	fmt.Println("the ID is", id)

	err := r.DB.Where("id = ?", id).First(bookModel).Error
	if err != nil{
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message":"could not get the book"
		})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"book id fetched succsefully",
		"data": bookModel
	})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App){
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_books?:id", r.GetBookByID)
	pai.Get("/books", r.GetBooks)
}

func main()  {
	err := godotenv.Load(".env")

	if err != nil{
		log.Fatal(err)
	}
	config := &storage.Config{
		Host: os.Getenv("DB_HOST"),
		Port:  os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:    os.Getenv("DB_USER"),
		SSLMode: os.Getenv("DB_SSLMODE"),
		DBName:  os.Getenv("DB_DBNAME"),

	}

	db, err := storage.NewConnection(config) //this error

	if err != nil {
		log.Fatal("Could not load database")
	}
	err = models.MigrateBooks(db)
	if err := nil {
		log.Fatal("could not migrate")
	}

	r := Repository(
		DB: db,
	)

	app := fiber.New(type)
	r.SetupRoutes(app)   //r is a repository defined above on line 7
	app.Listen(":8080")
}