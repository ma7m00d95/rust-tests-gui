package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Exercise struct {
	Name string `json:"name"`
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Static("/", "static")

	e.GET("/api/exercises", func(c echo.Context) error {
		entries, err := os.ReadDir("solutions")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		exercises := []Exercise{}
		for _, entry := range entries {
			if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") {
				exercises = append(exercises, Exercise{Name: entry.Name()})
			}
		}
		return c.JSON(http.StatusOK, exercises)
	})

	e.GET("/api/run", func(c echo.Context) error {
		exercise := c.QueryParam("exercise")
		verbose := c.QueryParam("verbose") == "true"
		format := c.QueryParam("format") == "true"
		clippy := c.QueryParam("clippy") == "true"

		c.Response().Header().Set(echo.HeaderContentType, "text/event-stream")
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")
		c.Response().WriteHeader(http.StatusOK)

		// Smart Sync from student-solution folder
		studentFolder := "student-solution"
		if exercise != "" {
			// Find the target file in the solutions folder
			targetFile := "solutions/" + exercise + "/src/lib.rs"
			if _, err := os.Stat("solutions/" + exercise + "/src/main.rs"); err == nil {
				targetFile = "solutions/" + exercise + "/src/main.rs"
			}

			// 1. CLEAR the original solution first to ensure we aren't testing old code
			os.WriteFile(targetFile, []byte("// TODO: Student solution goes here\n"), 0644)

			if _, err := os.Stat(studentFolder); err == nil {
				sendSSE(c.Response().Writer, "output", "Searching for your work in student-solution...")
				
				variations := []string{exercise, strings.ReplaceAll(exercise, "_", "-")}
				found := false

				// Check for individual .rs files in student-solution
				for _, v := range variations {
					filePath := filepath.Join(studentFolder, v+".rs")
					if _, err := os.Stat(filePath); err == nil {
						sendSSE(c.Response().Writer, "output", fmt.Sprintf("Found your file: %s.rs", v))
						exec.Command("cp", filePath, targetFile).Run()
						found = true
						break
					}
				}

				// If not found as a root file, check for a directory
				if !found {
					for _, v := range variations {
						dirPath := filepath.Join(studentFolder, v)
						if info, err := os.Stat(dirPath); err == nil && info.IsDir() {
							sendSSE(c.Response().Writer, "output", fmt.Sprintf("Found your directory: %s", v))
							
							// Check if student provided a src folder or just files
							if _, err := os.Stat(filepath.Join(dirPath, "src")); err == nil {
								// They have a src folder, copy everything
								exec.Command("cp", "-a", dirPath+"/.", "solutions/"+exercise+"/").Run()
							} else {
								// No src folder, copy individual files to the right place
								files := []string{"main.rs", "lib.rs", "Cargo.toml"}
								for _, f := range files {
									fPath := filepath.Join(dirPath, f)
									if _, err := os.Stat(fPath); err == nil {
										if f == "Cargo.toml" {
											exec.Command("cp", fPath, "solutions/"+exercise+"/").Run()
										} else {
											exec.Command("cp", fPath, targetFile).Run()
										}
									}
								}
							}
							found = true
							break
						}
					}
				}

				if !found {
					sendSSE(c.Response().Writer, "output", "No student work found. Testing against a blank file.")
				}
			}
		}

		args := []string{}
		if verbose {
			args = append(args, "-v")
		}
		if format {
			args = append(args, "-f")
		}
		if clippy {
			args = append(args, "-c")
		}
		if exercise != "" {
			args = append(args, exercise)
		}

		fullArgs := append([]string{"tests/test_exercises.sh"}, args...)
		cmd := exec.Command("bash", fullArgs...)
		// Set CWD to the project root so the script finds its paths correctly
		cmd.Dir = "."

		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()

		if err := cmd.Start(); err != nil {
			sendSSE(c.Response().Writer, "error", err.Error())
			return nil
		}

		multi := io.MultiReader(stdout, stderr)
		scanner := bufio.NewScanner(multi)

		for scanner.Scan() {
			line := scanner.Text()
			sendSSE(c.Response().Writer, "output", line)
			c.Response().Flush()
		}

		if err := cmd.Wait(); err != nil {
			sendSSE(c.Response().Writer, "exit", fmt.Sprintf("Process exited with error: %v", err))
		} else {
			sendSSE(c.Response().Writer, "exit", "Process completed successfully")
		}

		return nil
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func sendSSE(w io.Writer, event string, data string) {
	fmt.Fprintf(w, "event: %s\ndata: %s\n\n", event, data)
}
