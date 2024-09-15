<script>
  import RuleTile from "./RuleTile.svelte";

</script>
<div
    class="h-screen py-8 overflow-y-auto  border-l border-r sm:w-64 w-60 min-w-60 bg-gray-900 border-gray-700"
  >
    <!-- <h2 class="px-5 text-lg font-medium text-gray-800 dark:text-white">Accounts</h2> -->


    <div class="space-y-4">
      <RuleTile icon={"bi bi-shuffle"} title={"Redirect"}/>
      <RuleTile icon={"bi bi-body-text"} title={"Modify Request Body"}/>
      <RuleTile icon={"bi bi-body-text"} title={"Modify Response Body"}/>
      <RuleTile icon={"bi bi-h-square"} title={"Modify Request Headers"}/>
      <RuleTile icon={"bi bi-h-square"} title={"Modify Response Headers"}/>
      <RuleTile icon={"bi bi-x-square"} title={"Cancel Request"}/>
      <RuleTile icon={"bi bi-clock"} title={"Add Delay"}/>
    </div>
  </div>