<script>
  import "bootstrap-icons/font/bootstrap-icons.css";
  import { currentRule, webServerRunning, proxyServerRunning } from "../stores";
  import { onDestroy } from "svelte";
  import { StartProxy, StopProxy } from "../../wailsjs/go/main/App";
  import { currentPage } from "../stores";
  import { HOME, LOGS, SERVER, SETTINGS } from "../constants";

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
  <!-- Rules -->
  <a
    on:click={() => setPage(HOME)}
    href="#!"
    class="{currPage == HOME
      ? activePageCss
      : inactivePageCss} p-1.5 transition-colors duration-200 rounded-lg text-blue-400"
  >
    <i class="text-2xl bi bi-list-task"></i>
  </a>

  <!-- Start Proxy -->
  <a
    title="Start Proxy"
    href="#!"
    class="{currPage == HOME
      ? activePageCss
      : inactivePageCss} p-1.5 transition-colors duration-200 rounded-lg text-blue-400"
  >
    <i class="text-2xl bi bi-list-task"></i>
  </a>

  <!-- Start Proxy -->
  <div
    class="{currPage == HOME
      ? activePageCss
      : inactivePageCss} p-1.5 transition-colors duration-200 rounded-lg text-blue-400 relative inline-block"
  >
    <button
      class="text-gray-600 transition-colors duration-200 focus:outline-none dark:text-gray-200 dark:hover:text-blue-400 hover:text-blue-500"
    >
      <i class="text-2xl bi bi-list-task"></i>
    </button>

    <div
      class="absolute flex items-center justify-center w-48 p-3 text-gray-600 bg-white rounded-lg shadow-lg left-12 -top-4 dark:shadow-none shadow-gray-200 dark:bg-gray-800 dark:text-white"
    >
      <span class="truncate">This is a tooltip</span>

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
  </div>

  <a
    on:click={() => setPage(SERVER)}
    href="#!"
    class="{inactivePageCss} p-1.5 focus:outline-nones transition-colors duration-200 rounded-lg"
  >
    <button type="button" class="relative">
      <!-- Web server indicator -->
      <div
        title="Web server"
        class="absolute inline-flex items-center justify-center w-4 h-4 text-xs font-bold text-white {$webServerRunning
          ? 'bg-green-500'
          : 'bg-red-500'} border-2 border-white rounded-full -bottom-2 -end-2 dark:border-gray-900"
      ></div>
      <i class="text-2xl bi bi-hdd-network"></i>
      <!-- Proxy server indicator -->
      <div
        title="Proxy server"
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
