<img src="https://github.com/matiassingers/awesome-readme/blob/master/icon.png" align="right" />

# Game Theory App

The past year I've learned a lot about artificial intelligence, heuristic algorithms, crafting & optimization of algorithms. For this reason I've decided to test my skills and code a basic Tic-tac-toe solver...

The actual logic of the game-theory algorithms is written in [GO](https://go.dev) and on top of that I've added a user interface written in [Svelte](https://svelte.dev) and served via [Lorca](https://github.com/zserge/lorca) from [GO](https://go.dev).

Tic-tac-toe has a very limited space complexity (there are 3^9 = 19683 possible field layouts) which was easily handled by a basic [Minimax algorithm](https://en.wikipedia.org/wiki/Minimax), nevertheless I've optimized it with [Alpha-beta pruning](https://en.wikipedia.org/wiki/Alpha–beta_pruning) and state [Memoization](https://en.wikipedia.org/wiki/Memoization) with a simple [Dynamic programming](https://en.wikipedia.org/wiki/Dynamic_programming) approach.

![No image](https://github.com/JanezSedeljsak/game-theory-app/blob/main/docs/banner.png)

### Install locally

* Clone from git
```terminal
$ git clone https://github.com/JanezSedeljsak/game-theory-app.git
```
* Build client
```terminal
cd app
npm install
npm run build
cd ..
```

* Build Core
```
go get
go run .
```

### Prerequisites

* [GO](https://go.dev)
* [Node](https://nodejs.org/en/) > 14.0

### License

[![CC0](https://upload.wikimedia.org/wikipedia/commons/thumb/0/0c/MIT_logo.svg/220px-MIT_logo.svg.png)](https://en.wikipedia.org/wiki/MIT_License)

### Authors

```JS
const AUTHORS = ['Janez Sedeljsak']
```
