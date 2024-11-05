<script>
  import "bootstrap-icons/font/bootstrap-icons.css";
  import {
    webServerRunning,
    proxyServerRunning,
  } from "../../stores";
  import { onDestroy } from "svelte";
  import {
    StartProxy,
    StartWebServer,
    StopProxy,
    StopWebServer,
  } from "../../../wailsjs/go/main/App";
  import { currentPage } from "../../stores";
  import { HOME, LOGS, SERVER, SETTINGS } from "../../constants";
  import Button from "./Button.svelte";
  import { GetWebServerPath } from "../../../wailsjs/go/main/App";

  let currPage;
  let isProxyRunning;
  let isWebRunning;
  let proxyServerDesc = null;
  let webServerDesc = null;
  let webServerPath;

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

  (isWebRunning) ? `Web server running on ${webServerPath}` : null

  $: webServerDescription2 = (webServerPath != null && isWebRunning === true) ? `Web server running run on ${webServerPath}` : null;

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

  GetWebServerPath().then((res) => {
    webServerPath = res;
  });

  onDestroy(() => {
    unCurrentPage();
    unsubProxyServerRunning();
    unsubWebServerRunning();
  });
</script>

<div
  class="flex flex-col items-center w-16 h-screen py-8 space-y-8 bg-gray-800 b border-r border-gray-800"
>
  <!-- Rules -->
  <Button
    indicatorStore={null}
    bsIcon={"bi bi bi-list-task"}
    page={HOME}
    descriptipn={null}
    onClick={() => setPage(HOME)}
    descriptipn2={null}
  />

  <Button
    onClick={() => setPage(LOGS)}
    indicatorStore={null}
    bsIcon={"bi bi-arrow-down-up"}
    page={LOGS}
    descriptipn={null}
    descriptipn2={null}
  />

  <Button
    onClick={() => setPage(SETTINGS)}
    indicatorStore={null}
    bsIcon={"bi bi-gear"}
    page={SETTINGS}
    descriptipn={null}
    descriptipn2={null}
  />

  <div class="w-full px-2">
    <hr class="border border-gray-500 w-full" />
  </div>

  <Button
    onClick={toggleProxyServer}
    indicatorStore={proxyServerRunning}
    bsIcon={"bi bi-hdd-network"}
    page={null}
    descriptipn={proxyServerDesc}
    descriptipn2={null}
  />

  <Button
    onClick={toggleWebServer}
    indicatorStore={webServerRunning}
    bsIcon={"bi bi-globe2"}
    page={null}
    descriptipn={webServerDesc}
    descriptipn2={webServerDescription2}
  />
</div>
