<script>
    export let changed;
    export let save;
    export let remove;
    export let enableDisable;
    export let enabled;
</script>


<div class="flex space-x-2 mt-2 justify-end">
    {#if changed}
        <button
            on:click={save}
            class="px-4 py-2 font-medium tracking-wide text-white capitalize transition-colors duration-300 transform bg-blue-600 rounded-lg hover:bg-blue-500 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-80"
        >
            <i class="bi bi-floppy"></i>
        </button>
    {:else}
        <button
            on:click={remove}
            class="px-4 py-2 font-medium tracking-wide text-white capitalize transition-colors duration-300 transform bg-red-600 rounded-lg hover:bg-red-500 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-80"
        >
            <i class="bi bi-trash"></i>
        </button>

        {#if !changed}
        <button
            on:click={enableDisable}
            class="px-4 py-2 font-medium tracking-wide text-white capitalize transition-colors duration-300 transform focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-80"
        >   
            {#if enabled}
            <i class="text-green-500 text-4xl bi bi-toggle-on"></i>
            {:else}
            <div class="rotate-180">
                <i class=" text-red-500 text-4xl bi bi-toggle-on"></i>
            </div>
            {/if}
        </button>
        {/if}
    {/if}
</div>