<script>
  export const prerender = true;
  
  import "$/app.css";
  import PrivateLayout from "$/lib/components/layouts/PrivateLayout.svelte";
</script>

<PrivateLayout>
  <slot />
</PrivateLayout>
