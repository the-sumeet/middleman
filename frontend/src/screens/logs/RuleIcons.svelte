<script>
    export let log;
</script>
<div class="flex gap-2 mt-2">
    {#if log.cancelled}
      <i class="text-md text-yellow-500 bi bi-x-square" title="Cancelled"
      ></i>
    {/if}

    {#if log.redirected}
      <i class="text-md text-yellow-500 bi bi-shuffle" title="Redirected"
      ></i>
    {/if}

    {#if log.requestHeaderModified}
      <i class="text-md text-yellow-500 bi bi-h-square" title="Request Header Modified"
      ></i>
    {/if}

    {#if log.responseHeaderModified}
      <i class="text-md text-yellow-500 bi bi-h-square" title="Response Header Modified"
      ></i>
    {/if}

    {#if log.requestBodyModified}
      <i class="text-md text-yellow-500 bi bi-body-text" title="Request Body Modified"
      ></i>
    {/if}

    {#if log.responseBodyModified}
      <i class="text-md text-yellow-500 bi bi-body-text" title="Response Body Modified"
      ></i>
    {/if}

    {#if log.delayed != 0}
      <i class="text-md text-yellow-500 bi bi-clock" title="Delayed by {log.delayed} second(s)"
      ></i>
    {/if}

  </div>