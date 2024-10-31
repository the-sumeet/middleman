<script>
  import "bootstrap-icons/font/bootstrap-icons.css";
  import { onDestroy } from "svelte";
  import { currentPage } from "./stores";
  import HomeScreen from "./screens/home/HomeScreen.svelte";
  import { REQUESTS, SERVER, SETTINGS, LOGS } from "./constants";
  import Requests from "./screens/requests/Requests.svelte";
  import Sidebar from "./widgets/sidebar/Sidebar.svelte";
  import ServerScreen from "./screens/server/ServerScreen.svelte";
  import Settings from "./screens/settings/Settings.svelte";
  import LogsScreen from "./screens/logs/LogsScreen.svelte";
  import {
    errorMessage,
    infoMessage,
    warningMessage,
    successMessage,
  } from "./stores";
  import Alert from "./widgets/Alert.svelte";
  import {
    ALERT_ERROR,
    ALERT_INFO,
    ALERT_SUCCESS,
    ALERT_WARNING,
  } from "./constants";
  
  let currPage;

  const unSubServerRunning = currentPage.subscribe((value) => {
    currPage = value;
  });

  onDestroy(() => {
    unSubServerRunning();
  });
</script>

<div class="relative">
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

  {#if $errorMessage != "" || $infoMessage != "" || $warningMessage != "" || $successMessage != ""}
    <div class="absolute bottom-0 w-full">
      {#if $errorMessage != ""}
        <Alert
          message={$errorMessage}
          type={ALERT_ERROR}
          clear={() => {
            errorMessage.set("");
          }}
          clearAfter={null}
        />
      {/if}

      {#if $infoMessage != ""}
        <Alert
          message={$infoMessage}
          type={ALERT_INFO}
          clear={() => {
            infoMessage.set("");
          }}
          clearAfter={5000}
        />
      {/if}

      {#if $warningMessage != ""}
        <Alert
          message={$warningMessage}
          type={ALERT_WARNING}
          clear={() => {
            warningMessage.set("");
          }}
          clearAfter={5000}
        />
      {/if}

      {#if $successMessage != ""}
        <Alert
          message={$successMessage}
          type={ALERT_SUCCESS}
          clear={() => {
            successMessage.set("");
          }}
          clearAfter={5000}
        />
      {/if}
    </div>
  {/if}
</div>
