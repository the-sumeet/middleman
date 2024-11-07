<script>
    export let title;
    export let rule;
    export let icon = "text-xl font-bold text-white bi bi-shuffle";

    import { currentRule } from "../../stores";

    const selectedCss = "bg-gray-100 dark:bg-gray-700";
    const unselectedCss = "dark:hover:bg-gray-800 gap-x-2 hover:bg-gray-100";

    function select() {
        currentRule.set(rule);
    }
</script>

<button
    on:click={select}
    class="{$currentRule == rule
        ? selectedCss
        : unselectedCss} flex space-x-2 items-center w-full px-5 py-2 transition-colors duration-200 focus:outline-none"
>
    <div class="flex gap-2 items-center text-left rtl:text-right">
        <i class="text-xl font-bold text-white {icon}"></i>
        <h1
            class="text-sm font-medium text-gray-700 capitalize dark:text-white"
        >
            {title}
        </h1>

        <!-- <p class="text-xs  text-gray-400">11.2 Followers</p> -->
    </div>
</button>
