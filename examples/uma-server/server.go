package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
	"github.com/uma-universal-money-address/uma-go-sdk/uma"
	"github.com/uma-universal-money-address/uma-go-sdk/uma/errors"
	"github.com/uma-universal-money-address/uma-go-sdk/uma/generated"
	umautils "github.com/uma-universal-money-address/uma-go-sdk/uma/utils"
)

/**
 * This is a simple Gin server (https://gin-gonic.com) that implements the UMA protocol using the Lightspark SDK.
 *
 * By default, this server will run on port 8080, but you can set the PORT environment variable to change that. You can
 * make a request to the API through curl to make sure the server is working properly (replace bob with the username you
 * have configured). If you're running 2 instances of this server locally, one on port 8080 and one on port 8081, you
 * can test the UMA protocol by running commands like:
 *
 * curl -X GET http://localhost:8080/api/umalookup/\$bob@localhost:8081
 * curl -X GET "http://localhost:8080/api/umapayreq/52ca86cd-62ed-4110-9774-4e07b9aa1f0e?amount=100&currencyCode=USD"
 * curl -X POST http://localhost:8080/api/sendpayment/e26cbee9-f09d-4ada-a731-965cbd043d50
 *
 * Configuration parameters (API keys, etc.) and information on how to set them can be found in config.go.
 */

func main() {
	config := NewConfig()
	log.Printf("Starting server with config: %+v", config)
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	store := cookie.NewStore([]byte(config.CookieSecret))
	engine.Use(sessions.Sessions("uma_session", store))
	pubKeyCache := uma.NewInMemoryPublicKeyCache()
	oneDayAgo := time.Now().AddDate(0, 0, -1)
	userService := NewUserServiceFromEnv(config)

	engine.Use(UmaErrorHandler())
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())

	// Require authentication for all API routes.
	engine.Use(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, "/api/") {
			c.Next()
			return
		}
		user, err := userService.GetUserFromContext(c)
		if err != nil {
			c.Error(err)
			return
		}
		c.Set("user_id", user.ID)
	})

	vasp1 := NewVasp1(&config, pubKeyCache, userService)
	vasp2 := Vasp2{
		config:      &config,
		pubKeyCache: pubKeyCache,
		nonceCache:  uma.NewInMemoryNonceCache(oneDayAgo),
	}

	// VASP1 Routes:
	engine.GET("/api/umalookup/:receiver", func(c *gin.Context) {
		vasp1.handleClientUmaLookup(c)
	})

	engine.GET("/api/umapayreq/:callbackUuid", func(c *gin.Context) {
		vasp1.handleClientPayReq(c)
	})

	engine.POST("/api/sendpayment/:callbackUuid", func(c *gin.Context) {
		vasp1.handleClientPaymentConfirm(c)
	})

	engine.POST("/api/uma/pay_invoice", func(c *gin.Context) {
		vasp1.handlePayInvoice(c)
	})

	engine.POST("/uma/request_invoice_payment", func(c *gin.Context) {
		vasp1.handleRequestPayInvoice(c)
	})

	// End VASP1 Routes

	// VASP2 Routes:
	engine.GET("/.well-known/lnurlp/:username", func(c *gin.Context) {
		vasp2.handleWellKnownLnurlp(c)
	})

	engine.GET("/uma/payreq/:uuid", func(c *gin.Context) {
		vasp2.handleLnurlPayreq(c)
	})

	engine.POST("/uma/payreq/:uuid", func(c *gin.Context) {
		vasp2.handleUmaPayreq(c)
	})

	engine.POST("/api/uma/create_invoice/:uuid", func(c *gin.Context) {
		vasp2.handleCreateInvoice(c)
	})

	engine.POST("/api/uma/create_and_send_invoice/:uuid", func(c *gin.Context) {
		vasp2.handleCreateAndSendInvoice(c)
	})

	// End VASP2 Routes

	// Shared:

	engine.GET("/.well-known/lnurlpubkey", func(c *gin.Context) {
		// It doesn't matter which vasp protocol handles this since they share a config and cache.
		vasp2.handlePubKeyRequest(c)
	})

	engine.GET("/uma/utxocallback", func(c *gin.Context) {
		// It doesn't matter which vasp protocol handles this since they share a config and cache.
		vasp2.handleUtxoCallback(c)
	})

	engine.POST("/.well-known/uma-configuration", func(c *gin.Context) {
		scheme := "https"
		if umautils.IsDomainLocalhost(c.Request.Host) {
			scheme = "http"
		}
		c.JSON(http.StatusOK, gin.H{
			"uma_major_versions":   uma.GetSupportedMajorVersions(),
			"uma_request_endpoint": fmt.Sprintf("%s://%s/uma/request_invoice_payment", scheme, c.Request.Host),
		})
	})

	// Frontend:

	engine.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	engine.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		hashedPassword := hashString(password)
		if username == config.Username && hashedPassword == config.HashedUserPassword {
			session := sessions.Default(c)
			session.Set("user_id", config.UserID)
			session.Save()
			c.Redirect(http.StatusFound, "/")
			return
		}
		c.HTML(http.StatusOK, "login.html", gin.H{
			"error":    "Invalid username or password",
			"username": username,
			"password": password,
		})
	})

	engine.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.Redirect(http.StatusFound, "/login")
	})

	engine.GET("/", func(c *gin.Context) {
		user, err := userService.GetUserFromContext(c)
		if err != nil || user == nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{
			"username": user.Name,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	engine.Run(":" + port)
}

func UmaErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case *errors.UmaError:
				c.AbortWithStatusJSON(e.ToHttpStatusCode(), umaErrorToMap(e))
				return
			default:
				internalError := &errors.UmaError{
					Reason:    c.Errors.Last().Error(),
					ErrorCode: generated.InternalError,
				}
				c.AbortWithStatusJSON(internalError.ToHttpStatusCode(), umaErrorToMap(internalError))
				return
			}
		}
	}
}

func umaErrorToMap(umaError *errors.UmaError) map[string]interface{} {
	jsonStr, marshalErr := umaError.ToJSON()
	if marshalErr != nil {
		return map[string]interface{}{
			"status": "ERROR",
			"reason": "Failed to serialize UMA error",
			"code":   generated.InternalError.Code,
		}
	} else {
		var jsonMap map[string]interface{}
		if err := json.Unmarshal([]byte(jsonStr), &jsonMap); err != nil {
			return map[string]interface{}{
				"status": "ERROR",
				"reason": "Failed to parse error JSON",
				"code":   generated.InternalError.Code,
			}
		} else {
			return jsonMap
		}
	}
}
