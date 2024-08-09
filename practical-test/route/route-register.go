package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"technical-test-icon-pln/practical-test/route/consumption"
	"technical-test-icon-pln/practical-test/route/transaction"
	transactionConsumption "technical-test-icon-pln/practical-test/route/transaction-consumption"
)

func RouteRegister(db *gorm.DB, routerGroup *gin.RouterGroup) {
	consumption.ConsumptionRoute(db, routerGroup)
	transaction.TransactionRoute(db, routerGroup)
	transactionConsumption.TransactionConsumptionRoute(db, routerGroup)
}
