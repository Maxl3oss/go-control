package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

		router.POST(fmt.Sprintf("/api/%s/build", siteTitle), func(c *gin.Context) {
			c.Header("Content-Type", "text/event-stream")
			c.Header("Cache-Control", "no-cache")
			c.Header("Connection", "keep-alive")

			// Run the build process and stream logs
			runCommandWithStream(c, siteCommand, siteFolder, fmt.Sprintf("Build %s", siteTitle))
		})

		router.POST(fmt.Sprintf("/api/%s/pull", siteTitle), func(c *gin.Context) {
			c.Header("Content-Type", "text/event-stream")
			c.Header("Cache-Control", "no-cache")
			c.Header("Connection", "keep-alive")

			// Run the build process and stream logs
			cmd := fmt.Sprintf("git pull %s %s", site.GitToken, site.GitBranch)
			runCommandWithStream(c, cmd, siteFolder, fmt.Sprintf("Pull %s", siteTitle))
		})

		router.POST(fmt.Sprintf("/api/%s/stop-service", siteTitle), func(c *gin.Context) {
			c.Header("Content-Type", "text/event-stream")
			c.Header("Cache-Control", "no-cache")
			c.Header("Connection", "keep-alive")

			// Run the build process and stream logs
			cmd := fmt.Sprintf(`appcmd stop site /site.name:%s`, siteTitle)
			runCommandWithStream(c, cmd, siteFolder, fmt.Sprintf("Deploy %s", siteTitle))
		})

		router.POST(fmt.Sprintf("/api/%s/start-service", siteTitle), func(c *gin.Context) {
			c.Header("Content-Type", "text/event-stream")
			c.Header("Cache-Control", "no-cache")
			c.Header("Connection", "keep-alive")

			// Run the build process and stream logs
			cmd := fmt.Sprintf(`appcmd start site /site.name:%s`, siteTitle)
			runCommandWithStream(c, cmd, siteFolder, fmt.Sprintf("Deploy %s", siteTitle))
		})

		router.POST(fmt.Sprintf("/api/%s/deploy", siteTitle), func(c *gin.Context) {
			c.Header("Content-Type", "text/event-stream")
			c.Header("Cache-Control", "no-cache")
			c.Header("Connection", "keep-alive")

			// Run the build process and stream logs
			cmd := fmt.Sprintf("xcopy /s /y %s %s", site.SiteClone, site.SiteDeploy)
			runCommandWithStream(c, cmd, siteFolder, fmt.Sprintf("Deploy %s", siteTitle))
		})

		router.POST(fmt.Sprintf("/api/%s/auto-deploy", siteTitle), func(c *gin.Context) {
			c.Header("Content-Type", "text/event-stream")
			c.Header("Cache-Control", "no-cache")
			c.Header("Connection", "keep-alive")

			// Define the list of commands
			commands := []string{
				fmt.Sprintf("git pull %s %s", site.GitToken, site.GitBranch),
				siteCommand,
				// "cp -r build /path/to/deploy",
			}

			// Run the commands in sequence
			runMultiCommandsWithStream(c, commands, siteFolder, fmt.Sprintf("Deploy %s", siteTitle), site.GitToken)
		})
	}

	// get all menu
	router.GET("/api/get-site", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": config.Sites,
		})
	})

	// Start the server
	port := "8032"
	if os.Getenv("ASPNETCORE_PORT") != "" { // get environment variable set by ACNM
		port = os.Getenv("ASPNETCORE_PORT")
	}
	log.Fatal(router.Run(":" + port))
}

func runCommandWithStream(c *gin.Context, commandStr string, workingDir string, action string) {
	var shell, flag string
	if runtime.GOOS == "windows" {
		shell, flag = "cmd", "/C"
	} else {
		shell, flag = "sh", "-c"
	}

	command := exec.Command(shell, flag, fmt.Sprintf("cd %s && %s", workingDir, commandStr))
	stdoutPipe, err := command.StdoutPipe()
	if err != nil {
		log.Printf("[%s] Error creating stdout pipe: %v", action, err)
		c.SSEvent("error", fmt.Sprintf("[%s] Failed to initialize output streaming", action))
		c.Writer.Flush()
		return
	}
	stderrPipe, err := command.StderrPipe()
	if err != nil {
		log.Printf("[%s] Error creating stderr pipe: %v", action, err)
		c.SSEvent("error", fmt.Sprintf("[%s] Failed to initialize error streaming", action))
		c.Writer.Flush()
		return
	}

	if err := command.Start(); err != nil {
		log.Printf("[%s] Command failed to start: %v", action, err)
		c.SSEvent("error", fmt.Sprintf("[%s] Failed to start command", action))
		c.Writer.Flush()
		return
	}

	c.SSEvent("status", fmt.Sprintf("[%s] Command started successfully", action))
	c.Writer.Flush()

	done := make(chan bool)
	go func() {
		if _, err := io.Copy(c.Writer, stdoutPipe); err != nil {
			log.Printf("[%s] Error streaming stdout: %v", action, err)
		}
		done <- true
	}()
	go func() {
		if _, err := io.Copy(c.Writer, stderrPipe); err != nil {
			log.Printf("[%s] Error streaming stderr: %v", action, err)
		}
		done <- true
	}()

	<-done
	<-done

	if err := command.Wait(); err != nil {
		log.Printf("[%s] Command execution failed: %v", action, err)
		c.SSEvent("error", fmt.Sprintf("[%s] Command execution failed: %v", action, err))
		c.Writer.Flush()
		return
	}

	c.SSEvent("done", fmt.Sprintf("[%s] Command executed successfully", action))
	c.Writer.Flush()
}

func runMultiCommandsWithStream(c *gin.Context, commands []string, workingDir string, action string, hiddenCommand string) {
	var shell, flag string
	if runtime.GOOS == "windows" {
		shell, flag = "cmd", "/C"
	} else {
		shell, flag = "sh", "-c"
	}

	for _, commandStr := range commands {
		command := exec.Command(shell, flag, fmt.Sprintf("cd %s && %s", workingDir, commandStr))
		showCommand := filterCommand(commandStr, hiddenCommand)

		stdoutPipe, err := command.StdoutPipe()
		if err != nil {
			log.Printf("[%s] Error creating stdout pipe: %v", action, err)
			c.SSEvent("error", fmt.Sprintf("[%s] Failed to initialize output streaming for command: %s", action, showCommand))
			c.Writer.Flush()
			return
		}
		stderrPipe, err := command.StderrPipe()
		if err != nil {
			log.Printf("[%s] Error creating stderr pipe: %v", action, err)
			c.SSEvent("error", fmt.Sprintf("[%s] Failed to initialize error streaming for command: %s", action, showCommand))
			c.Writer.Flush()
			return
		}

		if err := command.Start(); err != nil {
			log.Printf("[%s] Command failed to start: %v", action, err)
			c.SSEvent("error", fmt.Sprintf("[%s] Failed to start command: %s", action, showCommand))
			c.Writer.Flush()
			return
		}

		c.SSEvent("status", fmt.Sprintf("[%s] Running command: %s", action, showCommand))
		c.Writer.Flush()

		done := make(chan bool)
		go func() {
			if _, err := io.Copy(c.Writer, stdoutPipe); err != nil {
				log.Printf("[%s] Error streaming stdout for command: %s, %v", action, showCommand, err)
			}
			done <- true
		}()
		go func() {
			if _, err := io.Copy(c.Writer, stderrPipe); err != nil {
				log.Printf("[%s] Error streaming stderr for command: %s, %v", action, showCommand, err)
			}
			done <- true
		}()

		<-done
		<-done

		if err := command.Wait(); err != nil {
			log.Printf("[%s] Command execution failed for command: %s, %v", action, showCommand, err)
			c.SSEvent("error", fmt.Sprintf("[%s] Command execution failed for command: %s, %v", action, showCommand, err))
			c.Writer.Flush()
			return
		}

		c.SSEvent("done", fmt.Sprintf("[%s] Command executed successfully: %s", action, showCommand))
		c.Writer.Flush()
	}

	c.SSEvent("done", fmt.Sprintf("[%s] All commands executed successfully", action))
	c.Writer.Flush()
}

func filterCommand(commandStr, hiddenCommand string) string {
	// Split the command string into parts
	parts := strings.Split(commandStr, " ")

	// Filter out the hiddenCommand
	var filteredParts []string
	for _, part := range parts {
		if part != hiddenCommand {
			filteredParts = append(filteredParts, part)
		} else {
			filteredParts = append(filteredParts, "********")
		}
	}

	// Join the filtered parts back into a single string
	return strings.Join(filteredParts, " ")
}
