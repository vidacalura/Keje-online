
# noname API

## Sumário

* []()

## Status Codes

Os status de respostas possíveis para esta API são:

| Status Code | Description |
| :--- | :--- |
| 200 | `OK` |
| 400 | `BAD REQUEST` |
| 404 | `NOT FOUND` |
| 500 | `INTERNAL SERVER ERROR` |

## Responses

Os endpoints terão sempre 2 possíveis responses, 
sendo elas:

```javascript
{
  ...
  "message": String
}
```

```javascript
{
  "error": String
}
```

O atributo `message` estará presente caso uma request seja concluída com sucesso, e ausente caso contrário.

O atributo `error` estará presente caso uma request não seja devidamente concluída, retornando o devido erro.

Além disso, os endpoints podem retornar outros atributos específicos daquele endpoint, mas sempre
serão retornados estes dois atributos.

## Endpoints:

### /api/jogo

#### Retornar dados de uma sala

```http
GET /api/jogo/${idSala}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `idSala`      | `string` | **Required**. Id da sala a ser retornada |

Response:

```javascript
{
  "sala": object
}
```

#### Retornar dados de uma partida passada

```http
GET /api/jogo/analise/${codJogo}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `codJogo`      | `string` | **Required**. cod_jogo da partida no banco de dados |

Response:

```javascript
{
  "jogo": object
}
```

#### Criar nova sala

```http
POST /api/jogo/salas
```

Request body:

```javascript
{
  "socketId": String,
  "username": String, // opcional
  "tempo": Number     // em segundos
}
```

Exemplo:
```javascript
{
  "socketId": "aAbBcCdDeEfFgG",
  "username": "vidacalura",
  "tempo": 180
}
```

Response:
```javascript
{
  "sala": object
}
```

#### Entrar em uma sala

```http
POST /api/jogo/salas/entrar
```

Request body:

```javascript
{
  "salaId": Number,
  "socketId": String,
  "username": String, // opcional
  "tempo": Number     // em segundos
}
```

Exemplo:
```javascript
{
  "salaId": 100000
  "socketId": "aAbBcCdDeEfFgG",
  "username": "vidacalura",
  "tempo": 180
}
```

#### Fazer um movimento

```http
POST /api/jogo
```

Request body:

```javascript
{
  "salaId": Number,
  "socketId": String,
  "movimentos": []object
}
```

Exemplo:
```javascript
{
  "salaId": 100000,
  "socketId": "teste",
  "movimentos": [
    {
      "x": 0,
      "y": 0,
      "lado": "B"
    },
    {
      "x": 0,
      "y": 1,
      "lado": "B"
    }
  ]
}
```