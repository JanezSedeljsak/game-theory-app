<script>
  import { getSymbol, GMEnum, GEnum, getNextPlayer, evalGameStatus, boardAction, sleep } from "../util";
  import { mdiHome, mdiRestart, mdiGamepad } from "@mdi/js";
  import { Button } from "svelte-chota";
  import { toasts } from "svelte-toasts";

  const { ttt_init: init, ttt_multiplayer, ttt_mutateAI, ttt_mutateRand } = window;
  export let visible, gameMode, showModal;

  var board = null;
  var xStart = false;
  var winningLine = new Set();
  var playerMoveCount = 0;
  var nextPlayer = null;
  var [hoverRow, hoverCol] = [-1, -1];

  // reset game on modal dispatch
  $: (() => !showModal && resetGame())();

  async function resetGame() {
    xStart = !xStart;
    board = await init(xStart && gameMode != GMEnum.Multiplayer, gameMode == GMEnum.AdvancedAI);
    winningLine = new Set();
    playerMoveCount = 0;
    nextPlayer = getNextPlayer(xStart, playerMoveCount, gameMode);
  }

  async function move(row, col) {
    if (board[row][col] != 0 || winningLine.size != 0) return;
    board[row][col] = nextPlayer

    await sleep();
    const response = await boardAction(gameMode, ttt_mutateAI, ttt_mutateRand, ttt_multiplayer, board);
    [board, winningLine] = await evalGameStatus(response, toasts, GEnum.TicTacToe);
    if (winningLine.size > 0) {
      [hoverRow, hoverCol] = [-1, -1];
    }

    playerMoveCount += gameMode != GMEnum.Multiplayer ? 2 : 1;
    nextPlayer = getNextPlayer(xStart, playerMoveCount, gameMode);
  }

  function setHover(row, col) {
    if (winningLine.size == 0) {
      [hoverRow, hoverCol] = [row, col];
    }
  }
</script>

<div class="grid-wrapper grid-wrapper-ttt">
  {#if board}
    <div class="grid-container">
      {#each board as row, i}
        {#each row as cell, j}
          <!-- svelte-ignore a11y-mouse-events-have-key-events -->
          <div
            class="grid-item flex-container full-size"
            on:click={() => move(i, j)}
            on:mouseover={() => setHover(i, j)}
            on:mouseout={() => [hoverRow, hoverCol] = [-1, -1]}
          >
            {#if cell != 0}
              <div class={winningLine.has(i * 3 + j) ? "text-primary" : ""}>
                {getSymbol(cell)}
              </div>
            {:else if i == hoverRow && j == hoverCol}
                {getSymbol(nextPlayer)}
            {/if}
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
