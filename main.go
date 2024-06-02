package main

import (
	"books/config"
	"books/controller"
	"books/middleware"
	"books/repository"
	"books/service"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()

	bookRepository repository.BookRepository = repository.NewBookRepository(db)

	BookService service.BookService = service.NewBookService(bookRepository)

	bookController controller.BookController = controller.NewBookService(BookService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	booksRoutes := r.Group("api/v1")
	{
		booksRoutes.GET("/books", bookController.All)
		booksRoutes.GET("/books/:id", bookController.FindByID)
		booksRoutes.POST("/books/create", bookController.Create)
		booksRoutes.PUT("/books/update/:id", bookController.Update)
		booksRoutes.DELETE("/books/delete/:id", bookController.Delete)
	}
	r.Run(os.Getenv("API_PORT"))
}
