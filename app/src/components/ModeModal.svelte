<script>
  import { Modal, Button, Card } from "svelte-chota";
  import { GMEnum } from "../util";
  export let visible, gameMode;

  function pickAndClose(option) {
    gameMode = option;
    visible = false;
  }
</script>

<Modal bind:open={visible}>
  <Card>
    <h4 slot="header" class="text-primary">Pick gamemode</h4>
    <p>You can either play against a friend or some version of the game AI?</p>

    <div slot="footer" class="flex-container">
      <Button primary on:click={() => pickAndClose(GMEnum.Multiplayer)}>Multiplayer</Button>
      <Button primary on:click={() => pickAndClose(GMEnum.EasyAI)}>EasyAI</Button>
      <Button primary on:click={() => pickAndClose(GMEnum.AdvancedAI)}>AdvancedAI</Button>
    </div>
  </Card>
</Modal>
