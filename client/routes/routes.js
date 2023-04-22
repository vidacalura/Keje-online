require("dotenv").config();
let path = require("path");
const express = require("express");
const fetch = require("node-fetch");

const router = express.Router();

router.get("/", (req, res) => {
    res.sendFile(path.join(__dirname, "../public/index.html"));
});

router.get("/local", (req, res) => {
    res.sendFile(path.join(__dirname, "../public/local.html"));
});

router.post("/jogo/criar-sala", async (req, res) => {
    const { socketId, tempo } = req.body;

    if (socketId == null || tempo == null){
        res.status(400).json({ "error": "Dados inválidos." });
        return
    }

    fetch(process.env.API + "/api/jogo/salas", {
        method: "POST",
        headers: {
            "Content-type": "Application/JSON"
        },
        body: JSON.stringify({
            socketId,
            tempo
        })
    })
    .then((response) => { return response.json(); })
    .then((response) => {
        res.json(response);
    });
});

router.post("/jogo/entrar-sala", async (req, res) => {
    const { salaId, socketId, username, tempo } = req.body;

    if (salaId == null || socketId == null || tempo == null){
        res.status(400).json({ "error": "Dados inválidos." });
        return
    }

    fetch(process.env.API + "/api/jogo/salas/entrar", {
        method: "POST",
        headers: {
            "Content-type": "Application/JSON"
        },
        body: JSON.stringify({
            salaId,
            socketId,
            username,
            tempo
        })
    })
    .then((response) => { return response.json(); })
    .then((response) => {
        res.json(response);
    });
});

router.post("/jogo", async (req, res) => {
    const { salaId, socketId, movimentos } = req.body;

    if (salaId == null || socketId == null || movimentos == null || movimentos == []){
        res.status(400).json({ "error": "Dados inválidos." });
        return
    }

    fetch(process.env.API + "/api/jogo", {
        method: "POST",
        headers: {
            "Content-type": "Application/JSON"
        },
        body: JSON.stringify({
            salaId,
            socketId,
            movimentos
        })
    })
    .then((response) => { return response.json(); })
    .then((response) => {
        res.json(response);
    });
});

router.post("/jogo/desistir", (req, res) => {
    const { salaId, socketId } = req.body;

    if (salaId == null || socketId == null) {
        res.status(400).json({ "error": "Dados inválidos." });
        return;
    }
    
    fetch(process.env.API + "/api/jogo/desistir", {
        method: "POST",
        headers: {
            "Content-type": "Application/JSON"
        },
        body: JSON.stringify({
            salaId,
            socketId
        })
    })
    .then((response) => { return response.json(); })
    .then((response) => {
        res.json(response);
    });
});

router.post("/jogo/restart", (req, res) => {
    const { salaId, socketId } = req.body;

    if (salaId == null || socketId == null) {
        res.status(400).json({ "error": "Dados inválidos." });
        return;
    }

    fetch(process.env.API + "/api/jogo/restart", {
        method: "POST",
        headers: {
            "Content-type": "Application/JSON"
        },
        body: JSON.stringify({
            salaId,
            socketId
        })
    })
    .then((response) => { return response.json(); })
    .then((response) => {
        res.json(response);
    });
});


module.exports = router;