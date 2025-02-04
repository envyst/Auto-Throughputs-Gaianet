# Auto-Throughputs-Gaianet
Forked from : https://github.com/Aethereal-Collective/Auto-Throughputs-Gaianet

## Overview
Auto-Throughputs-Gaianet is a script designed to send automated Throughputs the Gaianet API using multiple threads. It continuously processes user messages and interacts with the API to retrieve responses.

## Node Join Domains
Join your nodes to domains like:
- cryptology.gaia.domains
- p*rnhub.gaia.domains
- esiyk.gaia.domains
- doge.gaia.domains
- ionet.gaia.domains
- openai.gaia.domains
- huggingface.gaia.domains

by accessing https://www.gaianet.ai/setting/nodes in three dots (...)

## Installation
To get started, download releases zip file from this repository:

```sh
# If you dont have curl and unzip
sudo apt update
sudo apt install -y curl unzip
```

```sh
# for LINUX/MAC AMD64
curl -L -o "amd64.zip" "https://github.com/envyst/Auto-Throughputs-Gaianet/releases/download/v2/amd64.zip"
unzip -o amd64.zip -d gaianet-auto-throughputs/
cd gaianet-auto-throughputs/
```

```sh
# for LINUX ARM64 / TERMUX / RASPBERRY PIE
curl -L -o "arm64.zip" "https://github.com/envyst/Auto-Throughputs-Gaianet/releases/download/v2/arm64.zip"
unzip -o arm64.zip -d gaianet-auto-throughputs/
cd gaianet-auto-throughputs/
```

### For windows, just download and extract from [releases](https://github.com/envyst/Auto-Throughputs-Gaianet/releases) and extract

## Setup
1. **Edit `api.txt`**
   - Open a terminal and run:
     ```sh
     nano api.txt
     ```
   - Add the following details:
     - **First line:** API Key (Obtain from [Gaianet Settings](https://www.gaianet.ai/setting))
     - **Second line:** API URL (Navigate to [Gaianet Chat](https://www.gaianet.ai/chat?domain) and select your domain)
   - Click on **API Tutorial** at the bottom of the "Go to Chat" screen.
   - Copy `/v1/chat/completions`.
   - Copy the API key and replace it in `api.txt`.
   - Format `api.txt` file
   ```txt
   APIKEY-GAIANET
   API-DOMAINURL
   ```

## Usage
Run the script with:
```sh
# for LINUX/MAC AMD64 and LINUX ARM64 / TERMUX / RASPBERRY PIE

# optional : screen -S gaianet-throughputs-v2
./agt_*
# you will be prompted to input thread count

# optional : Ctrl + A + D
```

```sh
# for Windows open terminal in the folder
.\agt_win.exe
# you will be prompted to input thread count
```


## License
This project is open-source and available under the MIT License.
