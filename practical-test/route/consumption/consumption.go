package consumption

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"technical-test-icon-pln/practical-test/config/database/paginate"
	"technical-test-icon-pln/practical-test/controller"
	"technical-test-icon-pln/practical-test/repository"
	"technical-test-icon-pln/practical-test/service"
)

func ConsumptionRoute(db *gorm.DB, routerGroup *gin.RouterGroup) {
	// Repositories
	consumptionRepo := repository.NewConsumptionRepository(db)

	// Services
	consumptionService := service.NewConsumptionService(consumptionRepo)

	//paginate
	pagination := paginate.NewPagination()

	// Controllers
	consumptionController := controller.NewConsumptionController(consumptionService, *pagination)

	// Endpoints
	routerGroup.POST("/consumption-service/", consumptionController.Create)
	routerGroup.GET("/consumption-service/:id", consumptionController.FindById)
	routerGroup.GET("/consumption-service/", consumptionController.FindAll)
	routerGroup.GET("/consumption-service/without-pagination", consumptionController.FindAllWithoutPagination)
	routerGroup.DELETE("/consumption-service/:id", consumptionController.Delete)
}
