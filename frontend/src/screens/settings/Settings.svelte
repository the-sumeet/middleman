<script>
    import { GetConfig, GenerateCert } from "../../../wailsjs/go/main/App";
    import { onDestroy, onMount } from "svelte";
    import { AddProxyPort, AddWebPort } from "../../../wailsjs/go/main/App";
    import PortSetting from "./PortSetting.svelte";
    import { errorMessage } from "../../../src/stores";
    import { config } from "../../../src/stores";
    import { refreshConfig } from "../../../src/stores";

    let certPath;
    let generatingCert = false;
    let configValue;
    let proxyServerPort;
    let webServerPort;

    function generateCert() {
        if (generatingCert === true) {
            return;
        }

        generatingCert = true;
        GenerateCert().then((result) => {
            if (result.error != "") {
                errorMessage.set(result.error);
            }
            generatingCert = false;
        });
    }

    const unSubConfig = config.subscribe((value) => {
        configValue = value;
    });

    const addProxyPort = async (newPort) => {
        const result = await AddProxyPort(newPort);
        refreshConfig();
        return result;
    };

    const addWebPort = async (newPort) => {
        const result = await AddWebPort(newPort);
        refreshConfig();
        return result;
    };

    onMount(async () => {
        certPath = configValue.certPath;
        proxyServerPort = configValue.proxyServerPort;
        webServerPort = configValue.webServerPort;
    });

    onDestroy(() => {
        unSubConfig();
    });
</script>

<div class="p-4 w-full">
    <!-- Certificates -->
    <div>
        <h2 class="text-2xl font-medium text-white">Certificates</h2>

        <div class="mt-4">
            <label for="username" class="block text-sm text-gray-300"
                >Certificate Path</label
            >

            <input
                bind:value={certPath}
                readonly
                type="text"
                placeholder="Certificate Path"
                class="block mt-2 w-full placeholder-gray-400/70 placeholder-gray-500 rounded-lg border px-5 py-2.5 focus:outline-none focus:ring-opacity-40 border-gray-600 bg-gray-900 text-gray-300 focus:border-1"
            />
        </div>

        <button
            on:click={generateCert}
            class="w-48 mt-4 px-6 py-2 font-medium tracking-wide text-white capitalize transition-colors duration-300 transform bg-blue-600 rounded-lg hover:bg-blue-500 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-80"
        >
            {#if generatingCert}
                <div class="animate-spin">
                    <i class="bi bi-arrow-clockwise"></i>
                </div>
            {:else}
                New Certificate
            {/if}
        </button>
    </div>

    <PortSetting
        title="Proxy Server Port"
        description="Port for proxy server to run on"
        port={proxyServerPort}
        addPortFunc={addProxyPort}
    />

    <PortSetting
        title="Web Server Port"
        description="Port for web server to run on"
        port={webServerPort}
        addPortFunc={addWebPort}
    />
</div>
