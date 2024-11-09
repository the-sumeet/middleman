<script>
    import "bootstrap-icons/font/bootstrap-icons.css";
    import { proxyServerRunning, webServerRunning } from "../../../src/stores";
    import { onDestroy, onMount } from "svelte";
    import { config } from "../../../src/stores";
    import ServerInput from "./ServerInput.svelte";
    import {
        StartProxy,
        StopProxy,
        StartWebServer,
        StopWebServer,
    } from "../../../wailsjs/go/main/App";

    let port;
    let proxyServerPort;
    let webServerPort;
    let webServerPath;
    let configValue;

    const unSubConfig = config.subscribe((value) => {
        configValue = value;
        proxyServerPort = configValue.proxyServerPort;
        webServerPort = configValue.webServerPort;
    });

    onDestroy(() => {
        unSubConfig();
    });
</script>

<div class="flex flex-col w-full p-4">
    <ServerInput
        startProxy={StartProxy}
        stopProxy={StopProxy}
        title={"Proxy Server"}
        serverRunning={proxyServerRunning}
        port={proxyServerPort}
    />
    <ServerInput
        startProxy={StartWebServer}
        stopProxy={StopWebServer}
        title={"Web Server"}
        serverRunning={webServerRunning}
        port={webServerPort}
    />

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
