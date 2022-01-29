<script>
    import { sleep, getSymbol, GMEnum, getNextPlayer } from "../util";
    import { mdiHome, mdiRestart, mdiGamepad } from '@mdi/js';
    import { Button } from "svelte-chota";
    import { toasts } from "svelte-toasts";
    
    export let visible, gameMode, showModal;
  
    var board = null;
    var xStart = false;
    var winningLine = new Set();
    var playerMoveCount = 0;
  </script>
  
  <div class="grid-wrapper">
    {#if board}
      <div class="grid-container">
        {#each board as row, i}
          {#each row as column, j}
            <div
              class="grid-item flex-container full-size"
            >
              <div class={winningLine.has(i * 3 + j) ? "text-primary" : ""}>
                {getSymbol(column)}
              </div>
            </div>
          {/each}
        {/each}
      </div>
    {/if}
    <div class="stick-bottom-right">
      <Button primary class="is-rounded" icon={mdiHome} on:click={() => (visible = "landing")} />
      <Button primary class="is-rounded" icon={mdiRestart} on:click={() => {}} />
      <Button primary class="is-rounded" icon={mdiGamepad} on:click={() => (showModal = true)} />
    </div>
  </div>
  