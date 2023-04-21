export class Jogador {
    constructor(socketId, username, tempo, lado) {
        this.socketId = socketId;
        this.username = username;
        this.tempo = tempo;
        this.lado = lado;
        this.pontos = 0; 
    }
}

export class Movimento {
    constructor(y, x, lado) {
        this.y = y;
        this.x = x;
        this.lado = lado;
    }
}