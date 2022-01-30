<script>
  import { getSymbol, GMEnum, GEnum, getNextPlayer, evalGameStatus, boardAction } from "../util";
  import { mdiHome, mdiRestart, mdiGamepad } from "@mdi/js";
  import { Button } from "svelte-chota";
  import { toasts } from "svelte-toasts";

  const { ttt_init: init, ttt_multiplayer, ttt_mutateAI, ttt_mutateRand } = window;
  export let visible, gameMode, showModal;

  var board = null;
  var xStart = false;
  var winningLine = new Set();
  var playerMoveCount = 0;

  // reset game on modal dispatch
  $: (() => !showModal && resetGame())();

  async function resetGame() {
    xStart = !xStart;
    board = await init(xStart && gameMode != GMEnum.Multiplayer, gameMode != GMEnum.AdvancedAI);
    winningLine = new Set();
    playerMoveCount = 0;
  }

  async function move(row, col) {
    if (board[row][col] != 0 || winningLine.size != 0) return;
    board[row][col] = getNextPlayer(xStart, playerMoveCount, gameMode);

    const response = await boardAction(gameMode, ttt_mutateAI, ttt_mutateRand, ttt_multiplayer, board);
    playerMoveCount += gameMode != GMEnum.Multiplayer ? 2 : 1;
    [board, winningLine] = await evalGameStatus(response, toasts, GEnum.TicTacToe);
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
