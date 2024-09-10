import { writable } from "svelte/store";
import { REDIRECT_RULE } from "./constants";
import { GetRedirects } from "../wailsjs/go/main/App";


// This is object and not int because integer does not refresh store when assigned same value.
export const currentCollection = writable({collectionId: -1});
export const currentRule = writable(REDIRECT_RULE);
export const redirects = writable([]);
GetRedirects().then((res) => {
    console.log(res);
    redirects.set(res.redirects);
});
