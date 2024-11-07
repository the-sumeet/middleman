<script>
    export let method;
</script>


{#if method.toLowerCase() === "get"}
<td class="font-bold text-green-400 text-sm w-1/12 truncate border-r border-gray-700 px-1">{method}</td>
{:else if method.toLowerCase() === "post" || method.toLowerCase() === "put" || method.toLowerCase() === "patch"}
<td class="font-bold text-yellow-400 text-sm w-1/12 truncate border-r border-gray-700 px-1">{method}</td>
{:else if method.toLowerCase() === "delete"}
<td class="font-bold text-red-400 text-sm w-1/12 truncate border-r border-gray-700 px-1">{method}</td>
{:else}
<td class="font-bold text-gray-300 text-sm w-1/12 truncate border-r border-gray-700 px-1">{method}</td>
{/if}
