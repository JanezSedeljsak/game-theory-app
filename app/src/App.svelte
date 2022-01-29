<script>
  import { ToastContainer, FlatToast } from "svelte-toasts";
  import TicTacToe from "./components/TicTacToe.svelte";
  import Connect4 from "./components/Connect4.svelte";
  import Landing from "./components/Landing.svelte";
  import ModeModal from "./components/ModeModal.svelte";
  import { GMEnum, GMEnumStr } from "./util";
  import "chota";

  let visible = "landing";
  let showModal = false;
  let gameMode = GMEnum.AdvancedAI;
</script>

<svelte:head>
	<title>Game Theory App - {GMEnumStr(gameMode)}</title>
</svelte:head>

<div class="{visible === 'landing' ? 'wrapper-background' : ''} wrapper flex-container ">
  <div class="{visible !== 'landing' ? 'force-top' : ''} container flex-container">
    {#if visible == "tictactoe"}
      <TicTacToe bind:visible bind:gameMode bind:showModal />
    {:else if visible == "connect4"}
      <Connect4 bind:visible bind:gameMode bind:showModal />
    {:else}
      <Landing bind:visible bind:showModal />
    {/if}
  </div>
  <ModeModal bind:visible={showModal} bind:gameMode  />
  <ToastContainer let:data>
    <FlatToast {data} />
  </ToastContainer>
</div>
