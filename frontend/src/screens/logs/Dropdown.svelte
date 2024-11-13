<script>
    export let title;
    export let onChange;
    export let values;

    import { dropdownNames } from "./constants";

    var selected = new Set();
    var displayTitle = title;
    var opened = false;

    $: displayTitle =
        selected.size > 0
            ? `(${selected.size}) ${values.filter((value) => selected.has(value)).join(",")}`
            : title;
    $: textColor = selected.size > 0 ? "text-green-400" : "text-white";

    function toggle(value) {
        if (selected.has(value)) {
            selected.delete(value);
            selected = new Set(selected);
        } else {
            selected.add(value);
            selected = new Set(selected);
        }
        onChange(Array.from(selected));
    }
</script>

<div class="relative">
    <!-- Button -->
    <button
        on:click={() => {
            opened = !opened;
        }}
        class="{textColor} m-0 flex items-center w-full relative z-10 p-2 bg-white border border-transparent rounded-md focus:border-blue-500 focus:ring-opacity-40 dark:focus:ring-opacity-40 focus:ring-blue-300 dark:focus:ring-blue-400 focus:ring dark:bg-gray-800 focus:outline-none"
    >
        <span class="flex gap-2 text-xs truncate">
            {#if opened}
            <i class=" bi bi-chevron-up"></i>
            {:else}
            <i class=" bi bi-chevron-down"></i>
            {/if}
            {dropdownNames[displayTitle] || displayTitle}
        </span>
    </button>

    <!-- Dropdown content -->
    {#if opened}
        <div
            class="absolute right-0 z-20 w-full py-2 mt-2 origin-top-right bg-white rounded-md shadow-xl dark:bg-gray-800"
        >
            {#each values as value}
                <a
                    href="#!"
                    on:click={() => toggle(value)}
                    class="w-full block content-start px-4 py-3 text-xs transition-colors duration-300 transform text-gray-300 hover:bg-gray-700 dark:hover:text-white"
                >
                    {#if selected.has(value)}
                        <i class="text-green-400 bi bi-check"></i>
                    {/if}
                    {dropdownNames[value] || value}
                </a>
            {/each}
        </div>
    {/if}
</div>
