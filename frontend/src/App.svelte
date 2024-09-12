<script>
  import "bootstrap-icons/font/bootstrap-icons.css";
  import Rules from "./screens/rules/Rules.svelte";
  import { currentRule, serverRunning } from "./stores";
  import { onDestroy } from "svelte";
  import { StartProxy, StopProxy } from "../wailsjs/go/main/App";
  import RedirectList from "./screens/redirects/redirect_list.svelte";
  import { REDIRECT_RULE } from "./constants";
  import { AddRedirect, GetRedirects } from "../wailsjs/go/main/App";
  import { main } from "../wailsjs/go/models";
  import { redirects } from "./stores";

  let selectedRule;
  let isServerRunning = false;

  const unSubCurrentRule = currentRule.subscribe((value) => {
    selectedRule = value;
  });

  function add() {
    if (selectedRule == REDIRECT_RULE) {
      AddRedirect(new main.Redirect());
      GetRedirects().then((res) => {
        redirects.set(res.redirects);
      });
    }
  }

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

  onDestroy(() => {
    unSubCurrentRule();
  });
</script>

<aside class="flex bg-gray-900">
  <!-- Sidebar 1 -->
  <div
    class="flex flex-col items-center w-16 h-screen py-8 space-y-8 bg-gray-900 border-gray-700"
  >
    <a href="#">
      <img
        class="w-auto h-6"
        src="https://merakiui.com/images/logo.svg"
        alt=""
      />
    </a>

    <a
      on:click={toggleProxy}
      href="#!"
      class="p-1.5 text-gray-500 focus:outline-nones transition-colors duration-200 rounded-lg dark:text-gray-400 dark:hover:bg-gray-800 hover:bg-gray-100"
    >
      {#if isServerRunning}
        <i class="text-3xl text-red-700 bg-red-700 bi bi-pause"></i>
      {:else}
        <i class="text-3xl text-green-700 bg-green-700 bi bi-play"></i>
      {/if}
    </a>

    <a
      href="#"
      class="p-1.5 text-blue-500 transition-colors duration-200 rounded-lg dark:text-blue-400 bg-gray-800"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="w-6 h-6"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M15 19.128a9.38 9.38 0 002.625.372 9.337 9.337 0 004.121-.952 4.125 4.125 0 00-7.533-2.493M15 19.128v-.003c0-1.113-.285-2.16-.786-3.07M15 19.128v.106A12.318 12.318 0 018.624 21c-2.331 0-4.512-.645-6.374-1.766l-.001-.109a6.375 6.375 0 0111.964-3.07M12 6.375a3.375 3.375 0 11-6.75 0 3.375 3.375 0 016.75 0zm8.25 2.25a2.625 2.625 0 11-5.25 0 2.625 2.625 0 015.25 0z"
        />
      </svg>
    </a>

    <a
      href="#"
      class="p-1.5 text-gray-500 focus:outline-nones transition-colors duration-200 rounded-lg dark:text-gray-400 dark:hover:bg-gray-800 hover:bg-gray-100"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="w-6 h-6"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M10.5 6a7.5 7.5 0 107.5 7.5h-7.5V6z"
        />
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M13.5 10.5H21A7.5 7.5 0 0013.5 3v7.5z"
        />
      </svg>
    </a>

    <a
      href="#"
      class="p-1.5 text-gray-500 focus:outline-nones transition-colors duration-200 rounded-lg dark:text-gray-400 dark:hover:bg-gray-800 hover:bg-gray-100"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="w-6 h-6"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M14.857 17.082a23.848 23.848 0 005.454-1.31A8.967 8.967 0 0118 9.75v-.7V9A6 6 0 006 9v.75a8.967 8.967 0 01-2.312 6.022c1.733.64 3.56 1.085 5.455 1.31m5.714 0a24.255 24.255 0 01-5.714 0m5.714 0a3 3 0 11-5.714 0"
        />
      </svg>
    </a>

    <a
      href="#"
      class="p-1.5 text-gray-500 focus:outline-nones transition-colors duration-200 rounded-lg dark:text-gray-400 dark:hover:bg-gray-800 hover:bg-gray-100"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
        class="w-6 h-6"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M9.594 3.94c.09-.542.56-.94 1.11-.94h2.593c.55 0 1.02.398 1.11.94l.213 1.281c.063.374.313.686.645.87.074.04.147.083.22.127.324.196.72.257 1.075.124l1.217-.456a1.125 1.125 0 011.37.49l1.296 2.247a1.125 1.125 0 01-.26 1.431l-1.003.827c-.293.24-.438.613-.431.992a6.759 6.759 0 010 .255c-.007.378.138.75.43.99l1.005.828c.424.35.534.954.26 1.43l-1.298 2.247a1.125 1.125 0 01-1.369.491l-1.217-.456c-.355-.133-.75-.072-1.076.124a6.57 6.57 0 01-.22.128c-.331.183-.581.495-.644.869l-.213 1.28c-.09.543-.56.941-1.11.941h-2.594c-.55 0-1.02-.398-1.11-.94l-.213-1.281c-.062-.374-.312-.686-.644-.87a6.52 6.52 0 01-.22-.127c-.325-.196-.72-.257-1.076-.124l-1.217.456a1.125 1.125 0 01-1.369-.49l-1.297-2.247a1.125 1.125 0 01.26-1.431l1.004-.827c.292-.24.437-.613.43-.992a6.932 6.932 0 010-.255c.007-.378-.138-.75-.43-.99l-1.004-.828a1.125 1.125 0 01-.26-1.43l1.297-2.247a1.125 1.125 0 011.37-.491l1.216.456c.356.133.751.072 1.076-.124.072-.044.146-.087.22-.128.332-.183.582-.495.644-.869l.214-1.281z"
        />
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
        />
      </svg>
    </a>
  </div>

  <!-- RuleType Sidebar -->
  <Rules />

  <!-- Collection -->
  <div class="p-4 flex flex-col w-full h-screen">
    <!-- Toolbar -->
    <div class="flex justify-end rounded-md">
      <button
        on:click={add}
        class="px-4 py-2 font-medium tracking-wide text-white capitalize transition-colors duration-300 transform bg-blue-600 rounded-lg hover:bg-blue-500 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-80"
      >
        <i class="bi bi-plus-lg"></i>
      </button>
    </div>

    <div class="mt-4 overflow-y-scroll">
      {#if selectedRule == REDIRECT_RULE}
        <RedirectList />
      {/if}
    </div>

    <!-- {#if currentColl !== undefined && currentColl >= 0}
      <CollectionPage />
    {/if} -->
  </div>
</aside>
