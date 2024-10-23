<script>
  export let log;

  console.log(log);

  function getUrl() {
    return log.scheme + "://" + log.host + log.path;
  }
</script>

<div class="mt-4">
  <div
    class="mt-4 flex flex-col bg-gray-800 p-4 rounded border-2 border-gray-700 items-start overflow-x-auto"
  >
    <!-- Datetime -->
    <p class="text-white text-xs">{new Date(log.timestamp)}</p>

    <div class="mt-2 flex space-x-4 justify-between">
      <!-- Right Side -->
      <div class="flex gap-2 items-center">
        <!-- Rule Icons -->
        {#if log.cancelled}
          <i class="text-2xl text-yellow-500 bi bi-x-square" title="Cancelled"
          ></i>
        {/if}

        <!-- Text -->
        <div class="gap-2 text-lg text-emerald-500">
          <div>{log.method}</div>
          <div class="">{getUrl() + getUrl() + getUrl()}</div>
        </div>
      </div>

      <!-- Expand / Collapse -->
      <!-- <i class="text-2xl text-white bi bi-chevron-down"></i> -->
      <!-- <i class="text-2xl text-gray-800 dark:text-white bi bi-chevron-up"></i> -->
    </div>
  </div>
</div>
