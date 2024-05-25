package env
/*

    usage --

    key := GetEnv("ENV_NAME")

*/
import (
    "os"
    "log"

    "github.com/joho/godotenv"
)

var envVars map[string]string

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    envVars = make(map[string]string)

    //TODO: make this more efficient
    keys := []string{
        "AHMED", 
        "HCP_CLIENT_ID",
        "MONGO_CONNECTION_STRING",
        "JWT_SECRET_KEY", 
        /*... add more keys here ...*/}

    for _, key := range keys {
        if value, exists := os.LookupEnv(key); exists {
            envVars[key] = value
        } else {
            log.Printf("Key \"%s\" not found", key)
        }
    }
}

func GetEnv(key string) string {
    if value, exists := envVars[key]; exists {
        return value
    }
    log.Printf("Key \"%s\" not found", key)
    return ""
}