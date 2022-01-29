<script>
  import { sleep, getSymbol, GMEnum, getNextPlayer } from "../util";
  import { mdiHome, mdiRestart } from '@mdi/js';
  import { onMount } from "svelte";
  import { Button } from "svelte-chota";
  import { toasts } from "svelte-toasts";
  
  export let visible, gameMode;

  var board = null;
  var xStart = false;
  var winningLine = new Set();
  var playerMoveCount = 0;

  onMount(resetGame);
  async function resetGame() {
    xStart = !xStart;
    board = await window.init(
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
        response = await window.mutateAI(board);
        break;
      case GMEnum.EasyAI:
        response = await window.mutateRand(board);
        break;
      case GMEnum.Multiplayer:
        response = await window.multiplayer(board);
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
  <div class="stick-bottom-right">
    <Button primary class="is-rounded" icon={mdiHome} on:click={() => (visible = "landing")} />
    <Button primary class="is-rounded" icon={mdiRestart} on:click={resetGame} />
  </div>
</div>
