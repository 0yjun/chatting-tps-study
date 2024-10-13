const express = require("express");
const http = require("http");
const { WebSocket } = require("ws");
const cors = require("cors");
const { socketLogger } = require("./logs/winston");
const { newRoom } = require("./type/Room");
const app = express();

app.use(
  cors({
    origin: "*",
  })
);
app.use(express.json);
app.use(express.urlencoded({ extended: false }));

const server = http.createServer();
const wss = new WebSocket.Server({ server });

const room = newRoom();

wss.on("connection", (ws, req) => {
  const cookie = req.headers.cookie;

  const [_, user] = cookie.split("=");

  room.join(ws);

  ws.on("message", (msg) => {
    const jsonMsg = JSON.parse(msg);
    jsonMsg.Name = user;
    room.forwardMessage(jsonMsg);
    console.log("msg: {}", msg);
  });
  ws.on("close", () => {
    room.leave(ws);
    console.log(`socket close`);
  });
});

const PORT = 8080;

server.listen(PORT, () => {
  socketLogger.info(`server started on port ${PORT}`);
});
