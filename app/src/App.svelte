<script>
  import { ToastContainer, FlatToast } from "svelte-toasts";
  import TicTacToe from "./components/TicTacToe.svelte";
  import Connect4 from "./components/Connect4.svelte";
  import Landing from "./components/Landing.svelte";
  import "chota";

  let visible = "landing";
</script>

<div class="wrapper flex-container">
  <div class="container flex-container">
    {#if visible == "tictactoe"}
      <TicTacToe bind:visible={visible} />
    {:else if visible == "connect4"}
      <Connect4 bind:visible={visible} />
    {:else}
      <Landing bind:visible={visible} />
    {/if}
  </div>
  <ToastContainer let:data>
    <FlatToast {data} />
  </ToastContainer>
</div>

<style>
  .wrapper {
    width: 100vw;
    height: 100vh;
    background-color: #444;
  }

  .container {
    width: 90%;
    height: 90%;
    padding: 20px;
  }
</style>
