<script>
    import "bootstrap-icons/font/bootstrap-icons.css";
    import { StartProxy, StopProxy } from "../../../wailsjs/go/main/App";
    import { serverRunning } from "../../../src/stores";
    import { onDestroy } from "svelte";

    let isServerRunning;
    const startButtonCss = "bg-blue-600 hover:bg-blue-500 focus:ring-blue-300";
    const stopButtonCss = "bg-red-600 hover:bg-red-500 focus:ring-red-300";
    let error = "";
    let port = "8080";

    const ubSubServerRunning = serverRunning.subscribe((value) => {
        isServerRunning = value;
    });

    function validatePort() {
        let portNumber = parseInt(port);
        if (isNaN(portNumber)) {
            error = "Invalid number";
            return false;
        }

        if (portNumber.toString().length != 4) {
            error = "Port must be of length 4";
            return false;
        }
        return true;
    }

    function toggleProxy() {
        error = "";

        if (isServerRunning === true) {
            console.log("Stopping proxy");
            StopProxy().then(() => {
                serverRunning.set(false);
                // serverRunning.set(false);
                console.log("Proxy stopped");
            });
        } else {
            if (!validatePort()) {
                return;
            }
            const portNumber = parseInt(port);
            console.log("Starting proxy");
            StartProxy(portNumber).then((result) => {
                if (result.error != "") {
                    error = result.error;
                    serverRunning.set(false);

                    return;
                } else {
                    serverRunning.set(true);

                    console.log("Proxy started");
                }
            });
        }
    }

    onDestroy(() => {
        ubSubServerRunning();
    });
</script>

<div class="flex flex-col w-full">
    <h1 class="mt-4 text-2xl text-white mx-auto">Proxy Server</h1>

    <div class="mt-10 flex gap-2 mx-auto items-center">
        <input
            bind:value={port}
            type="text"
            min="0"
            placeholder="8080"
            class="max-w-sm block placeholder-gray-400/70 dark:placeholder-gray-500 rounded-lg border border-gray-200 bg-white px-5 py-2.5 text-gray-700 focus:border-blue-400 focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-40 dark:border-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:focus:border-blue-300"
        />
        <button
            on:click={toggleProxy}
            class="{isServerRunning
                ? stopButtonCss
                : startButtonCss} flex items-center rounded-lg px-4 py-2.5 font-medium tracking-wide text-white capitalize transition-colors duration-300 transform focus:outline-none focus:ring focus:ring-opacity-80"
        >
            {#if isServerRunning}
                Stop Server
            {:else}
                Start Server
            {/if}
        </button>
    </div>

    {#if error != ""}
        <div
            class="mt-4 mx-auto flex w-full max-w-sm overflow-hidden bg-white rounded-lg shadow-md dark:bg-gray-800"
        >
            <div class="flex items-center justify-center w-12 bg-red-500">
                <i class="text-white text-xl bi bi-exclamation-triangle"></i>
            </div>

            <div class="px-4 py-2 -mx-3">
                <div class="mx-3">
                    <span class="font-semibold text-red-500 dark:text-red-400"
                        >Error</span
                    >
                    <p class="text-sm text-gray-600 dark:text-gray-200">
                        {error}
                    </p>
                </div>
            </div>
        </div>
    {/if}
</div>
