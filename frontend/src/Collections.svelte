<script>
    import CollectionsTile from "./CollectionsTile.svelte";
</script>
<div
    class="h-screen py-8 overflow-y-auto  border-l border-r sm:w-64 w-60 min-w-60 bg-gray-900 border-gray-700"
  >
    <!-- <h2 class="px-5 text-lg font-medium text-gray-800 dark:text-white">Accounts</h2> -->

    <div class="space-y-4">
      <CollectionsTile collectionName={"collectionName"} />
    </div>
  </div>