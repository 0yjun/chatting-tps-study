const express = require("express");
const http = require("http");
const cors = require("cors");
const WebSocket = require("ws");
const app = express();

app.use(
  cors({
    origin: "*",
  })
);
app.use(express.json());
app.use(express.urlencoded({ extended: false }));

const server = http.createServer();
//const wss = new WebSocket.Server({ server });

const PORT = 8080;

server.listen(PORT, () => {
  console.log(`server started on port ${PORT}`);
});
