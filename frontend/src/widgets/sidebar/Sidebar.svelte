<script>
  import "bootstrap-icons/font/bootstrap-icons.css";
  import {
    currentRule,
    webServerRunning,
    proxyServerRunning,
  } from "../../stores";
  import { onDestroy } from "svelte";
  import { StartProxy, StartWebServer, StopProxy, StopWebServer } from "../../../wailsjs/go/main/App";
  import { currentPage } from "../../stores";
  import { HOME, LOGS, SERVER, SETTINGS } from "../../constants";
  import Button from "./Button.svelte";

  let currPage;
  let isProxyRunning;
  let isWebRunning;
  let proxyServerDesc = null;
  let webServerDesc = null;

  $: if (isProxyRunning) {
    proxyServerDesc = "Stop Proxy Server";
  } else {
    proxyServerDesc = "Start Proxy Server";
  }

  $: if (isWebRunning) {
    webServerDesc = "Stop Web Server";
  } else {
    webServerDesc = "Start Web Server";
  }

  const unCurrentPage = currentPage.subscribe((value) => {
    currPage = value;
  });

  const unsubProxyServerRunning = proxyServerRunning.subscribe((value) => {
    isProxyRunning = value;
  });

  const unsubWebServerRunning = webServerRunning.subscribe((value) => {
    isWebRunning = value;
  });

  function setPage(page) {
    currentPage.set(page);
  }

  async function toggleProxyServer() {
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

  async function toggleWebServer() {
    if ($webServerRunning === true) {
      console.log("Stopping web server");
      await StopWebServer();
      webServerRunning.set(false);
      console.log("Web server stopped");
    } else {
      console.log("Starting web server");
      const result = await StartWebServer();
      if (result.error != "") {
        webServerRunning.set(false);
        return result.error;
      } else {
        webServerRunning.set(true);
        console.log("Web server started");
      }
    }

    return null;
  }

  onDestroy(() => {
    unCurrentPage();
    unsubProxyServerRunning();
    unsubWebServerRunning();
  });
</script>

<div
  class="flex flex-col items-center w-16 h-screen py-8 space-y-8 bg-gray-900 b border-r border-gray-800"
>
  <!-- Rules -->
  <Button
    indicatorStore={null}
    bsIcon={"bi bi bi-list-task"}
    page={HOME}
    descriptipn={null}
    onClick={() => setPage(HOME)}
  />

  <Button
    onClick={() => setPage(LOGS)}
    indicatorStore={null}
    bsIcon={"bi bi-arrow-down-up"}
    page={null}
    descriptipn={null}
  />

  <Button
    onClick={() => setPage(SETTINGS)}
    indicatorStore={null}
    bsIcon={"bi bi-gear"}
    page={null}
    descriptipn={null}
  />

  <hr class="border border-gray-800 w-full" />

  <Button
    onClick={toggleProxyServer}
    indicatorStore={proxyServerRunning}
    bsIcon={"bi bi-hdd-network"}
    page={null}
    descriptipn={proxyServerDesc}
  />

  <Button
    onClick={toggleWebServer}
    indicatorStore={webServerRunning}
    bsIcon={"bi bi-hdd-network"}
    page={null}
    descriptipn={webServerDesc}
  />

  
</div>
