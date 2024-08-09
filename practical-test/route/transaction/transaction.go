package transaction

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"technical-test-icon-pln/practical-test/config/database/paginate"
	"technical-test-icon-pln/practical-test/controller"
	"technical-test-icon-pln/practical-test/repository"
	"technical-test-icon-pln/practical-test/service"
)

func TransactionRoute(db *gorm.DB, routerGroup *gin.RouterGroup) {
	// Repositories
	transactionRepo := repository.NewTransactionRepository(db)

	// Services
	transactionService := service.NewTransactionService(transactionRepo)

	//paginate
	pagination := paginate.NewPagination()

	// Controllers
	transactionController := controller.NewTransactionController(transactionService, *pagination)

	// Endpoints
	routerGroup.POST("/transaction-service/", transactionController.Create)
	routerGroup.GET("/transaction-service/:id", transactionController.FindById)
	routerGroup.GET("/transaction-service/", transactionController.FindAll)
	routerGroup.GET("/transaction-service/without-pagination", transactionController.FindAllWithoutPagination)
	routerGroup.DELETE("/transaction-service/:id", transactionController.Delete)
}
