<script>
  import "bootstrap-icons/font/bootstrap-icons.css";
  import { currentRule, webServerRunning, proxyServerRunning } from "./stores";
  import { onDestroy } from "svelte";
  import { StartProxy, StopProxy } from "../wailsjs/go/main/App";
  import { currentPage } from "./stores";
  import { HOME, LOGS, SERVER, SETTINGS } from "./constants";

  // let isServerRunning = false;
  let currPage;
  const activePageCss = "text-blue-500 bg-gray-800";
  const inactivePageCss = "text-gray-200 hover:bg-gray-800";

  const unCurrentPage = currentPage.subscribe((value) => {
    currPage = value;
  });

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
  <a
    on:click={() => setPage(HOME)}
    href="#!"
    class="{currPage == HOME
      ? activePageCss
      : inactivePageCss} p-1.5 transition-colors duration-200 rounded-lg text-blue-400"
  >
    <i class="text-2xl bi bi-list-task"></i>
  </a>

  <a
    on:click={() => setPage(SERVER)}
    href="#!"
    class="{inactivePageCss} p-1.5 focus:outline-nones transition-colors duration-200 rounded-lg"
  >
    <button type="button" class="relative">
      <!-- Web server indicator -->
      <div
        class="absolute inline-flex items-center justify-center w-4 h-4 text-xs font-bold text-white {$webServerRunning
          ? 'bg-green-500'
          : 'bg-green-500'} border-2 border-white rounded-full -bottom-2 -end-2 dark:border-gray-900"
      ></div>
      <i class="text-2xl bi bi-hdd-network"></i>
      <!-- Proxy server indicator -->
      <div
        class="absolute inline-flex items-center justify-center w-4 h-4 text-xs font-bold text-white {$proxyServerRunning
          ? 'bg-green-500'
          : 'bg-red-500'} border-2 border-white rounded-full -top-2 -end-2 dark:border-gray-900"
      ></div>
      
    </button>
  </a>

  <a
    on:click={() => setPage(LOGS)}
    href="#!"
    class="{currPage == LOGS
      ? activePageCss
      : inactivePageCss} p-1.5 transition-colors duration-200 rounded-lg text-blue-400"
  >
    <i class="text-2xl bi bi-arrow-down-up"></i>
  </a>

  <!-- <a
        on:click={() => setPage(REQUESTS)}
        href="#!"
        class="{currPage == REQUESTS
            ? activePageCss
            : inactivePageCss} p-1.5 transition-colors duration-200 rounded-lg text-blue-400"
    >
        <i class="text-2xl bi bi-arrow-down-up"></i>
    </a> -->

  <!-- Settings -->
  <a
    on:click={() => setPage(SETTINGS)}
    href="#!"
    class="{currPage == SETTINGS
      ? activePageCss
      : inactivePageCss} p-1.5 transition-colors duration-200 rounded-lg text-blue-400"
  >
    <i class="text-2xl bi bi-gear"></i>
  </a>
</div>
