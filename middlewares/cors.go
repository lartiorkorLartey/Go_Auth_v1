package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func CORSMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//         c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
//         c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
//         c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

//         if c.Request.Method == "OPTIONS" {
//             c.AbortWithStatus(http.StatusNoContent)
//             return
//         }

//         c.Next()
//     }
// }

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        allowedOrigins := []string{
            "http://localhost:4200",            // Default Angular development URL
            "https://gatekeeperpro.netlify.app", // Netlify app URL
            "http://localhost:5173",            // Default Vite React development URL
        }

        origin := c.Request.Header.Get("Origin")

        var isAllowedOrigin bool
        for _, o := range allowedOrigins {
            if origin == o {
                isAllowedOrigin = true
                break
            }
        }

        if isAllowedOrigin {
            c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
        } else {
            c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        }

        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-APN")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(http.StatusNoContent)
            return
        }

        c.Next()
    }
}

