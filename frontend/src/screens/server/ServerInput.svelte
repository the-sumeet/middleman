<script>
    export let serverRunning;
    export let port;
    export let title;
    export let startProxy;
    export let stopProxy;

    import "bootstrap-icons/font/bootstrap-icons.css";
    import { onDestroy, onMount } from "svelte";
    import { validatePort } from "./utils";

    let isServerRunning;
    const startButtonCss = "bg-blue-600 hover:bg-blue-500 focus:ring-blue-300";
    const stopButtonCss = "bg-red-600 hover:bg-red-500 focus:ring-red-300";
    let error = "";
    const unsubSubServerRunning = serverRunning.subscribe((value) => {
        isServerRunning = value;
    });

    function toggleProxy() {
        error = "";

        if (isServerRunning === true) {
            console.log("Stopping proxy");
            stopProxy().then(() => {
                serverRunning.set(false);
                // serverRunning.set(false);
                console.log("Proxy stopped");
            });
        } else {
            const err = validatePort(port);
            if (err != "") {
                error = err;
                return;
            }
            const portNumber = parseInt(port);
            console.log("Starting proxy");
            startProxy(portNumber).then((result) => {
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
        unsubSubServerRunning();
    });
</script>

<h1 class="text-2xl font-medium text-white">{title}</h1>

<div class="mt-10 flex gap-2 mx-auto items-center">
    <input
        readonly={isServerRunning}
        bind:value={port}
        type="text"
        min="0"
        placeholder="8080"
        class="max-w-sm block placeholder-gray-400/70 placeholder-gray-500 rounded-lg border  px-5 py-2.5   focus:outline-none focus:ring focus:ring-blue-300 focus:ring-opacity-40 border-gray-600 bg-gray-900 text-gray-300 focus:border-blue-300"
    />
    <button
        on:click={toggleProxy}
        class="{isServerRunning
            ? stopButtonCss
            : startButtonCss} flex items-center rounded-lg px-4 py-2.5 font-medium tracking-wide text-white capitalize transition-colors duration-300 transform focus:outline-none focus:ring focus:ring-opacity-80"
    >
        {#if isServerRunning}
            Stop {title}
        {:else}
            Start {title}
        {/if}
    </button>
</div>



{#if error != ""}
    <div
        class="mt-4 mx-auto flex w-full max-w-sm overflow-hidden rounded-lg shadow-md bg-gray-800"
    >
        <div class="flex items-center justify-center w-12 bg-red-500">
            <i class="text-white text-xl bi bi-exclamation-triangle"></i>
        </div>

        <div class="px-4 py-2 -mx-3">
            <div class="mx-3">
                <span class="font-semibold  text-red-400"
                    >Error</span
                >
                <p class="text-sm  text-gray-200">
                    {error}
                </p>
            </div>
        </div>
    </div>
{/if}
