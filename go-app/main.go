package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type RequestBody struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Response struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func clearConsole() {
	fmt.Print("\033[2J\033[H")
}

func processChats(nodeID string) error {
	clearConsole()

	// Print the NodeID for testing
	fmt.Printf("Using NodeID: %s\n", nodeID)

	content, err := os.ReadFile("cersex.txt")
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	totalChats := len(lines) - 11
	if totalChats < 0 {
		totalChats = 0
	}
	fmt.Printf("Total chats to process: %d\n", totalChats)

	for index := 11; index < len(lines); index++ {
		chet := strings.TrimSpace(lines[index])
		if chet == "" {
			continue
		}
		currentIndex := index - 10
		fmt.Printf("Processing Chat %d/%d\n", currentIndex, totalChats)
		fmt.Println("Content Chat: " + chet + "\n")

		url := fmt.Sprintf("https://%s.gaia.domains/v1/chat/completions", nodeID)
		messages := []Message{
			{
				Role:    "system",
				Content: "You are a helpful assistant.",
			},
			{
				Role:    "user",
				Content: chet,
			},
		}

		jsonData, err := json.Marshal(RequestBody{Messages: messages})
		if err != nil {
			fmt.Printf("Error marshaling JSON: %v\n", err)
			continue
		}

		client := &http.Client{Timeout: 30 * time.Second}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Printf("Error creating request: %v\n", err)
			continue
		}
		req.Header.Set("accept", "application/json")
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error sending request: %v\n", err)
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response body: %v\n", err)
			continue
		}

		var response Response
		if err := json.Unmarshal(body, &response); err != nil {
			fmt.Printf("Error unmarshaling response: %v\n", err)
			continue
		}

		if len(response.Choices) > 0 {
			fmt.Printf("Response: [%s]\n\n", response.Choices[0].Message.Content)
		} else {
			fmt.Println("Response: No choices available\n")
		}

		fmt.Printf("⌛ Waiting 20 seconds... (%d/%d)\n", currentIndex, totalChats)
		time.Sleep(20 * time.Second)
	}

	fmt.Printf("✅ Completed processing %d chats.\n", totalChats)
	fmt.Println("⏩ Restarting the process...\n")
	return nil
}

func main() {
	logo := `
	@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@@%@@@@@@@@@@@@**@@@@@@@@@@@@@@@@@%*=+@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@%+@@@@@@@@@@%=.+@@@@@@@@@@@@@%#+-:. -@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@+=@@@@@@@@@*:..*@@@@@@@@@%#+-:......-@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@%:=@@@@@@@#-....*@@@@@%#+-:..........-@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@* =@@@@@%+......*@@#*=:..............-@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@-..................-@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@- =@@@@@=.......*@:..................-@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@- +@@@@@=.......*@:..................=@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@-.#@@@@@=.....:*@@:..............:=*%@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@--@@@@@@=....=%@@@:..........:=*#@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@-*@@@@@@=..:#@@@@@:......:-+#%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@+%@@@@@@-.+%@@@@@@:..:-+#%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@%@@@@@@@+#@@@@@@@@=+#%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@\n@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
	`
	fmt.Println(logo)
	fmt.Printf("\x1b[36m%s\x1b[0m\n", `
╔═══════════════════════════════════════╗
║        Auto Throughputs Gaianet       ║
║         By: Aethereal Team            ║
╚═══════════════════════════════════════╝
`)
	fmt.Println("🔄 Initializing...")
	// Define a command-line flag for NodeID
	nodeID := flag.String("nodeID", "", "The Node ID to use for processing chats")
	flag.Parse()

	// Validate the NodeID
	if *nodeID == "" {
		log.Fatal("\n🚫 Error: Node ID is required. Usage: go run main.go -nodeID <your-node-id>\n🚫 Error: Node ID is required. Usage: ./program -nodeID <your-node-id>")
	}

	// Print the NodeID for testing
	fmt.Printf("NodeID provided: %s\n", *nodeID)

	// Run the process in a loop
	for {
		err := processChats(*nodeID)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			fmt.Println("⏩ Restarting due to an error...\n")
		}
		time.Sleep(2 * time.Second)
	}
}
