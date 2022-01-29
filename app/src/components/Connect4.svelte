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
    console.log(board);
    winningLine = new Set();
    playerMoveCount = 0;
  }
</script>

<div class="grid-wrapper grid-wrapper-cf">
  {#if board}
    <div class="grid-container">
      {#each board as row, i}
        {#each row as _, j}
          <div class="grid-item flex-container full-size">
              {#if board[5 - i][j] != 0}
                <div class="{board[5 - i][j] == 1 ? 'circle-first' : 'circle-second'} circle" />
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
    <Button primary class="is-rounded" icon={mdiRestart} on:click={() => {}} />
    <Button
      primary
      class="is-rounded"
      icon={mdiGamepad}
      on:click={() => (showModal = true)}
    />
  </div>
</div>
