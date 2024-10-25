<script>
    import "bootstrap-icons/font/bootstrap-icons.css";
    import { proxyServerRunning, webServerRunning } from "../../../src/stores";
    import { onMount } from "svelte";
    import { GetConfig } from "../../../wailsjs/go/main/App";
    import { GetWebServerPath } from "../../../wailsjs/go/main/App";
    import ServerInput from "./ServerInput.svelte";

    let isServerRunning;
    const startButtonCss = "bg-blue-600 hover:bg-blue-500 focus:ring-blue-300";
    const stopButtonCss = "bg-red-600 hover:bg-red-500 focus:ring-red-300";
    let error = "";
    let port;
    let proxyServerPort;
    let webServerPort;
    let webServerPath;

    GetWebServerPath().then((path) => {
        webServerPath = path;
    });

    onMount(async () => {
        const config = await GetConfig();
        proxyServerPort = config.proxyServerPort;
        webServerPort = config.webServerPort;
    });
</script>

<div class="flex flex-col w-full p-4">
    <ServerInput title={"Proxy Server"} serverRunning={proxyServerRunning} port={proxyServerPort} />
    <ServerInput title={"Web Server"} serverRunning={webServerRunning} port={webServerPort} />

    {#if webServerPath}
        <div class="mx-auto mt-4 text-white">
            <p>
                Webserver running on <span class="p-2 text-blue-500 font-mono"
                    >127.0.0.1:{port}{webServerPath}</span
                >
            </p>
        </div>
    {/if}
</div>
