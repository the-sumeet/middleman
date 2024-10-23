<script>
  export let log;

  import RuleIcons from "./RuleIcons.svelte";

  console.log(log);

  function getUrl() {
    return `${log.scheme}://${log.host}${log.path}`;
  }
</script>

<div class="mt-4">
  <div
    class="flex flex-col bg-gray-800 p-2 rounded border-2 border-gray-700 items-start overflow-x-auto"
  >
    <!-- Datetime -->
    <p class="text-white text-xs">{new Date(log.timestamp)}</p>

    {#if log.cancelled || log.redirected || log.requestHeaderModified || log.responseHeaderModified || log.requestBodyModified || log.responseBodyModified || log.delayed}
      <RuleIcons {log} />
    {/if}

    <div class="mt-2 flex space-x-4 justify-between">
      <!-- Right Side -->
      <div class="flex gap-2 items-center">
        <!-- Text -->
        <div class="flex gap-2 text-md text-emerald-500">
          <div>{log.method}</div>
          <div class="break-all">{getUrl()}</div>
        </div>
      </div>

      <!-- Expand / Collapse -->
      <!-- <i class="text-2xl text-white bi bi-chevron-down"></i> -->
      <!-- <i class="text-2xl text-gray-800 dark:text-white bi bi-chevron-up"></i> -->
    </div>
  </div>
</div>
