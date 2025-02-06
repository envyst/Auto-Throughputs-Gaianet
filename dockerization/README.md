# Auto Throughputs Gaianet

## Prerequisites
- SSH access to a VPS/VM
- Installed Git

## Installation

Follow these steps to set up the project:

1. Clone the repository:
   ```sh
   git clone https://github.com/envyst/Auto-Throughputs-Gaianet.git
   ```
2. Change to the `dockerization` directory:
   ```sh
   cd Auto-Throughputs-Gaianet/dockerization
   ```
3. Install Docker:
   ```sh
   chmod +x install_docker.sh
   ./install_docker.sh
   ```
4. Navigate to the `instruct/` directory:
   ```sh
   cd instruct/
   ```
5. Build the Docker image:
   ```sh
   chmod +x build.sh run.sh
   ./build.sh
   ```

## Usage 
### Ensure You're in directory `Auto-Throughputs-Gaianet/dockerization/instruct`
```sh
pwd
```

### Step 1: Run Node (2 Nodes Recommended)
1. Start a node:
   ```sh
   ./run.sh
   ```
   Note the name of the new container.
2. View logs, wait until finish to get `node-id` and `device-id`:
   ```sh
   docker logs -f <container_name>
   ```
3. Register your node to your account at [Gaianet Settings](https://www.gaianet.ai/setting/nodes).
4. Join your node to a domain (e.g., `traders`, `vkvik`, or any other of your choice).
5. Press `Ctrl + C` on your terminal to exit from logs

### Step 2 (Optional): Run Auto Throughputs
1. Edit the API configuration:
   ```sh
   docker exec <container_name> nano /truput/api.txt
   ```
   - Fill in your **API-KEY** ([Get API-KEY](https://www.gaianet.ai/setting/gaia-api-keys)).
   - Fill in the **API-URL** ([Find your domain URL](https://www.gaianet.ai/chat)).
   - Save the file (`Ctrl + X`, then `Y`, then `Enter`).
2. Start auto throughputs:
   ```sh
   docker exec -d <container_name> /truput/agt
   ```
   This will run auto throughputs with 2 threads.
3. Verify that auto throughputs is running:
   ```sh
   docker exec <container_name> ps aux | grep agt
   ```
   Ensure that the process has a **PID** and is running.

---

Your node should now be set up and participating in the Gaianet network successfully!

