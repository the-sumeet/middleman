import { writable } from "svelte/store";
import { RULE_REDIRECT } from "./constants";
import { GetCancels, GetDelays, GetRedirects } from "../wailsjs/go/main/App";
import { HOME } from "./constants";

// This is object and not int because integer does not refresh store when assigned same value.
export const currentRule = writable(RULE_REDIRECT);
export const redirects = writable([]);
GetRedirects().then((res) => {
    redirects.set(res.redirects);
});
export const cancels = writable([]);
GetCancels().then((res) => {
    cancels.set(res.cancels);
});
export const delays = writable([]);
GetDelays().then((res) => {
    delays.set(res.delays);
});
export const serverRunning = writable(false);
// Leftmost sidebar buttons
export const currentPage = writable(HOME);