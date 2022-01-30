<script>
  import { sleep, getSymbol, GMEnum, getNextPlayer } from "../util";
  import { mdiHome, mdiRestart, mdiGamepad } from "@mdi/js";
  import { Button } from "svelte-chota";
  import { toasts } from "svelte-toasts";

  export let visible, gameMode, showModal;

  var board = null;
  var xStart = false;
  var winningLine = new Set();
  var playerMoveCount = 0;

  // reset game on modal dispatch
  $: (() => !showModal && resetGame())();

  async function resetGame() {
    xStart = !xStart;
    board = await window.cf_init(
      xStart && gameMode != GMEnum.Multiplayer,
      gameMode != GMEnum.AdvancedAI
    );
    winningLine = new Set();
    playerMoveCount = 0;
  }

  async function move(col) {
    if (board[5][col] != 0) return;
    const player = getNextPlayer(xStart, playerMoveCount, gameMode);
    board = await window.cf_playerDrop(col, player);
    var response;
    switch (gameMode) {
      case GMEnum.AdvancedAI:
        response = await window.cf_mutateAI();
        break;
      case GMEnum.EasyAI:
        response = await window.cf_mutateRand();
        break;
      case GMEnum.Multiplayer:
        response = await window.cf_multiplayer();
        break;
      default:
        throw new Error(`Invalid gameMode enum - ${gameMode}!`);
    }

    playerMoveCount += gameMode != GMEnum.Multiplayer ? 2 : 1;
    board = response;
  }
</script>

<div class="grid-wrapper grid-wrapper-cf">
  {#if board}
    <div class="grid-container">
      {#each board as row, i}
        {#each row as _, j}
          <div
            class="grid-item flex-container full-size"
            on:click={() => move(j)}
          >
            {#if board[5 - i][j] != 0}
              <div
                class="{board[5 - i][j] == 1
                  ? 'circle-first'
                  : 'circle-second'} circle"
              />
            {/if}
          </div>
        {/each}
      {/each}
    </div>
  {/if}
  <div class="stick-bottom-right">
    <Button
      primary
      class="is-rounded"
      icon={mdiHome}
      on:click={() => (visible = "landing")}
    />
    <Button primary class="is-rounded" icon={mdiRestart} on:click={resetGame} />
    <Button
      primary
      class="is-rounded"
      icon={mdiGamepad}
      on:click={() => (showModal = true)}
    />
  </div>
</div>
