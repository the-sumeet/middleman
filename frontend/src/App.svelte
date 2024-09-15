<script>
  import "bootstrap-icons/font/bootstrap-icons.css";
  import { onDestroy } from "svelte";
  import { currentPage } from "./stores";
  import HomeScreen from "./screens/home/HomeScreen.svelte";
  import {REQUESTS } from "./constants";
  import Requests from "./screens/requests/Requests.svelte";
  import Sidebar from "./Sidebar.svelte";


  let currPage;

  const unSubServerRunning = currentPage.subscribe((value) => {
    currPage = value;
  });


  onDestroy(() => {
    unSubServerRunning();
  });
</script>

<aside class="flex bg-gray-900">

  <Sidebar />

  {#if currPage === REQUESTS}
  <Requests />
  {:else}
  <HomeScreen />
  {/if}

</aside>
