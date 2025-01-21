# Auto Throughputs Gaianet

## 🔧 Overview
Auto Throughputs Gaianet is a Node.js script developed by the Aethereal Team. It automates the processing of chat messages through Gaianet's chat API, handling multiple chat messages at once and outputting the response.

## ⚙️ Features
- Processes chat messages from a text file
- Automatically sends requests to the Gaianet API.
- Waits for 20 seconds between requests.
- Loops the script continuously until manually stopped.

## 🚀 Installation

### Prerequisites:
- Node.js (v14 or higher)
- `cersex.txt` file containing the chat messages you want to process.

### Steps:
1. Clone the repository:
   ```bash
   git clone https://github.com/Aethereal-Collective/Auto-Throughputs-Gaianet
   ```
2. Create screen
   ```bash
   screen -S gaianet-throughputs
   ```
3. Navigate to directory:
   ```bash
   cd Auto-Throughputs-Gaianet
   ```
4. Install dependencies:
   ```bash
   npm install
   ```
5. Run the script:
   ```bash
   node index.js
   ```
Follow the prompts to enter your Node ID.

## 📄 Usage
Once you run the script, it will process the chats listed in `cersex.txt`. For each chat, the script will:

- Send the chat message to the Gaianet API.
- Log the response from the API.
- Wait 20 seconds before proceeding to the next chat message.
- Loop indefinitely until you manually stop the process.

## 🤝Support
Stay connected and be part of our community:

- [Join our Discord](https://discord.gg/aethereal)  
- [Follow us on Twitter](https://x.com/aethereal_co)
