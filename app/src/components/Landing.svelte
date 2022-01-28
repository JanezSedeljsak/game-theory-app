<script>
  import { Button } from "svelte-chota";
  export let visible;
  export let showModal;

  function pickModeAndMove(game) {
    visible = game;
    showModal = true;
  }
</script>

<div class="landing-container flex-container">
  <h1>
    <span class="cursive-text">Test yourself against</span><br />
    <b
      ><span class="text-primary">AI</span> <span class="cursive-text">&</span>
      <span class="text-primary">game theory</span> algorithms</b
    >
  </h1>
  <div>
    <Button primary on:click={() => pickModeAndMove("tictactoe")}
      >Tic-Tac-Toe</Button
    >
    <Button primary on:click={() => pickModeAndMove("connect4")} disabled
      >Connect 4</Button
    >
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
