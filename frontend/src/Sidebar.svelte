<script>
    import "bootstrap-icons/font/bootstrap-icons.css";
    import { currentRule } from "./stores";
    import { onDestroy } from "svelte";
    import { StartProxy, StopProxy } from "../wailsjs/go/main/App";
    import { currentPage } from "./stores";
    import { HOME, REQUESTS } from "./constants";

    let isServerRunning = false;
    let currPage;
    const activePageCss = "text-blue-500 bg-gray-800";
    const inactivePageCss = "text-gray-200 hover:bg-gray-800";

    const unCurrentPage = currentPage.subscribe((value) => {
        currPage = value;
    });

    function toggleProxy() {
        if (isServerRunning) {
            console.log("Stopping proxy");
            StopProxy().then(() => {
                isServerRunning = false;
                // serverRunning.set(false);
                console.log("Proxy stopped");
            });
        } else {
            console.log("Starting proxy");
            StartProxy().then(() => {});
            isServerRunning = true;
            console.log("Proxy started");
        }
    }

    function setPage(page) {
        currentPage.set(page);
    }

    onDestroy(() => {
        unCurrentPage();
    });
</script>

<div
    class="flex flex-col items-center w-16 h-screen py-8 space-y-8 bg-gray-900 b border-r border-gray-800"
>
    <a on:click={() => setPage(HOME)} href="#!">
        <img
            class="w-auto h-6"
            src="https://merakiui.com/images/logo.svg"
            alt=""
        />
    </a>

    <a
        on:click={toggleProxy}
        href="#!"
        class="{inactivePageCss} p-1.5 focus:outline-nones transition-colors duration-200 rounded-lg"
    >
        {#if isServerRunning}
            <i class=" text-red-700 bg-red-700 bi bi-pause"></i>
        {:else}
            <i class="text-2xl bi bi-play"></i>
        {/if}
    </a>

    <a  
        on:click={() => setPage(REQUESTS)}
        href="#!"
        class="{currPage == REQUESTS ? activePageCss : inactivePageCss } p-1.5 transition-colors duration-200 rounded-lg text-blue-400 "
    >
        <i class="text-2xl bi bi-arrow-down-up"></i>
    </a>
</div>
