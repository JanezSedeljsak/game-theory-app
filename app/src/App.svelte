<script>
  import { onMount } from 'svelte';
  import { sleep, transSymbol } from './util';
  import { toasts, ToastContainer, FlatToast }  from "svelte-toasts";

  var board = null;
  var aiStart = true;
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
    const description = !jsonResp.winner ? 'Stalemate' : (`${jsonResp.winner == 1 ? 'O' : 'X'} won!!!`);
    toasts.add({
      title: 'Game finished',
      description: description,
      type: 'info',
      onRemove: async () => {
        board = await window.init(aiStart)
      }
    });
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
  <ToastContainer let:data={data}>
		<FlatToast {data}  />
	</ToastContainer>
</div>
