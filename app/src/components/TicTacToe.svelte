<script>
  import { getSymbol, GMEnum, getNextPlayer } from "../util";
  import { mdiHome, mdiRestart, mdiGamepad } from '@mdi/js';
  import { Button } from "svelte-chota";
  import { toasts } from "svelte-toasts";
  
  const {ttt_init, ttt_multiplayer, ttt_mutateAI, ttt_mutateRand} = window;
  export let visible, gameMode, showModal;

  var board = null;
  var xStart = false;
  var winningLine = new Set();
  var playerMoveCount = 0;

  // reset game on modal dispatch
  $: (() => !showModal && resetGame())()

  async function resetGame() {
    xStart = !xStart;
    board = await ttt_init(
      xStart && gameMode != GMEnum.Multiplayer,
      gameMode != GMEnum.AdvancedAI
    );
    winningLine = new Set();
    playerMoveCount = 0;
  }

  async function move(row, col) {
    if (board[row][col] != 0 || winningLine.size != 0) return;
    board[row][col] = getNextPlayer(xStart, playerMoveCount, gameMode);
    var response;
    switch (gameMode) {
      case GMEnum.AdvancedAI:
        response = await ttt_mutateAI(board);
        break;
      case GMEnum.EasyAI:
        response = await ttt_mutateRand(board);
        break;
      case GMEnum.Multiplayer:
        response = await ttt_multiplayer(board);
        break;
      default:
        throw new Error(`Invalid gameMode enum - ${gameMode}!`);
    }

    playerMoveCount += gameMode != GMEnum.Multiplayer ? 2 : 1;
    const jsonResp = JSON.parse(response);
    board = jsonResp.board;
    evalResult(jsonResp);

  }

  async function evalResult(jsonResp) {
    if (!jsonResp.isdone) return;
    const description = !jsonResp.winner
      ? "Stalemate"
      : `${jsonResp.winner == 1 ? "O" : "X"} won!`;

    if (jsonResp.winner != 0) {
      winningLine = new Set(
        jsonResp.coords.map((coord) => coord.Row * 3 + coord.Col)
      );
    }

    toasts.add({
      title: "Game finished",
      description: description,
      type: "info",
    });
  }
</script>

<div class="grid-wrapper grid-wrapper-ttt">
  {#if board}
    <div class="grid-container">
      {#each board as row, i}
        {#each row as column, j}
          <div
            class="grid-item flex-container full-size"
            on:click={() => move(i, j)}
          >
            <div class={winningLine.has(i * 3 + j) ? "text-primary" : ""}>
              {getSymbol(column)}
            </div>
          </div>
        {/each}
      {/each}
    </div>
  {/if}
  <div class="stick-bottom-right">
    <Button primary class="is-rounded" icon={mdiHome} on:click={() => (visible = "landing")} />
    <Button primary class="is-rounded" icon={mdiRestart} on:click={resetGame} />
    <Button primary class="is-rounded" icon={mdiGamepad} on:click={() => (showModal = true)} />
  </div>
</div>
