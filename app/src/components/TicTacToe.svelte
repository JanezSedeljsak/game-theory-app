<script>
  import { onMount } from "svelte";
  import { sleep, getSymbol } from "../util";
  import { toasts } from "svelte-toasts";

  export let visible;
  export let showModal;
  export let gameMode;

  var board = null;
  var aiStart = false;
  var winningLine = new Set();

  onMount(async () => {
    board = await window.init(aiStart);
  });

  async function move(row, col) {
    if (board[row][col] != 0) return;
    board[row][col] = 1;
    const response = await window.mutate(board);
    const jsonResp = JSON.parse(response);
    board = jsonResp.board;
    await sleep();

    if (!jsonResp.isdone) return;
    const description = !jsonResp.winner
      ? "Stalemate"
      : `${jsonResp.winner == 1 ? "O" : "X"} won!!!`;

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

<div class="grid-wrapper">
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
</div>

<style>
  .grid-wrapper {
    height: 90%;
    width: 90%;
    background-color: #ddd;
  }

  .grid-container {
    height: 100% !important;
    width: auto;
    display: grid;
    grid-template-columns: repeat(3, minmax(160px, 1fr));
    grid-auto-rows: 1fr;
    grid-gap: 1px;
  }

  .grid-item {
    background-color: #444;
  }

  .grid-item > div {
    font-family: "Architects Daughter", cursive;
    font-size: 6em;
  }
</style>
