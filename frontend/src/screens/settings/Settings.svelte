<script>
    import { GetConfig, GenerateCert } from "../../../wailsjs/go/main/App";
    import { onMount } from "svelte";
    import { AddConfigPort } from "../../../wailsjs/go/main/App";

    let certPath;
    let generatingCert = false;
    // Port
    let port;
    let portChanged = false;
    let portError = "";

    function generateCert() {
        if (generatingCert === true) {
            return;
        }

        generatingCert = true;
        GenerateCert().then(() => {
            generatingCert = false;
        });
    }

    function savePort() {
        if (port.length != 4) {
            portError = "Port must be of length 4";
            return;
        }
        AddConfigPort(port).then(() => {
            portError = "";
            portChanged = false;
        });
    }

    onMount(async () => {
        const config = await GetConfig();
        certPath = config.certPath;
        port = config.serverPort;
    });
</script>

<div class="p-4 w-full">
    <!-- Certificates -->
    <div>
        <h1 class="text-white font-bold text-xl">Certificates</h1>

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

    <!-- Proxy Server -->
    <div class="mt-8">
        <h1 class="text-white font-bold text-xl">Proxy Server</h1>

        <div class="mt-4">
            <label
                for="username"
                class="block text-sm text-gray-500 dark:text-gray-300"
                >Default port for server to run on</label
            >

            <input
                bind:value={port}
                on:input={() => (portChanged = true)}
                type="text"
                placeholder="Certificate Path"
                class="block mt-2 w-full placeholder-gray-400/70 dark:placeholder-gray-500 rounded-lg border border-gray-200 px-5 py-2.5 focus:outline-none focus:ring-opacity-40 dark:border-gray-600 bg-gray-900 text-gray-300 focus:border-1"
            />
        </div>

        {#if portError != ""}
        <p class="text-red-500">{portError}</p>
        {/if}

        {#if portChanged === true}
            <button
                on:click={savePort}
                class="mt-4 px-6 py-2 font-medium tracking-wide text-white capitalize transition-colors duration-300 transform bg-blue-600 rounded-lg hover:bg-blue-500 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-80"
            >
                <i class="bi bi-floppy"></i>
            </button>
        {/if}
    </div>
</div>
