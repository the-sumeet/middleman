<script>
  import { GetLogs } from "../../../wailsjs/go/main/App";
  import Log from "./Log.svelte";
  let logs = [];

  GetLogs().then((res) => {
    if (res.httpRequests) {
      logs = res.httpRequests;
    }
  });
</script>

<div class="flex flex-col h-screen mx-auto p-4 w-full">
  <h2 class="text-2xl font-medium text-white">Logs</h2>

  <div class="flex flex-col overflow-y-auto border-gray-700">
    {#each logs as log}
      <Log {log} />
    {/each}
  </div>
</div>
