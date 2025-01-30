const axios = require("axios");
const fs = require("fs");
const readline = require("readline");
const hitamlegam = require('./src/lomgo.js');

const delay = (ms) => new Promise((resolve) => setTimeout(resolve, ms));

const rli = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

const processChats = async (NodeID) => {
  try {
    console.clear();
    console.log(
      '\x1b[36m%s\x1b[0m',
      `
╔═══════════════════════════════════════╗
║        Auto Throughputs Gaianet       ║
║         By: Aethereal Team            ║
╚═══════════════════════════════════════╝
`
    );
    console.log("🔄 Initializing...");
    console.log(hitamlegam);

    const addressList = await fs.readFileSync("cersex.txt", "utf-8");
    const addressListArray = await addressList.split("\n");

    const totalChats = addressListArray.length - 11;
    console.log(`Total chats to process: ${totalChats}`);

    for (let index = 11; index < addressListArray.length; index++) {
      const chet = addressListArray[index];
      const currentIndex = index - 10;
      console.log(`Processing Chat ${currentIndex}/${totalChats}`);
      console.log("Content Chat: " + chet + "\n");

      try {
        const response = await axios.post(
          `https://${NodeID}.gaia.domains/v1/chat/completions`,
          {
            messages: [
              {
                role: "system",
                content: "You are a helpful assistant.",
              },
              {
                role: "user",
                content: `${chet}`,
              },
            ],
          },
          {
            headers: {
              accept: "application/json",
              "Content-Type": "application/json",
            },
          }
        );

        console.log("Response: [" + response.data.choices[0].message.content + "]\n");
        console.log(`⌛ Waiting 20 seconds... (${currentIndex}/${totalChats})`);
        await delay(20000);
      } catch (postError) {
        console.error("Error during axios post: ", postError);
      }
    }

    console.log(`? Completed processing ${totalChats} chats.`);
    console.log('⏩ Restarting the process...\n');
    await delay(2000);
    processChats(NodeID);
  } catch (error) {
    console.error("Error: ", error);
    console.log('⏩ Restarting due to an error...\n');
    await delay(2000);
    processChats(NodeID); 
  }
};

(async () => {
  try {
    console.clear();
    const NodeID = await new Promise((resolve) =>
      rli.question("↪ Enter your Node ID: ", (input) => resolve(input.trim()))
    );

    if (!NodeID) {
      console.error("🚫 Error: Node ID is required.");
      rli.close();
      return;
    }

    processChats(NodeID);
  } catch (error) {
    console.error("Error: ", error);
    rli.close();
  }
})();
