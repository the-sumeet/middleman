<script>
    export let title;
    export let description;
    export let port;
    export let addPortFunc;

    // Port
    let portChanged = false;
    let portError = "";

    function savePort() {
        if (port.length != 4) {
            portError = "Port must be of length 4";
            return;
        }
        addPortFunc(port).then((result) => {
            if (result.error != "") {
                portError = result.error;
            } else {
                portError = "";
                portChanged = false;
            }
        });
    }
</script>

<div class="w-full mt-8">

    <h1 class="text-white font-bold text-xl">{title}</h1>

        <div class="mt-4">
            <label
                for="username"
                class="block text-sm text-gray-300"
                >{description}</label
            >

            <input
                bind:value={port}
                on:input={() => (portChanged = true)}
                type="text"
                placeholder="Certificate Path"
                class="block mt-2 w-full placeholder-gray-400/70 placeholder-gray-500 rounded-lg border  px-5 py-2.5 focus:outline-none focus:ring-opacity-40 border-gray-600 bg-gray-900 text-gray-300 focus:border-1"
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
