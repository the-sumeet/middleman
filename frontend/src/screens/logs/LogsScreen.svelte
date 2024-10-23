<script>
  import { GetLogs } from "../../../wailsjs/go/main/App";
  import Log from "./Log.svelte";
  import { scrollToBottomIfAtBottom } from "../../../src/utils";

  let logs = [];

  fetchLogs();

  function fetchLogs() {
    GetLogs().then((res) => {
      if (res.httpRequests) {
        logs = res.httpRequests;
        scrollToBottomIfAtBottom(document.getElementById("logs"));
      }
    });
  }

  let clear;
  $: {
    clearInterval(clear);
    clear = setInterval(fetchLogs, 1000);
  }
</script>

<div class="flex flex-col h-screen p-4 w-full">
  <h2 class="text-2xl font-medium text-white">Logs</h2>

  <div id="logs" class="flex flex-col overflow-y-auto border-gray-700">
    {#each logs as log}
      <Log {log} />
    {/each}
  </div>
</div>
