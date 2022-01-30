<script>
  import { getNextRow, GMEnum, getNextPlayer } from "../util";
  import { mdiHome, mdiRestart, mdiGamepad } from "@mdi/js";
  import { Button } from "svelte-chota";
  import { toasts } from "svelte-toasts";

  const { cf_init, cf_multiplayer, cf_mutateAI, cf_mutateRand } = window;
  export let visible, gameMode, showModal;

  var board = null;
  var xStart = false;
  var winningLine = new Set();
  var playerMoveCount = 0;

  // reset game on modal dispatch
  $: (() => !showModal && resetGame())();

  async function resetGame() {
    xStart = !xStart;
    board = await cf_init(
      xStart && gameMode != GMEnum.Multiplayer,
      gameMode != GMEnum.AdvancedAI
    );
    winningLine = new Set();
    playerMoveCount = 0;
  }

  async function move(col) {
    if (board[5][col] != 0 || winningLine.size != 0) return;
    const nextRow = getNextRow(board, col);
    board[nextRow][col] = getNextPlayer(xStart, playerMoveCount, gameMode);

    var response;
    switch (gameMode) {
      case GMEnum.AdvancedAI:
        response = await cf_mutateAI(board, nextRow, col);
        break;
      case GMEnum.EasyAI:
        response = await cf_mutateRand(board, nextRow, col);
        break;
      case GMEnum.Multiplayer:
        response = await cf_multiplayer(board, nextRow, col);
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
      : `${jsonResp.winner == 1 ? "Green" : "White"} won!`;

    if (jsonResp.winner != 0) {
      winningLine = new Set(
        jsonResp.coords.map((coord) => coord.Row * 7 + coord.Col)
      );
    }

    toasts.add({
      title: "Game finished",
      description: description,
      type: "info",
    });
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
                  : 'circle-second'} circle full-size 
                  {winningLine.has((5 - i) * 7 + j) ? 'circle-border' : ''}"
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
