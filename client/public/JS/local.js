import { Jogador, Movimento } from '../jogo.js';

let tabuleiroContainer = document.getElementById("tabuleiro-div");
const terminarVezBtn = document.getElementById("terminar-vez-btn");
const desistirBtn = document.getElementById("desistir-btn");
const restartBtn = document.getElementById("restart-btn");

let startSound = new Audio("../static/startSound.mkv");
startSound.crossOrigin = "anonymous";
let endSound = new Audio("../static/lichessCheckmate.mkv");
endSound.crossOrigin = "anonymous";

let salaId, jogador1, jogador2, turno = 'B';
const tempo = 180;
let movimentos = [];

criarTabuleiro(tabuleiroContainer);


let socket = io();
socket.on("connect", async () => {
    terminarVezBtn.addEventListener("click", () => terminarVez(socket.id));
    desistirBtn.addEventListener("click", () => desistir(socket.id, salaId));
    restartBtn.addEventListener("click", () => restartJogo(socket.id, salaId));

    jogador1 = new Jogador(socket.id, null, tempo, null);
    jogador2 = new Jogador(socket.id, null, tempo, null);

    await fetch("/jogo/criar-sala", {
        method: "POST",
        headers: {
            "Content-type": "Application/JSON"
        },
        body: JSON.stringify({
            socketId: socket.id,
            tempo
        })
    })
    .then((res) => { return res.json(); })
    .then((res) => {
        if (res.error) {
            alert(res.error);
            window.location.href = "/";
        }

        salaId = res.salaId

        entrarSala(jogador1);
        entrarSala(jogador2);

        startSound.play();
    });

});


function entrarSala(jogador) {
    fetch("/jogo/entrar-sala", {
        method: "POST",
        headers: {
            "Content-type": "Application/JSON"
        },
        body: JSON.stringify({
            salaId,
            socketId: jogador.socketId,
            username: jogador.username,
            tempo: jogador.tempo
        })
    })
    .then((res) => { return res.json(); })
    .then((res) => {
        if (res.error) {
            alert(res.error);
            window.location.href = "/";
        }

        if (res.sala.jogadores[0].socketId == jogador.socketId)
            jogador.lado = res.sala.jogadores[0].lado
        else
            jogador.lado = res.sala.jogadores[1].lado
    });
}

function criarTabuleiro(tabuleiro) {
    for (let i = 0; i < 5; i++) {
        const linha = document.createElement("div");
        linha.classList.add("centralizador");

        for (let j = 0; j < 5; j++) {
            const casa = document.createElement("div");
            casa.classList.add("casa-tabuleiro");

            if (i === 0 && j === 0)
                casa.classList.add("casa-topo-esquerda");
            else if (i === 0 && j === 4)
                casa.classList.add("casa-topo-direita");
            else if (i === 4 && j === 0)
                casa.classList.add("casa-baixo-esquerda");
            else if (i === 4 && j === 4)
                casa.classList.add("casa-baixo-direita");

            casa.y = i;
            casa.x = j;

            casa.addEventListener("click", () => {
                movimentos.push(new Movimento(i, j, turno));
            });

            linha.appendChild(casa);
        }

        tabuleiro.appendChild(linha);
    }
}

function atualizarTabuleiro(tabuleiroDiv, tabuleiro) {
    // Limpar tabuleiro antigo
    for (let i = 0; i < tabuleiro.length; i++) {
        for (let j = 0; j < tabuleiro[i].length; j++) {
            while (tabuleiroDiv.children[i].children[j].hasChildNodes()) {
                tabuleiroDiv.children[i].children[j].firstChild.remove();
            }
        }
    }

    // Colocar peças no tabuleiro
    for (let i = 0; i < tabuleiro.length; i++) {
        for (let j = 0; j < tabuleiro[i].length; j++) {
            if (tabuleiro[i][j] != "") {
                const peca = document.createElement("div");
                (tabuleiro[i][j] === 'B' 
                    ? peca.classList.add("peca-branca")
                    : peca.classList.add("peca-preta")
                );

                tabuleiroDiv.children[i].children[j].appendChild(peca);
            }
        }
    }
}

function terminarVez(socketId) {
    fetch("/jogo", {
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
    .then((res) => { return res.json(); })
    .then((res) => {
        if (res.error) {
            movimentos = [];
            return;
        }

        atualizarTabuleiro(tabuleiroContainer, res.sala.jogo.tabuleiro);

        if (res.sala.jogoEncerrado) {
            encerrarPartida();
            return;
        }

        movimentos = [];
        turno = res.sala.jogo.turno;
    });
}

function encerrarPartida() {
    terminarVezBtn.classList.add("hidden");
    desistirBtn.classList.add("hidden");
    
    restartBtn.classList.remove("hidden");

    endSound.play();
}

function desistir(socketId) {
    fetch("/jogo/desistir", {
        method: "POST",
        headers: {
            "Content-type": "Application/JSON"
        },
        body: JSON.stringify({
            salaId,
            socketId
        })
    })
    .then((res) => { return res.json(); })
    .then((res) => {
        if (res.error) {
            alert(res.error);
            return;
        }

        encerrarPartida();
    });
}

function restartJogo(socketId, salaId) {
    fetch("/jogo/restart", {
        method: "POST",
        headers: {
            "Content-type": "Application/JSON"
        },
        body: JSON.stringify({
            salaId,
            socketId
        })
    })
    .then((res) => { return res.json(); })
    .then((res) => {
        if (res.error) {
            alert(res.error);
            return;
        }

        reiniciarPartida(tabuleiroContainer);
    });
}

function reiniciarPartida(tabuleiro) {
    turno = 'B';
    movimentos = [];

    tabuleiro.remove();
    tabuleiroContainer = document.createElement("div");
    tabuleiroContainer.id = "tabuleiro-div";

    criarTabuleiro(tabuleiroContainer);

    const tabuleiroDivContainer = document.getElementById("tabuleiro-div-container");
    tabuleiroDivContainer.appendChild(tabuleiroContainer);

    jogador1.tempo = tempo;
    jogador2.tempo = tempo;

    terminarVezBtn.classList.remove("hidden");
    desistirBtn.classList.remove("hidden");

    restartBtn.classList.add("hidden");

    startSound.play();
}

document.addEventListener("keydown", (e) => {
    const key = e.key;

    if (key == 'p' || key == 'P'){
        terminarVez(socket.id);
    }
    else if(key == 'd' || key == 'D'){
        desistir(socket.id);
    }
});

tabuleiroContainer.addEventListener("contextmenu", (e) => {
    e.preventDefault();
});