import { writable } from 'svelte/store';
import type { ISiteItem } from '../types/base';

export let siteStore = writable<ISiteItem[]>([]);
