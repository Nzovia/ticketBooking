package utils

import ("log" 
       "github.com/joho/godotenv")

func LoadEnvVariables(){
	err :=godotenv.Load()

	if(err != nil){
		log.Fatal("error while loading the ienv file")
	}
}