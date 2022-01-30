<script>
  import { Button } from "svelte-chota";
  import { mdiGamepad } from "@mdi/js";
  export let visible, showModal;
</script>

<div class="landing-container flex-container">
  <h1>
    <span class="cursive-text">Test yourself against</span><br />
    <b>
      <span class="text-primary">AI</span>
      <span class="cursive-text">&</span>
      <span class="text-primary">game theory</span>
      algorithms
    </b>
  </h1>
  <div>
    <Button primary on:click={() => (visible = "tictactoe")}>Tic-Tac-Toe</Button>
    <Button primary on:click={() => (visible = "connect4")}>Connect 4</Button>
  </div>
  <Button outline on:click={() => (showModal = true)} icon={mdiGamepad}>Gamemode options</Button>
</div>
