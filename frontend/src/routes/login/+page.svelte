<script lang="ts">
  import "$/app.css";
  
  let email = "";
  let password = "";
  let showPassword = false;

  function handleSubmit() {
    // Login logic
    console.log("Login attempt:", { email, password });
  }
</script>

<div class="flex items-center justify-center min-h-screen bg-gray-100 p-4">
  <div class="w-full max-w-md bg-white shadow-md rounded-lg p-6">
    <h2 class="text-2xl text-center mb-6">Login</h2>

    <form on:submit|preventDefault={handleSubmit} class="space-y-4">
      <div>
        <label for="email" class="block text-sm font-medium mb-2">Email</label>
        <input
          type="email"
          id="email"
          bind:value={email}
          placeholder="Enter your email"
          required
          class="w-full px-3 py-2 border rounded-md"
        />
      </div>

      <div class="relative">
        <label for="password" class="block text-sm font-medium mb-2"
          >Password</label
        >
        <input
          type={showPassword ? "text" : "password"}
          id="password"
          bind:value={password}
          placeholder="Enter your password"
          required
          class="w-full px-3 py-2 border rounded-md pr-10"
        />
        <button
          type="button"
          on:click={() => (showPassword = !showPassword)}
          class="absolute right-2 top-9 text-gray-500"
        >
          {#if showPassword}
            <span>Hide</span>
          {:else}
            <span>Show</span>
          {/if}
        </button>
      </div>

      <div class="flex items-center justify-between">
        <div class="flex items-center">
          <input
            type="checkbox"
            id="remember-me"
            class="h-4 w-4 text-indigo-600 rounded"
          />
          <label for="remember-me" class="ml-2 text-sm">Remember me</label>
        </div>
        <a href="" class="text-sm text-indigo-600 hover:text-indigo-500">
          Forgot password?
        </a>
      </div>

      <button
        type="submit"
        class="w-full bg-indigo-600 text-white py-2 rounded-md hover:bg-indigo-700"
      >
        Sign In
      </button>

      <div class="text-center mt-4">
        <p class="text-sm text-gray-600">
          Don't have an account?
          <a href="" class="text-indigo-600 hover:text-indigo-500">
            Sign up
          </a>
        </p>
      </div>
    </form>
  </div>
</div>

<style>
  /* Additional custom styles if needed */
</style>
