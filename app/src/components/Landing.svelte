<script>
  import { Button } from "svelte-chota";
  export let visible;
</script>

<div class="landing-container flex-container">
  <h1>
    Test yourself against<br />
    <b>AI & game theory algorithms</b>
  </h1>
  <div>
    <Button primary on:click={() => visible = "tictactoe"}>TicTacToe</Button>
    <Button primary on:click={() => visible = "connect4"} disabled>Connect4</Button>
  </div>
</div>

<style>
  .landing-container {
    flex-direction: column;
  }

  .landing-container > h1 {
    text-align: center;
    font-size: 2.3em;
  }

  .landing-container > h1 > b {
    text-transform: uppercase;
  }
</style>
