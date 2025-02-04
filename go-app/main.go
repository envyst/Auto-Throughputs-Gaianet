package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

const logo = `
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@%@@@@@@@@@@@@**@@@@@@@@@@@@@@@@@%*=+@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@%+@@@@@@@@@@%=.+@@@@@@@@@@@@@%#+-:. -@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@+=@@@@@@@@@*:..*@@@@@@@@@%#+-:......-@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@%:=@@@@@@@#-....*@@@@@%#+-:..........-@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@* =@@@@@%+......*@@#*=:..............-@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@-..................-@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@- =@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................=@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@-.#@@@@@=.....:*@@:..............:=*%@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@--@@@@@@=....=%@@@:..........:=*#@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@-*@@@@@@=..:#@@@@@:......:-+#%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@+%@@@@@@-.+%@@@@@@:..:-+#%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@%@@@@@@@+#@@@@@@@@=+#%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
`

var (
	client = &http.Client{}
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Messages []Message `json:"messages"`
}

func sendRequest(message string, apiKey string, apiURL string) {
	headers := map[string]string{
		"Authorization": "Bearer " + apiKey,
		"accept":        "application/json",
		"Content-Type":  "application/json",
	}

	data := RequestBody{
		Messages: []Message{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: message},
		},
	}

	for {
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Printf("Error marshaling JSON: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Printf("Error creating request: %v. Retrying...\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for k, v := range headers {
			req.Header.Set(k, v)
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Request failed with error: %v. Retrying...\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			var responseJSON map[string]interface{}
			if err := json.Unmarshal(bodyBytes, &responseJSON); err != nil {
				fmt.Printf("Error decoding JSON response for message '%s': %v\n", message, err)
				fmt.Printf("Response Text: %s\n", string(bodyBytes))
			} else {
				fmt.Printf("Response for message: '%s'\n", message)
				fmt.Println(responseJSON)
				return
			}
		} else {
			fmt.Printf("Error: %d, %s. Retrying...\n", resp.StatusCode, string(bodyBytes))
		}

		time.Sleep(5 * time.Second)
	}
}

func runThread(userMessages []string, apiKey string, apiURL string) {
	for {
		randIndex := rand.Intn(len(userMessages))
		msg := userMessages[randIndex]
		sendRequest(msg, apiKey, apiURL)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	apiContent, err := os.ReadFile("api.txt")
	if err != nil {
		log.Fatalf("Error reading api.txt: %v", err)
	}
	apiLines := strings.Split(strings.ReplaceAll(string(apiContent), "\r\n", "\n"), "\n")
	if len(apiLines) < 2 {
		log.Fatal("api.txt must contain at least two lines")
	}
	apiKey := strings.TrimSpace(apiLines[0])
	apiURL := strings.TrimSpace(apiLines[1])

	cersexContent, err := os.ReadFile("cersex.txt")
	if err != nil {
		log.Fatalf("Error reading cersex.txt: %v", err)
	}
	cersexLines := strings.Split(strings.ReplaceAll(string(cersexContent), "\r\n", "\n"), "\n")
	var userMessages []string
	for _, line := range cersexLines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			userMessages = append(userMessages, trimmed)
		}
	}
	if len(userMessages) == 0 {
		log.Fatal("cersex.txt must contain at least one message")
	}

	fmt.Print("\033[36m")
	fmt.Println(`
╔═══════════════════════════════════════╗
║        Auto Throughputs Gaianet       ║
║         By: Aethereal Team            ║
╚═══════════════════════════════════════╝`)
	fmt.Print("\033[0m")
	fmt.Print(logo)
	fmt.Println("🔄 Initializing...")

	var numThreads int
	fmt.Print("Please specify the number of threads to be used: ")
	_, err = fmt.Scanln(&numThreads)
	if err != nil || numThreads < 1 {
		fmt.Println("Invalid input. Enter a number greater than zero.")
		return
	}

	for i := 0; i < numThreads; i++ {
		go runThread(userMessages, apiKey, apiURL)
	}

	select {}

	// this line blocks the main goroutine from exiting, allowing the threads to run indefinitely
	select {}
}
