<script>
  import { errorMessage, selectedRequest } from "../../../src/stores";
  import { GetLogs } from "../../../wailsjs/go/main/App";
  import { scrollToBottomIfAtBottom } from "../../../src/utils";
  import RequestContent from "./RequestContent.svelte";
  export const tabSelectedStyle =
    "border-blue-400 text-blue-300  focus:outline-none";
  export const tabUnselectedStyle =
    "border-transparent text-white cursor-base focus:outline-none hover:border-gray-400";
  import { formatDate } from "../../../src/utils";
  import MethodTableCol from "./MethodTableCol.svelte";
  import StatusTableCol from "./StatusTableCol.svelte";
  import { GetOneRequest } from "../../../wailsjs/go/main/App";
  let logs = [];
  let clear;


  fetchLogs();

  function selectRequest(log) {
    GetOneRequest(log.id).then((res) => {
      if (res.error != "") {
        errorMessage.set(res.error);
      } else {
        selectedRequest.set(res.httpRequests[0]);
      }
    });
  }

  function getUrl(log) {
    return `${log.scheme}://${log.host}${log.path}`;
  }


  function fetchLogs() {
    GetLogs(logs.length).then((res) => {
      if (res.error != "") {
        errorMessage.set(res.error);
      } else if (res.httpRequests) {
        logs = [...logs, ...res.httpRequests];
        scrollToBottomIfAtBottom(document.getElementById("logs"));
      }
    });
  }

  $: {
    clearInterval(clear);
    clear = setInterval(fetchLogs, 1000);
  }
</script>

<div class="flex flex-col h-screen p-4 w-full">
  <h2 class="text-2xl font-medium text-white">HTTP(S) Requests</h2>

  <!-- <div id="logs" class="flex flex-col overflow-y-auto border-gray-700">
    {#each logs as log}
      <Log {log} />
    {/each}
  </div> -->

  <!--
  Heads up! 👋

  This component comes with some `rtl` classes. Please remove them if they are not needed in your project.
-->

  <div
    id="logs"
    class="text-white overflow-y-auto overflow-x-auto max-w-full mt-2 rounded border border-gray-700"
  >
    <table class="table-fixed w-full">
      <tbody>
        {#each logs as log}
          <tr
            on:click={() => selectRequest(log)}
            class=" odd:bg-gray-800/50 hover:bg-blue-500/30"
          >
            <td title="{log.timestamp}" class="text-sm w-1/12 truncate border-r border-gray-700 px-1"
              >{formatDate(log.timestamp)}</td
            >
            <MethodTableCol method={log.method} />
            <td class="text-sm w-1/12 truncate border-r border-gray-700 px-1"
              >{"application/json"}</td
            >
            <td class="text-sm w-7/12 truncate border-r border-gray-700 px-1"
              >{getUrl(log)}</td
            >
            <StatusTableCol status={log.status} />
            <td class="text-sm w-1/12 truncate">RULES</td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>

  <!-- Request content -->
  {#if $selectedRequest}
    <RequestContent />
  {/if}
</div>
