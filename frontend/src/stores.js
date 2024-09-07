import { writable } from "svelte/store";
import { GetCollections } from "../wailsjs/go/main/App";


export const collections = writable([]);
GetCollections().then((value) => {
    collections.set(value);
})

export const currentCollection = writable(-1);
