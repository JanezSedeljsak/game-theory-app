<script>
  import { ToastContainer, FlatToast } from "svelte-toasts";
  import TicTacToe from "./components/TicTacToe.svelte";
  import Connect4 from "./components/Connect4.svelte";
  import Landing from "./components/Landing.svelte";
  import ModeModal from "./components/ModeModal.svelte";
  import { GetGMEnum } from "./util";
  import "chota";

  const GMEnum = GetGMEnum();

  let visible = "landing";
  let showModal = false;
  let gameMode = GMEnum.AdvancedAI;
</script>

<div class="{visible === 'landing' ? 'wrapper-background' : ''} wrapper flex-container ">
  <div class="container flex-container">
    {#if visible == "tictactoe"}
      <TicTacToe bind:visible bind:showModal />
    {:else if visible == "connect4"}
      <Connect4 bind:visible />
    {:else}
      <Landing bind:visible bind:showModal />
    {/if}
  </div>
  <ModeModal bind:visible={showModal} bind:gameMode  />
  <ToastContainer let:data>
    <FlatToast {data} />
  </ToastContainer>
</div>

<style>
  .wrapper-background {
    background-image: url('../assets/landing.jpg') !important;
    background-size: cover;
    background-repeat: no-repeat;
  }
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
