<script>
    export let message;
    export let type = ALERT_INFO;
    export let clear;
    export let clearAfter;

    import {
        ALERT_BG_COLOR,
        ALERT_INFO,
        ALERT_ERROR,
        ALERT_SUCCESS,
        ALERT_WARNING,
    } from "../../src/constants";

    const bgColor = ALERT_BG_COLOR[type];

    if (clearAfter) {
        setInterval(() => {
            clear();
        }, clearAfter);
    }
</script>

<div class="w-full text-white {bgColor}">
    <div class="container flex items-center justify-between px-6 py-4 mx-auto">
        <div class="flex">
            {#if type === ALERT_ERROR}
                <i class="w-6 h-6 text-xl fill-current bi bi-lightning-fill"
                ></i>
            {:else if type === ALERT_INFO}
                <i class="w-6 h-6 text-xl fill-current bi bi-info-circle-fill"
                ></i>
            {:else if type === ALERT_WARNING}
                <i
                    class="w-6 h-6 text-xl fill-current bi bi-exclamation-triangle-fill"
                ></i>
            {:else if type === ALERT_SUCCESS}
                <i class="w-6 h-6 text-xl fill-current bi bi-check-circle-fill"
                ></i>
            {/if}

            <p class="mx-3">{message}</p>
        </div>

        <button
            on:click={clear}
            class="p-1 transition-colors duration-300 transform rounded-md hover:bg-opacity-25 hover:bg-gray-600 focus:outline-none"
        >
            <svg
                class="w-5 h-5"
                viewBox="0 0 24 24"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
            >
                <path
                    d="M6 18L18 6M6 6L18 18"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                />
            </svg>
        </button>
    </div>
</div>
