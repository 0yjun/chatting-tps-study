const { socketLogger } = require("../logs/winston");

class Room {
  constructor(params) {
    this.forward = new Map();
    this.clients = new Set();
  }

  join(client) {
    socketLogger.info(`client join`);
    this.clients.add(client);
  }

  leave(client) {
    socketLogger.info(`client leave`);
    this.clients.delete(client);
  }

  forwardMessage(message) {
    socketLogger.info(`send message all client`);
    for (const client of clients) {
      client.send(JSON.stringify(message));
    }
  }
}

function newRoom() {
  return new Room();
}

module.exports = { newRoom };
