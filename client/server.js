require("dotenv").config();
const express = require("express");
const socket = require("socket.io");
const fetch = require("node-fetch");

const app = express();

const port = process.env.PORT || 5000;
const server = app.listen(port);

app.use(express.static(__dirname + "/public/"));
app.use(express.static(__dirname + "/modules/"));
app.use(express.json());
app.use(require("./routes/routes"));

let io = socket(server);
io.on("connection", (socket) => {

    socket.on("disconnect", () => {
        fecharSala(socket.id);
    });
});


function fecharSala(socketId) {
    fetch(process.env.API + "/api/jogo/salas", {
        method: "DELETE",
        headers: {
            "Content-type": "Application/JSON"
        },
        body: JSON.stringify({
            socketId
        })
    })
    .then((response) => { return response.json(); })
    .then((response) => {
        // Avisar outros integrantes da sala que jogador saiu
    });
}