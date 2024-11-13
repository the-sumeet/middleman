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
  import Dropdown from "./Dropdown.svelte";

  let logs = [];
  let clear;

  let filterUrl = "";
  let filterStatus = "";
  let filterStatusError;
  let filterMethods = [];
  let filterResponseTypes = [];
  let filterAppliedRules = [];

  fetchLogs();

  function statusesNumbers() {
    if (filterStatus == "") {
      return [];
    }
    return filterStatus.split(",");
  }

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
    let statuses = statusesNumbers();

    GetLogs(
      filterUrl,
      filterMethods,
      statuses,
      filterResponseTypes,
      filterAppliedRules,
      logs.length,
    ).then((res) => {
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

  <!-- Filters -->
  <div class="flex gap-2 mt-4">
    <!-- URL Search -->
    <div class="w-4/12">
      <div class="relative flex items-center">
        <span class="absolute">
          <i class="w-6 h-6 mx-3 text-gray-500 bi bi-search"></i>
        </span>

        <input
          bind:value={filterUrl}
          on:input={() => {
            logs = [];
            fetchLogs();
            console.log(filterUrl);
          }}
          type="text"
          placeholder="URL contains"
          class="block w-full py-1 text-gray-700 placeholder-gray-400/70 bg-white border border-gray-200 rounded-lg pl-8 pr-5 rtl:pr-11 rtl:pl-5 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-600 focus:border-blue-400 dark:focus:border-blue-300 focus:ring-blue-300 focus:outline-none focus:ring focus:ring-opacity-40"
        />
      </div>
    </div>

    <!-- Status -->
    <div class="w-2/12">
      <div class="relative flex items-center">
        <input
          bind:value={filterStatus}
          on:input={() => {
            const statuses = statusesNumbers();
            logs = [];
            fetchLogs();
          }}
          type="text"
          placeholder="Status starting with"
          class="block w-full py-1 text-gray-700 placeholder-gray-400/70 bg-white border border-gray-200 rounded-lg pl-2 pr-5 rtl:pr-11 rtl:pl-5 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-600 focus:border-blue-400 dark:focus:border-blue-300 focus:ring-blue-300 focus:outline-none focus:ring focus:ring-opacity-40"
        />
      </div>
      <!-- Statuses error -->
      <!-- svelte-ignore a11y-label-has-associated-control -->
      {#if filterStatusError}
        <label class="text-red-400 mt-1 block text-xs"
          >{filterStatusError}</label
        >
      {/if}
    </div>

    <div class="w-2/12">
      <Dropdown
        title="Methods"
        onChange={(selected) => {
          filterMethods = selected;
          logs = [];
          fetchLogs();
        }}
        values={["GET", "POST", "PUT", "PATCH", "DELETE"]}
      />
    </div>

    <div class="w-2/12">
      <Dropdown
        title="Response Type"
        onChange={(selected) => {
          filterResponseTypes = selected;
          logs = [];
          fetchLogs();
        }}
        values={["HTML", "JSON"]}
      />
    </div>

    <div class="w-2/12">
      <Dropdown
        title="Applied Rules"
        onChange={(selected) => {
          filterAppliedRules = selected;
          logs = [];
          fetchLogs();
        }}
        values={[
          "cancel",
          "redirect",
          "modifyRequestHeader",
          "modifyResponseHeader",
          "modifyRequestBody",
          "modifyResponseBody",
          "delay",
        ]}
      />
    </div>
  </div>

  <div
    id="logs"
    class="text-white overflow-y-auto overflow-x-auto max-w-full mt-2 rounded border border-gray-700"
  >
    <table class="table-fixed w-full">
      <tbody>
        {#each logs as log, i (log.id)}
          <tr
            on:click={() => selectRequest(log)}
            class=" odd:bg-gray-800/50 hover:bg-blue-500/30"
          >
            <td
              title={log.timestamp}
              class="text-sm w-1/12 truncate border-r border-gray-700 px-1"
              >{formatDate(log.timestamp)}</td
            >
            <MethodTableCol method={log.method} />

            <td class="text-sm w-1/12 truncate border-r border-gray-700 px-1"
              >{"application/json"}</td
            >
            <td class="text-sm flex truncate border-r border-gray-700 px-1"
              >{getUrl(log)}</td
            >
            <StatusTableCol status={log.status} />

            <td class="text-sm w-36 truncate">
              <div class="flex justify-between">
                <!-- Delay, Cancelled -->
                {#if log.delayed > 0 || log.cancelled || log.redirected}
                  <div class="flex gap-2">
                    {#if log.delayed > 0}
                      <i
                        class="bi bi-clock"
                        title="Request Delayed by {log.delayed}s"
                      ></i>
                    {/if}
                    {#if log.cancelled}
                      <i class="bi bi-x-square" title="Request Cancelled"></i>
                    {/if}
                    {#if log.redirected}
                      <i class="bi bi-shuffle" title="Request Redirected"></i>
                    {/if}
                  </div>
                {/if}

                <!-- Request header, body -->
                {#if log.requestHeaderModified || log.requestBodyModified}
                  <div class="flex gap-2 text-green-400">
                    {#if log.requestBodyModified}
                      <i class="bi bi-body-text" title="Request Body Modified"
                      ></i>
                    {/if}
                    {#if log.requestHeaderModified}
                      <i class="bi bi-h-square" title="Request Headers Modified"
                      ></i>
                    {/if}
                  </div>
                {/if}

                <!-- Response header, body -->
                {#if log.responseHeaderModified || log.responseBodyModified}
                  <div class="flex gap-2 text-yellow-400">
                    {#if log.responseBodyModified}
                      <i class="bi bi-body-text" title="Response Body Modified"
                      ></i>
                    {/if}
                    {#if log.responseHeaderModified}
                      <i
                        class="bi bi-h-square"
                        title="Response Headers Modified"
                      ></i>
                    {/if}
                  </div>
                {/if}
              </div>
            </td>
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
