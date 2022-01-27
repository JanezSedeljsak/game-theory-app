<script>
  const board = [
    [-1, 1, 1],
    [0, 1, 0],
    [0, -1, -1],
  ];

  function transSymbol(symbol) {
    switch (symbol) {
      case 0:
        return "";
      case 1:
        return "O";
      case -1:
        return "X";
    }
  }

  function move(row, col) {
    if (board[row][col] == 0) {
      board[row][col] = 1;
    }
  }
</script>

<div class="grid-wrapper">
  <div class="grid-container">
    {#each board as row, i}
      {#each row as column, j}
        <div class="grid-item" on:click={() => move(i, j)}>
          <div>{transSymbol(column)}</div>
        </div>
      {/each}
    {/each}
  </div>
</div>
