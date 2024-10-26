<script>
  import "bootstrap-icons/font/bootstrap-icons.css";
  import {
    currentRule,
    webServerRunning,
    proxyServerRunning,
  } from "../../stores";
  import { onDestroy } from "svelte";
  import { StartProxy, StopProxy } from "../../../wailsjs/go/main/App";
  import { currentPage } from "../../stores";
  import { HOME, LOGS, SERVER, SETTINGS } from "../../constants";
  import Button from "./Button.svelte";

  let currPage;
  let isProxyRunning = false;

  const activePageCss = "text-blue-500 bg-gray-800";
  const inactivePageCss = "text-gray-200 hover:bg-gray-800";

  const unCurrentPage = currentPage.subscribe((value) => {
    currPage = value;
  });

  const unsubProxyServerRunning = proxyServerRunning.subscribe((value) => {
    console.log("Proxy server running: ", value);
    isProxyRunning = value;
  });

  function setPage(page) {
    currentPage.set(page);
  }

  export async function toggleProxy() {
    if ($proxyServerRunning === true) {
      console.log("Stopping proxy");
      await StopProxy();
      proxyServerRunning.set(false);
      console.log("Proxy stopped");
    } else {
      console.log("Starting proxy");
      const result = await StartProxy();
      if (result.error != "") {
        proxyServerRunning.set(false);
        return result.error;
      } else {
        proxyServerRunning.set(true);
        console.log("Proxy started");
      }
    }

    return null;
  }

  onDestroy(() => {
    unCurrentPage();
    unsubProxyServerRunning();
  });
</script>

<div
  class="flex flex-col items-center w-16 h-screen py-8 space-y-8 bg-gray-900 b border-r border-gray-800"
>
  <!-- Rules -->
  <Button
    green={null}
    bsIcon={"bi bi bi-list-task"}
    page={HOME}
    descriptipn={null}
    onClick={() => setPage(HOME)}
  />

  <!-- Toggle Proxy -->
  <Button
    onClick={toggleProxy}
    green={isProxyRunning}
    bsIcon={"bi bi-hdd-network"}
    page={null}
    descriptipn={"Proxy server"}
  />

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
