<script lang="ts">
  import { fetchGetSite } from "$/lib/hooks/site.hook";
  import { siteStore } from "$/lib/store/siteStore";
  import type { ISiteItem } from "$/lib/types/base";

  const { siteTitle } = $props();
  let dataSiteTitle: ISiteItem[] = $state([]);

  const fetchData = async () => {
    const response = await fetchGetSite();
    siteStore.set(response);
  };

  $effect(() => {
    fetchData();
  });

  $effect(() => {
    const unsubscribe = siteStore.subscribe((items) => {
      dataSiteTitle = items;
    });

    return unsubscribe;
  });
</script>

<div
  class="h-screen fixed w-56 rounded-none bg-gray-800 border border-r-0 border-t-0 border-gray-700"
>
  <h3
    class="p-1 pt-6 text-center text-lg font-medium text-gray-900 dark:text-white !rounded-b-none"
  >
    Site
  </h3>

  <a href="/form" class:site-active={siteTitle === "form"} class="sites-item">
    <div class="w-full flex justify-center py-2 rounded-sm bg-orange-700">
      ADD SITE
    </div>
  </a>

  {#each dataSiteTitle as item}
    <a
      href={item.site_title}
      class:site-active={siteTitle === item.site_title}
      class="sites-item"
    >
      {item.site_title.toLocaleUpperCase()}
    </a>
  {/each}
</div>

<style lang="postcss">
  @reference "tailwindcss/theme";

  .site-active {
    @apply bg-gray-700 rounded-md;
  }

  .sites-item {
    @apply text-gray-100 flex gap-2 text-sm font-semibold p-3 m-2 cursor-pointer;
  }
  .sites-item:hover {
    @apply bg-gray-700 rounded-sm;
  }
</style>
