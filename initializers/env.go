package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func Loadenv() {
  err := godotenv.Load(".env")
   if err != nil {
     log.Println(err, "error while loading file from env file ")
   }
}