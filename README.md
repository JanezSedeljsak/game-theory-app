<img src="https://github.com/matiassingers/awesome-readme/blob/master/icon.png" align="right" />

# Game Theory App

The past year I've learned a lot about artificial intelligence, heuristic algorithms, crafting & optimization of algorithms. For this reason I've decided to test my skills and code a basic Tic-tac-toe solver...

The actual logic for the game-theory algorithm is written in [GO](https://go.dev) and on top of that I've built a user interface written in [Svelte](https://svelte.dev) ~ served via [Lorca](https://github.com/zserge/lorca) from [GO](https://go.dev).

Tic-tac-toe has a very limited space complexity (there are 3^9 = 19683 possible layouts) which was easily handled by a basic [Minimax algorithm](https://en.wikipedia.org/wiki/Minimax). Nevertheless, I decided to implement [Alpha-beta pruning](https://en.wikipedia.org/wiki/Alpha–beta_pruning) and state [Memoization](https://en.wikipedia.org/wiki/Memoization) with a simple [Dynamic programming](https://en.wikipedia.org/wiki/Dynamic_programming) approach of caching solutions via hash values of the current layout.

## Connect 4

On the other hand connect 4 was a whole different problem, here the space complexity is much much bigger, there are more than 2^42 possible layouts, which is more than 500GB of data if you wanted to store that (and that is only taking into consideration storing the layouts nevermind the solution)...

In its current state the algorithm holds up quite well against the AI on: [coonect4.gamesolver](https://connect4.gamesolver.org).<br/>
Improvements made to the miniMax algorithm:
* [Negamax](https://en.wikipedia.org/wiki/Negamax)
* [Alpha-beta pruning](https://en.wikipedia.org/wiki/Alpha–beta_pruning)
* [Memoization](https://en.wikipedia.org/wiki/Memoization) with progressive deepening
* [Bitmap board](https://github.com/denkspuren/BitboardC4/blob/master/BitboardDesign.md)

#### Possible improvements in the future
* A good heuristic evaluation
* Hard coded dictionary `HashMap<"Unique Bitmap", "Best move">`

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

### Credits
* [Optimization ideas](http://blog.gamesolver.org)
* [Optimization ideas #2](https://towardsdatascience.com/creating-the-perfect-connect-four-ai-bot-c165115557b0)
* [Bitmap representation](https://github.com/denkspuren/BitboardC4)
