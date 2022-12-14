package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/i3ankjs/chat/pkg/common/db"
	"github.com/i3ankjs/chat/pkg/users"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)
	fmt.Println(port)
	r := gin.Default()
	h := db.Init(dbUrl)

	users.RegisterRoutes(r, h)
	// register more routes here

	http.ListenAndServe(":2222", r)
}
