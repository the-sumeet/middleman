<script>
    export let text;
    export let color;
</script>

{#if color == "red"}
<span class="inline-flex items-center rounded-md bg-red-500/10 px-2 py-1 text-xs font-medium text-red-500 ring-1 ring-inset ring-red-600/10">{text}</span>
{:else if color == "yellow"}
<span class="inline-flex items-center rounded-md bg-yellow-500/10 px-2 py-1 text-xs font-medium text-yellow-500 ring-1 ring-inset ring-yellow-600/20">{text}</span>
{:else if color == "green"}
<span class="inline-flex items-center rounded-md bg-green-500/10 px-2 py-1 text-xs font-medium text-green-500 ring-1 ring-inset ring-green-600/20">{text}</span>
{:else if color == "gray"}
<span class="inline-flex items-center rounded-md bg-gray-300/10 px-2 py-1 text-xs font-medium text-gray-300 ring-1 ring-inset ring-gray-500/10">{text}</span>
{:else}
<span class="inline-flex items-center rounded-md bg-blue-400/10 px-2 py-1 text-xs font-medium text-blue-400 ring-1 ring-inset ring-blue-700/10">{text}</span>
{/if}