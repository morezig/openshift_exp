package main

import (
	"openshift_exp/oc-integrator/go-api/models"
	"openshift_exp/oc-integrator/go-api/routers"
)

func main() {
	// err := os.Remove("db.sqlite")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("==> done deleting file")

	models.ConnectDB()
	// models.ConnectRedis()
	// models.CreateMenuTemp()
	// if gin.Mode() == gin.DebugMode {
	// 	models.CreateRoleTestData()
	// 	// models.CreateLayerTestData()
	// 	// models.CreateGameTestData()
	// 	models.CreateGameHallTestData()
	// 	models.CreateCurrencyTestData()
	// 	// models.CreateUserTestData()
	// 	models.CreateUserTestDataFromSql()
	// }

	app := routers.SetupRouter()

	app.Run(":8888")
}
