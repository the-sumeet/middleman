import { writable } from "svelte/store";
import { RULE_REDIRECT } from "./constants";
import { GetCancels, GetRedirects } from "../wailsjs/go/main/App";
import { HOME } from "./constants";

// This is object and not int because integer does not refresh store when assigned same value.
export const currentRule = writable(RULE_REDIRECT);
export const redirects = writable([]);
GetRedirects().then((res) => {
    console.log("Redirects", res.redirects);
    redirects.set(res.redirects);
});
export const cancels = writable([]);
GetCancels().then((res) => {
    console.log("Cancels", res.cancels);
    cancels.set(res.cancels);
});

export const serverRunning = writable(false);
// Leftmost sidebar buttons
export const currentPage = writable(HOME);