<script>
  export const prerender = true;

  import "$/app.css";
  import PrivateLayout from "$/lib/components/layouts/PrivateLayout.svelte";
  import { Toaster } from "svelte-french-toast";
</script>

<PrivateLayout>
  <slot />
  <Toaster position="bottom-right" />
</PrivateLayout>
