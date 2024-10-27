<script>
    import { GetConfig, GenerateCert } from "../../../wailsjs/go/main/App";
    import { onMount } from "svelte";
    import { AddProxyPort, AddWebPort } from "../../../wailsjs/go/main/App";
    import PortSetting from "./PortSetting.svelte";

    let certPath;
    let generatingCert = false;

    let proxyServerPort;
    let webServerPort;

    function generateCert() {
        if (generatingCert === true) {
            return;
        }

        generatingCert = true;
        GenerateCert().then(() => {
            generatingCert = false;
        });
    }
    onMount(async () => {
        const config = await GetConfig();
        certPath = config.certPath;
        proxyServerPort = config.proxyServerPort;
        webServerPort = config.webServerPort;
    });
</script>

<div class="p-4 w-full">
    <!-- Certificates -->
    <div>
        <h2 class="text-2xl font-medium text-gray-800 dark:text-white">
            Certificates
        </h2>

        <div class="mt-4">
            <label
                for="username"
                class="block text-sm text-gray-500 dark:text-gray-300"
                >Certificate Path</label
            >

            <input
                bind:value={certPath}
                readonly
                type="text"
                placeholder="Certificate Path"
                class="block mt-2 w-full placeholder-gray-400/70 dark:placeholder-gray-500 rounded-lg border border-gray-200 px-5 py-2.5 focus:outline-none focus:ring-opacity-40 dark:border-gray-600 bg-gray-900 text-gray-300 focus:border-1"
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

    <div class="mt-8">
        <PortSetting
        title="Proxy Server Port"
        description="Default port for proxy server to run on"
        port={proxyServerPort}
        addPortFunc={AddProxyPort}
    />

    <PortSetting
        title="Web Server Port"
        description="Default port for web server to run on"
        port={webServerPort}
        addPortFunc={AddWebPort}
    />
    </div>
</div>
