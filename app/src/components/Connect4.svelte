<script>
  import { getNextRow, GMEnum, GEnum, getNextPlayer, evalGameStatus, boardAction, sleep } from "../util";
  import { mdiHome, mdiRestart, mdiGamepad } from "@mdi/js";
  import { Button } from "svelte-chota";
  import { toasts } from "svelte-toasts";

  const { cf_init: init, cf_multiplayer, cf_mutateAI, cf_mutateRand } = window;
  export let visible, gameMode, showModal;

  var board = null;
  var xStart = false;
  var winningLine = new Set();
  var playerMoveCount = 0;
  var hoverCol = -1;

  // reset game on modal dispatch
  $: (() => !showModal && resetGame())();

  async function resetGame() {
    xStart = !xStart;
    board = await init(xStart && gameMode != GMEnum.Multiplayer, gameMode == GMEnum.AdvancedAI);
    winningLine = new Set();
    playerMoveCount = 0;
  }

  async function move(col) {
    if (board[5][col] != 0 || winningLine.size != 0) return;
    const nextRow = getNextRow(board, col);
    board[nextRow][col] = getNextPlayer(xStart, playerMoveCount, gameMode);

    await sleep();
    const response = await boardAction(gameMode, cf_mutateAI, cf_mutateRand, cf_multiplayer, board, col);
    playerMoveCount += gameMode != GMEnum.Multiplayer ? 2 : 1;
    [board, winningLine] = await evalGameStatus(response, toasts, GEnum.Connect4);
    if (winningLine.size > 0) {
      hoverCol = -1;
    }
  }

  function setHover(col) {
    if (winningLine.size == 0) {
      hoverCol = col;
    }
  }
</script>

<div class="grid-wrapper grid-wrapper-cf">
  {#if board}
    <div class="grid-container">
      {#each board as row, i}
        {#each row as _, j}
          <!-- svelte-ignore a11y-mouse-events-have-key-events -->
          <div 
            class="flex-container full-size"
            on:mouseover={() => setHover(j)}
            on:mouseout={() => hoverCol = -1}
            on:click={() => move(j)}
            style="padding: 5px"
          >
            <div
              class="{j == hoverCol ? 'hover-column' : ''} grid-item flex-container full-size">
              {#if board[5 - i][j] != 0}
                <div
                  class="{board[5 - i][j] == 1
                    ? 'circle-first'
                    : 'circle-second'} circle full-size 
                    {winningLine.has((5 - i) * 7 + j) ? 'circle-border' : ''}"
                />
              {/if}
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
