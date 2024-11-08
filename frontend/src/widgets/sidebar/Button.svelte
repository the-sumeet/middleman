<script>
    export let bsIcon;
    export let page;
    export let descriptipn;
    export let descriptipn2;
    export let indicatorStore;
    export let onClick;

    import { currentPage, errorMessage } from "../../../src/stores";
    import { onDestroy } from "svelte";
    import { activePageCss, inactivePageCss } from "./constants";

    let on = null;

    let currPage;
    let showTooltip = false;

    const indicatorCssConst =
        "absolute inline-flex items-center justify-center w-4 h-4 text-xs font-bold text-white border-2 rounded-full -top-2 -end-2 border-gray-900";

    let indicatorCss = "";

    $: if (on === true) {
        indicatorCss = indicatorCssConst + " bg-green-500";
    } else if (on === false) {
        indicatorCss = indicatorCssConst + " bg-red-500";
    } else {
        indicatorCss = indicatorCssConst;
    }

    const unCurrentPage = currentPage.subscribe((value) => {
        currPage = value;
    });

    let ubSubIndicatorStore = null;
    if (indicatorStore) {
        ubSubIndicatorStore = indicatorStore.subscribe((value) => {
            on = value;
        });
    }

    async function onClickHandler() {
        const res = await onClick();
        if (res && res != "") {
            errorMessage.set(res);
        }
    }

    onDestroy(() => {
        unCurrentPage();
        if (ubSubIndicatorStore) {
            ubSubIndicatorStore();
        }
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
        class="text-gray-600 transition-colors duration-200 focus:outline-none text-gray-200 hover:text-blue-400 hover:text-blue-500"
    >
        <i
            class="{currPage == page
                ? activePageCss
                : inactivePageCss} text-2xl {bsIcon}"
        ></i>
    </button>

    <!-- Indicator -->
    {#if on !== null}
        <div title="Web server" class={indicatorCss}></div>
    {/if}

    {#if descriptipn && showTooltip}
        <div
            class="z-10 absolute flex items-center justify-start w-48 p-3 rounded-lg shadow-lg left-12 -top-4 shadow-none shadow-gray-200 bg-gray-800 text-white"
        >
            <div class="flex flex-col items-start">
                <span class="">{descriptipn}</span>
                {#if descriptipn2}
                    <span class="text-xs text-gray-400">{descriptipn2}</span>
                {/if}
            </div>

            <svg
                xmlns="http://www.w3.org/2000/svg"
                class="absolute w-6 h-6  transform rotate-45 -translate-y-1/2 fill-current -left-3 top-1/2 text-gray-800"
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
