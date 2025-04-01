package main

import (
	"encoding/json"
	"fmt"
	"go-control/dto"
	"go-control/helper"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Config struct for JSON configuration
type Config struct {
	Sites []Site `json:"sites"`
}

// Site struct for each site in the configuration
type Site struct {
	SiteTitle     string `json:"site_title"`
	SiteFolder    string `json:"site_folder"`
	SiteCommand   string `json:"site_command"`
	GitBranch     string `json:"git_branch"`
	GitRepository string `json:"git_repository"`
	GitToken      string `json:"git_token"`
	SiteClone     string `json:"site_clone"`
	SiteDeploy    string `json:"site_deploy"`
	Type          string `json:"type"`
}

func loadConfig(configFile string) (*Config, error) {
	// Open the configuration file
	file, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	// Decode the JSON configuration
	config := &Config{}
	if err := json.NewDecoder(file).Decode(config); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}
	return config, nil
}

func main() {
	// Load the configuration
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Ensure sites are configured
	if len(config.Sites) == 0 {
		log.Fatal("No sites configured in the config.json file")
	}

	// Create a new Gin router
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "go control is running",
		})
	})

	// Create dynamic routes for each site
	for _, site := range config.Sites {
		siteTitle := strings.ToLower(site.SiteTitle)
		siteFolder := site.SiteFolder
		siteCommand := site.SiteCommand

		router.POST(fmt.Sprintf("/api/%s/npm-install", siteTitle), func(c *gin.Context) {
			c.Header("Content-Type", "text/event-stream")
			c.Header("Cache-Control", "no-cache")
			c.Header("Connection", "keep-alive")

			cmd := fmt.Sprintf("npm install")
			helper.RunCommandWithStream(c, cmd, siteFolder, fmt.Sprintf("NPM Install %s", siteTitle))
		})

		router.POST(fmt.Sprintf("/api/%s/build", siteTitle), func(c *gin.Context) {
			c.Header("Content-Type", "text/event-stream")
			c.Header("Cache-Control", "no-cache")
			c.Header("Connection", "keep-alive")

			helper.RunCommandWithStream(c, siteCommand, siteFolder, fmt.Sprintf("Build %s", siteTitle))
		})

		router.POST(fmt.Sprintf("/api/%s/pull", siteTitle), func(c *gin.Context) {
			c.Header("Content-Type", "text/event-stream")
			c.Header("Cache-Control", "no-cache")
			c.Header("Connection", "keep-alive")

			cmd := fmt.Sprintf("git pull %s %s", site.GitToken, site.GitBranch)
			helper.RunCommandWithStream(c, cmd, siteFolder, fmt.Sprintf("Pull %s", siteTitle))
		})

		router.POST(fmt.Sprintf("/api/%s/stop-service", siteTitle), func(c *gin.Context) {
			c.Header("Content-Type", "text/event-stream")
			c.Header("Cache-Control", "no-cache")
			c.Header("Connection", "keep-alive")

			cmd := fmt.Sprintf(`appcmd stop site /site.name:%s`, siteTitle)
			helper.RunCommandWithStream(c, cmd, siteFolder, fmt.Sprintf("Deploy %s", siteTitle))
		})

		router.POST(fmt.Sprintf("/api/%s/start-service", siteTitle), func(c *gin.Context) {
			c.Header("Content-Type", "text/event-stream")
			c.Header("Cache-Control", "no-cache")
			c.Header("Connection", "keep-alive")

			cmd := fmt.Sprintf(`appcmd start site /site.name:%s`, siteTitle)
			helper.RunCommandWithStream(c, cmd, siteFolder, fmt.Sprintf("Deploy %s", siteTitle))
		})

		router.POST(fmt.Sprintf("/api/%s/deploy", siteTitle), func(c *gin.Context) {
			c.Header("Content-Type", "text/event-stream")
			c.Header("Cache-Control", "no-cache")
			c.Header("Connection", "keep-alive")

			cmd := fmt.Sprintf("xcopy /s /y %s %s", site.SiteClone, site.SiteDeploy)
			helper.RunCommandWithStream(c, cmd, siteFolder, fmt.Sprintf("Deploy %s", siteTitle))
		})

		router.POST(fmt.Sprintf("/api/%s/auto-deploy", siteTitle), func(c *gin.Context) {
			c.Header("Content-Type", "text/event-stream")
			c.Header("Cache-Control", "no-cache")
			c.Header("Connection", "keep-alive")

			removeFilesCommand := ""
			if site.Type == ".net" {
				removeFilesCommand = fmt.Sprintf(`powershell -Command "Start-Sleep -s 2; Remove-Item -Recurse -Force '%s\*'"`, site.SiteDeploy)
				err := helper.KillProcessesByWindowTitle(siteTitle)
				if err != nil {
					c.SSEvent("error", fmt.Sprintf("Failed to kill existing processes: %v", err))
					return
				}
			}

			commands := []string{
				fmt.Sprintf(`appcmd stop site /site.name:%s`, siteTitle),
				fmt.Sprintf("git pull %s %s", site.GitToken, site.GitBranch),
				siteCommand,
				removeFilesCommand,
				fmt.Sprintf("xcopy /s /y %s %s", site.SiteClone, site.SiteDeploy),
				fmt.Sprintf(`appcmd start site /site.name:%s`, siteTitle),
			}

			helper.RunMultiCommandsWithStream(c, commands, siteFolder, fmt.Sprintf("Deploy %s", siteTitle), site.GitToken)
		})
	}

	router.POST(fmt.Sprintf("/api/upload"), func(c *gin.Context) {
		// Check if the upload directory exists, if not, create it
		if _, err := os.Stat("./upload"); os.IsNotExist(err) {
			// Create the upload directory if it doesn't exist
			err := os.MkdirAll("./upload", os.ModePerm)
			if err != nil {
				// Log the error and return a response if directory creation fails
				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create upload directory: %v", err)})
				return
			}
		}

		// Parse the uploaded file
		file, header, err := c.Request.FormFile("file")
		log.Println(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to get file: %v", err)})
			return
		}
		defer file.Close()

		// Define the target path for saving the file
		filePath := filepath.Join("./upload", header.Filename)

		// Create the destination file
		outFile, err := os.Create(filePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to save file: %v", err)})
			return
		}
		defer outFile.Close()

		// Copy uploaded file data to the destination
		_, err = io.Copy(outFile, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to write file: %v", err)})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "path": filePath})
	})

	// get all menu
	router.GET("/api/get-site", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": config.Sites,
		})
	})

	router.POST("/api/login", func(c *gin.Context) {
		var user dto.UserLoginDTO
		json.NewDecoder(c.Request.Body).Decode(&user)

		if user.Email == "" || user.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
			return
		}
		if user.Email != "admin@hugcode.co.th" || user.Password != "hc-password" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// Generate JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": user.Email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})
		secretKey := []byte("h5G8sK3pQ1zY9vM2wX7dT6aJ0cL4NqR")
		tokenString, _ := token.SignedString(secretKey)

		json.NewEncoder(c.Writer).Encode(map[string]string{"token": tokenString})
	})

	// Start the server
	port := "8032"
	if os.Getenv("ASPNETCORE_PORT") != "" { // get environment variable set by ACNM
		port = os.Getenv("ASPNETCORE_PORT")
	}
	log.Fatal(router.Run(":" + port))
}
