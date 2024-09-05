<script>
    export let collectionName;
</script>
<button
    class="flex space-x-2 items-center w-full px-5 py-2 transition-colors duration-200 dark:hover:bg-gray-800 gap-x-2 hover:bg-gray-100 focus:outline-none"
>
    <i class="text-xl bi bi-send-fill"></i>

    <div class="text-left rtl:text-right">
        <h1
            class="text-sm font-medium text-gray-700 capitalize dark:text-white"
        >
            {collectionName}
        </h1>

        <!-- <p class="text-xs text-gray-500 dark:text-gray-400">11.2 Followers</p> -->
    </div>
</button>
