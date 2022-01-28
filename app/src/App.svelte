<script>
  import { onMount } from 'svelte';
  import { sleep, transSymbol } from './util';
  var board = null;

  onMount(async () => {
    board = await window.init()
    sleep();
	});

  async function move(row, col) {
    if (board[row][col] == 0) {
      board[row][col] = 1;
      const response = await window.mutate(board);
      const jsonResp = JSON.parse(response);
      board = jsonResp.board;
      await sleep();
      if (jsonResp.isdone) {
        if (jsonResp.winner == 0) {
          alert('Stalemate');
        } else {
          alert(`${jsonResp.winner == 1 ? 'O' : 'X'} won!!!`);
        }
        
        board = await window.init();
      }
    }
  }
</script>

<div class="grid-wrapper">
  {#if board}
  <div class="grid-container">
    {#each board as row, i}
      {#each row as column, j}
        <div class="grid-item" on:click={() => move(i, j)}>
          <div>{transSymbol(column)}</div>
        </div>
      {/each}
    {/each}
  </div>
  {/if}
</div>
