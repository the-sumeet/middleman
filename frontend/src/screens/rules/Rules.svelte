<script>
  import {
    RULE_CANCEL,
    RULE_DELAY,
    RULE_MODIFY_REQUEST_BODY,
    RULE_MODIFY_RESPONSE_BODY,
    RULE_MOD_REQUEST_HEADER,
    RULE_MOD_RESPONSE_HEADER,
    RULE_REDIRECT,
  } from "../../../src/constants";
  import RuleTile from "./RuleTile.svelte";
</script>

<div
  class="h-screen overflow-y-auto border-l border-r sm:w-64 w-60 min-w-60 bg-gray-900 border-gray-700"
>
  <!-- <h2 class="px-5 text-lg font-medium text-white">Accounts</h2> -->

  <div class="space-y-4">
    <RuleTile icon={"bi bi-shuffle"} title={"Redirect"} rule={RULE_REDIRECT} />
    <RuleTile
      icon={"bi bi-body-text"}
      title={"Modify Request Body"}
      rule={RULE_MODIFY_REQUEST_BODY}
    />
    <RuleTile
      icon={"bi bi-body-text"}
      title={"Modify Response Body"}
      rule={RULE_MODIFY_RESPONSE_BODY}
    />
    <RuleTile
      icon={"bi bi-h-square"}
      title={"Modify Request Headers"}
      rule={RULE_MOD_REQUEST_HEADER}
    />
    <RuleTile
      icon={"bi bi-h-square"}
      title={"Modify Response Headers"}
      rule={RULE_MOD_RESPONSE_HEADER}
    />
    <RuleTile
      icon={"bi bi-x-square"}
      title={"Cancel Request"}
      rule={RULE_CANCEL}
    />
    <RuleTile icon={"bi bi-clock"} title={"Add Delay"} rule={RULE_DELAY} />
  </div>
</div>
