<script lang="ts">
  import "$/app.css";
  import { login, logout } from "$/lib/utils/auth";
  import { goto } from "$app/navigation";
  import { Button } from "svelte-5-ui-lib";
  import toast, { Toaster } from "svelte-french-toast";

  let formData = $state({
    email: "",
    password: "",
    showPassword: false,
  })

  const handleSubmit = async () => {
    // Login logic
    const res = await login(formData.email, formData.password);

    if (res?.token) {
      goto("/")
    } else if(res?.error) {
      toast.error(res?.error || "");
    }
  };

  $effect(() => {
    logout();
  })
</script>

<div class="flex items-center justify-center min-h-screen bg-gray-100 p-4">
  <div class="w-full max-w-md bg-white shadow-md rounded-lg p-6">
    <h2 class="text-3xl text-center mb-6">Login</h2>

    <form onsubmit={handleSubmit} class="space-y-4">
      <div>
        <label for="email" class="block text-sm font-medium mb-2">Email</label>
        <input
          type="email"
          id="email"
          bind:value={formData.email}
          placeholder="Enter your email"
          required
          class="w-full px-3 py-2 border rounded-md border-gray-300 ring-0"
        />
      </div>

      <div class="relative">
        <label for="password" class="block text-sm font-medium mb-2"
          >Password</label
        >
        <input
          type={formData.showPassword ? "text" : "password"}
          id="password"
          bind:value={formData.password}
          placeholder="Enter your password"
          required
          class="w-full px-3 py-2 border rounded-md pr-10 border-gray-300"
        />
        <button
          type="button"
          onclick={() => (formData.showPassword = !formData.showPassword)}
          class="absolute right-2 top-9 text-gray-500"
        >
          {#if formData.showPassword}
            <span>Hide</span>
          {:else}
            <span>Show</span>
          {/if}
        </button>
      </div>

      <Button
        type="submit"
        class="mt-5 w-full cursor-pointer disabled:cursor-progress"
        color="red"
      >
        Sign In
      </Button>
    </form>
  </div>
</div>
<Toaster position="bottom-right" />

<style>
  /* Additional custom styles if needed */
</style>
