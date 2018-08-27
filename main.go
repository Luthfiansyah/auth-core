package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/auth-core/app/controllers"
	"github.com/auth-core/config"
	"github.com/auth-core/database"
	"github.com/auth-core/logging"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var log = logging.MustGetLogger("auth-core")
var client *redis.Client
var router *gin.Engine

func main() {

	configFile := flag.String("c", "", "configuration file")
	flag.Parse()

	if *configFile == "" {
		fmt.Println("\n\nUse -h to get more information on command line options")
		fmt.Println("You must specify a configuration file")
		os.Exit(1)
	}

	err := config.Initialize(*configFile)
	if err != nil {
		fmt.Printf("Error reading configuration: %s\n", err.Error())
		os.Exit(1)
	}

	// CALL DB CONN AND MIGRATION
	err = database.InitDB()
	if err != nil {
		fmt.Printf("Error reading configuration: %s\n", err.Error())
		// os.Exit(1)
	}

	// INIT KEY
	// err = InitKey()
	// if err != nil {
	// 	fmt.Printf("Erorr initiating key: %s\n", err.Error())
	// 	os.Exit(1)
	// }

	// nrConfig := newrelic.NewConfig("Zatrip Core", config.MustGetString(controllers.GetRunMode()+".newrelic_key"))
	// app, err := newrelic.NewApplication(nrConfig)

	// controllers.InitializeEndpoints(controllers.NewBackend(config.MustGetString("server.mode")))

	// SetupLogging()
	log.Infof(logging.INTERNAL, "Running in mode %s", config.MustGetString("server.mode"))
	log.Infof(logging.INTERNAL, "Running on port %d", config.MustGetInt("server.port"))
	log.Infof(logging.INTERNAL, "Serving document directory %s", config.MustGetString("server.docs"))
	log.Infof(logging.INTERNAL, "Debug logging: %t", config.MustGetBool("server.debug"))

	// Must be initialized before backends
	// err = services.InitializeMaps(config.MustGetString("server.maps"))
	// if err != nil {
	// 	log.Fatalf(logging.INTERNAL, "error reading map_data.json: %s", err.Error())
	// 	return
	// }

	// docSvc := services.NewDocumentService(config.MustGetString("server.docs"))
	// err = docSvc.ReadAll()
	// if err != nil {
	// 	log.Fatalf(logging.INTERNAL, "error reading docs: %s", err.Error())
	// 	return
	// }

	// endpoints.InitializeEndpoints(
	// 	backends.NewBackend(config.MustGetString("server.mode")),
	// 	docSvc)

	// router := NewRouter(config.MustGetBool("server.debug"), config.MustGetString("server.mode"))
	router = gin.Default()
	initRoutes(config.MustGetBool("server.debug"), config.MustGetString("server.mode"))

	port := os.Getenv("PORT")
	if port == "" {
		port = config.MustGetString("Server.port")
	}

	// EXAMPLE USE GET SINGATURE
	// signature, _ := controllers.GetSignature("test")
	// fmt.Println(signature)

	// TestRedisClient()

	router.Run(":" + port)

}

// func InitKey() error {
// 	key, ok := os.LookupEnv("PP_AUTH_CRED")
// 	if !ok {
// 		return fmt.Errorf("PP_AUTH_CRED is not set in system environment")
// 	}
//
// 	timeout, err := strconv.Atoi(config.MustGetString("server.jwt_timeout"))
// 	if err != nil {
// 		return fmt.Errorf("timeout setting is invalid")
// 	}
//
// 	err = crypto.InitKey([]byte(key), timeout)
// 	return err
// }

// func SetupLogging() {
// 	format := opLogger.MustStringFormatter(
// 		`%{color} %{time:2006-01-02T15:04:05.999Z07:00} %{shortfile:20.20s} %{shortfunc:10.10s} ▶ %{level:.4s} %{message}%{color:reset}`,
// 	)
// 	backend := opLogger.NewLogBackend(os.Stderr, "", 0)
// 	formatter := opLogger.NewBackendFormatter(backend, format)
// 	opLogger.SetBackend(formatter)
// 	log.Info(logging.INTERNAL, "logging initialized")

// }

func Ping(c *gin.Context) {

	res := map[string]string{
		"start_time": controllers.GetCurrentTimeTimeZone("Asia/Jakarta"),
		"message":    "Auth Core Run on " + controllers.GetRunMode() + " mode",
	}

	c.JSON(http.StatusOK, res)
}

func TestRedisClient() {
	client := database.RedisOpen()

	pong, err := client.Ping().Result()
	fmt.Println(pong, err) // Output: PONG <nil>

	// SET KEY
	err = client.Set("auth-core-run-mode", "Auth Core Run on production mode", 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	// GET KEY
	data, _ := client.Get("auth-core-run-mode").Result()
	fmt.Println(data)
}
