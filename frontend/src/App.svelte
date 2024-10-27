<script>
  import "bootstrap-icons/font/bootstrap-icons.css";
  import { onDestroy } from "svelte";
  import { currentPage } from "./stores";
  import HomeScreen from "./screens/home/HomeScreen.svelte";
  import {REQUESTS, SERVER, SETTINGS, LOGS } from "./constants";
  import Requests from "./screens/requests/Requests.svelte";
  import Sidebar from "./widgets/sidebar/Sidebar.svelte";
  import ServerScreen from "./screens/server/ServerScreen.svelte";
    import Settings from "./screens/settings/Settings.svelte";
  import LogsScreen from "./screens/logs/LogsScreen.svelte";
  let currPage;

  const unSubServerRunning = currentPage.subscribe((value) => {
    currPage = value;
  });


  onDestroy(() => {
    unSubServerRunning();
  });
</script>

<div class="flex bg-gray-900">

  <Sidebar />

  {#if currPage === REQUESTS}
  <Requests />
  {:else if currPage === SERVER}
  <ServerScreen />
  {:else if currPage === SETTINGS}
  <Settings />
  {:else if currPage === LOGS}
  <LogsScreen />
  {:else}
  <HomeScreen />
  {/if}

</div>