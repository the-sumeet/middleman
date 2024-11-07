<script>
    export let status;
</script>


{#if status >= 400}
<td class="font-bold text-red-400 text-sm w-1/12 truncate border-r border-gray-700 px-1">{status}</td>
{:else if status >= 300}
<td class="font-bold text-yellow-400 text-sm w-1/12 truncate border-r border-gray-700 px-1">{status}</td>
{:else if status >= 200}
<td class="font-bold text-green-400 text-sm w-1/12 truncate border-r border-gray-700 px-1">{status}</td>
{:else}
<td class="font-bold text-gray-300 text-sm w-1/12 truncate border-r border-gray-700 px-1">{status}</td>
{/if}
