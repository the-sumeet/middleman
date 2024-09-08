import { writable } from "svelte/store";
import { GetCollections } from "../wailsjs/go/main/App";


export const collections = writable([]);
GetCollections().then((value) => {
    collections.set(value);
})

// This is object and not int because integer does not refresh store when assigned same value.
export const currentCollection = writable({collectionId: -1});
