import { writable } from "svelte/store";
import { REDIRECT_RULE, REQUESTS } from "./constants";
import { GetRedirects } from "../wailsjs/go/main/App";
import { HOME } from "./constants";

// This is object and not int because integer does not refresh store when assigned same value.
export const currentCollection = writable({collectionId: -1});
export const currentRule = writable(REDIRECT_RULE);
export const redirects = writable([]);
GetRedirects().then((res) => {
    console.log(res);
    redirects.set(res.redirects);
});

export const serverRunning = writable(false);
// Leftmost sidebar buttons
export const currentPage = writable(REQUESTS);