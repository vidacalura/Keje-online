const tabuleiroContainer = document.getElementById("tabuleiro-div");
const localBtn = document.getElementById("btn-jogar-local");
const menuTempoLocal = document.getElementById("menu-tempo-container");
localBtn.addEventListener("click", mudarVisualizacaoMenu);
menuTempoLocal.addEventListener("click", mudarVisualizacaoMenu);

const sleep = ms => new Promise(r => setTimeout(r, ms));
const sleepTime = 500;

criarTabuleiro(tabuleiroContainer);
simulacaoJogo(tabuleiroContainer);


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

            casa.addEventListener("click", () => {
                
            });

            linha.appendChild(casa);
        }

        tabuleiro.appendChild(linha);
    }
}

async function simulacaoJogo(tabuleiro) {
    const pecaBranca = document.createElement("div");
    pecaBranca.classList.add("peca-branca");

    const pecaPreta = document.createElement("div");
    pecaPreta.classList.add("peca-preta");

    // 1. A1 B1 C1
    tabuleiro.children[0].children[0].appendChild(pecaBranca.cloneNode(true));
    await sleep(sleepTime);
    tabuleiro.children[0].children[1].appendChild(pecaBranca.cloneNode(true));
    await sleep(sleepTime);
    tabuleiro.children[0].children[2].appendChild(pecaBranca.cloneNode(true));
    await sleep(sleepTime);

    // 2. B2 B3 B4
    tabuleiro.children[1].children[1].appendChild(pecaPreta.cloneNode(true));
    await sleep(sleepTime);
    tabuleiro.children[2].children[1].appendChild(pecaPreta.cloneNode(true));
    await sleep(sleepTime);
    tabuleiro.children[3].children[1].appendChild(pecaPreta.cloneNode(true));
    await sleep(sleepTime);

    // 3. C4
    tabuleiro.children[3].children[2].appendChild(pecaBranca.cloneNode(true));
    await sleep(sleepTime);

    // 4. C5
    tabuleiro.children[4].children[2].appendChild(pecaPreta.cloneNode(true));
    await sleep(sleepTime);

    // 5. E4 E5
    tabuleiro.children[3].children[4].appendChild(pecaBranca.cloneNode(true));
    await sleep(sleepTime);
    tabuleiro.children[4].children[4].appendChild(pecaBranca.cloneNode(true));
    await sleep(sleepTime);

    // 6. D3
    tabuleiro.children[2].children[3].appendChild(pecaPreta.cloneNode(true));
    await sleep(sleepTime);

    // 7. E2
    tabuleiro.children[1].children[4].appendChild(pecaBranca.cloneNode(true));
    await sleep(sleepTime);

    // 8. D4
    tabuleiro.children[3].children[3].appendChild(pecaPreta.cloneNode(true));
    await sleep(sleepTime);

    // 9. A2 A3
    tabuleiro.children[1].children[0].appendChild(pecaBranca.cloneNode(true));
    await sleep(sleepTime);
    tabuleiro.children[2].children[0].appendChild(pecaBranca.cloneNode(true));
    await sleep(sleepTime);

    // 10. D1 E1
    tabuleiro.children[0].children[3].appendChild(pecaPreta.cloneNode(true));
    await sleep(sleepTime);
    tabuleiro.children[0].children[4].appendChild(pecaPreta.cloneNode(true));
    await sleep(sleepTime);

    // 11. C2
    tabuleiro.children[1].children[2].appendChild(pecaBranca.cloneNode(true));
    await sleep(sleepTime);

    // 12. A5 B5
    tabuleiro.children[4].children[0].appendChild(pecaPreta.cloneNode(true));
    await sleep(sleepTime);
    tabuleiro.children[4].children[1].appendChild(pecaPreta.cloneNode(true));
    await sleep(sleepTime);
}

function mudarVisualizacaoMenu(e) {
    if (menuTempoLocal.className.includes("hidden")) {
        menuTempoLocal.classList.add("centralizador");
        menuTempoLocal.classList.remove("hidden");
    }
    else {
        if (e.target.id == "menu-tempo-container") {
            menuTempoLocal.classList.remove("centralizador");
            menuTempoLocal.classList.add("hidden");
        }
    }
}