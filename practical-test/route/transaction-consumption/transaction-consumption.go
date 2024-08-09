package transaction_transactionConsumption

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"technical-test-icon-pln/practical-test/config/database/paginate"
	"technical-test-icon-pln/practical-test/controller"
	"technical-test-icon-pln/practical-test/repository"
	"technical-test-icon-pln/practical-test/service"
)

func TransactionConsumptionRoute(db *gorm.DB, routerGroup *gin.RouterGroup) {
	// Repositories
	transactionConsumptionRepo := repository.NewTransactionConsumptionRepository(db)

	// Services
	transactionConsumptionService := service.NewTransactionConsumptionService(transactionConsumptionRepo)

	//paginate
	pagination := paginate.NewPagination()

	// Controllers
	transactionConsumptionController := controller.NewTransactionConsumptionController(transactionConsumptionService, *pagination)

	// Endpoints
	routerGroup.POST("/transaction-consumption-service/", transactionConsumptionController.Create)
	routerGroup.GET("/transaction-consumption-service/:id", transactionConsumptionController.FindById)
	routerGroup.GET("/transaction-consumption-service/", transactionConsumptionController.FindAll)
	routerGroup.GET("/transaction-consumption-service/without-pagination", transactionConsumptionController.FindAllWithoutPagination)
	routerGroup.DELETE("/transaction-consumption-service/:id", transactionConsumptionController.Delete)
}
