<script>
    export let bsIcon;
    export let page;
    export let descriptipn;
    export let green
    export let onClick;

    import { currentPage } from "../../../src/stores";
    import { onDestroy } from "svelte";
    import { activePageCss, inactivePageCss } from "./constants";

    let currPage;
    let error = null;
    let showTooltip = false;

    let indicatorCss = "absolute inline-flex items-center justify-center w-4 h-4 text-xs font-bold text-white border-2 rounded-full -top-2 -end-2 border-gray-900";
    if (green === true) {
        indicatorCss += " bg-green-500";
    } else if (green === false) {
        indicatorCss += " bg-red-500";
    }
    console.log(indicatorCss);
    
    const unCurrentPage = currentPage.subscribe((value) => {
        currPage = value;
    });

    async function onClickHandler() {
        const res = await onClick();
        if (!res) {
            error = res;
        }
    }

    onDestroy(() => {
        unCurrentPage();
    });
</script>

<div
    class="{currPage == page
        ? activePageCss
        : inactivePageCss} p-1.5 transition-colors duration-200 rounded-lg text-blue-400 relative inline-block"
>
    <!-- svelte-ignore a11y-mouse-events-have-key-events -->
    <button
        on:mouseover={() => (showTooltip = true)}
        on:mouseleave={() => (showTooltip = false)}
        on:click={onClickHandler}
        class="text-gray-600 transition-colors duration-200 focus:outline-none dark:text-gray-200 dark:hover:text-blue-400 hover:text-blue-500"
    >
        <i class="{currPage == page
        ? activePageCss
        : inactivePageCss} text-2xl {bsIcon}"></i>
    </button>

    <!-- Indicator -->
    {#if error}
        <i
            class="absolute inline-flex items-center justify-center w-4 h-4 text-lg font-bold text-red-400 border-2 -top-2 -end-2 border-gray-900 bi bi-exclamation-triangle-fill"
        ></i>
    {:else if green !== null}
        <div
            title="Web server"
            class="{indicatorCss}"
        ></div>
    {/if}

    {#if descriptipn && showTooltip}
        <div
            class="z-10 absolute flex items-center justify-start w-48 p-3 rounded-lg shadow-lg left-12 -top-4 dark:shadow-none shadow-gray-200 bg-gray-800 text-white"
        >
            <div class="flex flex-col items-start">
                <span class="">{descriptipn}</span>
                <span class="text-xs text-red-400">{descriptipn}</span>
            </div>

            <svg
                xmlns="http://www.w3.org/2000/svg"
                class="absolute w-6 h-6 text-white transform rotate-45 -translate-y-1/2 fill-current -left-3 top-1/2 dark:text-gray-800"
                stroke="currentColor"
                viewBox="0 0 24 24"
            >
                <path
                    d="M20 3H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1z"
                ></path>
            </svg>
        </div>
    {/if}
</div>
