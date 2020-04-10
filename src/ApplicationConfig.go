package src

import (
	"./application"
	"./infrastructure/persistence"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hoanganf/pos_domain/entity"
	"github.com/hoanganf/pos_domain/service"
	"gopkg.in/gorp.v1"
	"os"
)

type Bean struct {
	DbMap             *gorp.DbMap
	RestaurantService *application.RestaurantService
	IngredientService *application.IngredientService
	UserService       *application.UserService
}

func (bean *Bean) DestroyBean() {
	bean.DbMap.Db.Close()
}

func InitBean() (*Bean, error) {
	user := getEnvWithDefault("DB_USER", "root")
	password := getEnvWithDefault("DB_PASSWORD", "")
	//	host := getEnvWithDefault("DB_HOST", "127.0.0.1")
	//	port := getEnvWithDefault("DB_PORT", "3306")
	dbName := getEnvWithDefault("DB_NAME", "anit_pos_server_new")
	//	dsn := fmt.Sprintf("%s:%s@unix(%s:%s)/%s?parseTime=true", user, password, host, port,dbName)
	dsn := fmt.Sprintf("%s:%s@unix(/Applications/XAMPP/xamppfiles/var/mysql/mysql.sock)/%s?parseTime=true", user, password, dbName)
	fmt.Printf("dns: %s", dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	restaurantRepository := persistence.NewRestaurantRepository(dbMap)
	ingredientRepository := persistence.NewIngredientRepository(dbMap)
	productRepository := persistence.NewProductRepository(dbMap)
	userRepository := persistence.NewUserRepository(dbMap)

	restaurantFactory := entity.NewRestaurantFactory()
	ingredientFactory := entity.NewIngredientFactory()
	productFactory := entity.NewProductFactory()
	userFactory := entity.NewUserFactory()

	restaurantService := application.NewRestaurantService(
		restaurantRepository,
		productRepository,
		restaurantFactory,
		productFactory)

	ingredientService := application.NewIngredientService(
		ingredientRepository,
		ingredientFactory)

	userService := application.NewUserService(service.NewUserService(
		userRepository),
		userFactory)

	return &Bean{DbMap: dbMap,
		RestaurantService: restaurantService,
		IngredientService: ingredientService,
		UserService:       userService}, nil
}

func getEnvWithDefault(name, def string) string {
	env := os.Getenv(name)
	if len(env) != 0 {
		return env
	}
	return def
}
