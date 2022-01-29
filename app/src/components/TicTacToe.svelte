<script>
  import { onMount } from "svelte";
  import { sleep, getSymbol } from "../util";
  import { toasts } from "svelte-toasts";
  import { GMEnum } from "../util";
  export let visible, gameMode;

  var board = null;
  var aiStart = false;
  var winningLine = new Set();

  onMount(async () => {
    board = await window.init(aiStart);
  });

  async function move(row, col) {
    if (board[row][col] != 0) return;
    board[row][col] = 1;
    var response;
    switch (gameMode) {
      case GMEnum.AdvancedAI:
        response = await window.mutateAI(board);
        break;
      case GMEnum.EasyAI:
        response = await window.mutateRand(board);
        break;
      case GMEnum.Multiplayer:
        break;
      default:
        throw new Error(`Invalid gameMode enum - ${gameMode}!`);
    }

    const jsonResp = JSON.parse(response);
    board = jsonResp.board;
    evalResult(jsonResp);
  }

  async function evalResult(jsonResp) {
    if (!jsonResp.isdone) return;
    await sleep();

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
