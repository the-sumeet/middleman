<script>
    export let entity;
    export let setChanged;
</script>
<select
  bind:value={entity}
  on:change={setChanged}
  class=""
  name=""
  id=""
>
  <option class="text-white" value="path">Path</option>
  <option class="text-white" value="host">Host</option>
</select>
