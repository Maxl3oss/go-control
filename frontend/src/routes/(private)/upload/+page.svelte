<script>
  import HomeCardList from "$/lib/components/home/HomeCardList.svelte";
  import { uploadFile } from "$/lib/hooks/site.hook";
  import { Button, Input } from "svelte-5-ui-lib";

  let formData = $state({
    pass: "",
  });
  let files = $state();
  let locked = $state(true);

  const handleSubmit = async () => {
    const response = await uploadFile(files[0]);
    if (response) {
      files = undefined;
    }
  };

  const handleUnLock = () => {
    locked = formData.pass != "HC-PIPELINE";
  };
</script>

<div class="flex w-full">
  <HomeCardList siteTitle={"form"} />
  <div class="min-h-screen ml-56 w-full inline-block pl-0 p-5 bg-gray-800">
    <div class={`h-full w-full p-6 rounded-4xl ${locked ? "bg-gray-400" : "bg-white"}`}>
      {#if locked}
        {@render uiLock()}
      {:else if upload}
        {@render upload()}
      {/if}
    </div>
  </div>
</div>

{#snippet uiLock()}
  <div class="h-full flex gap-2 justify-center items-center">
    <Input type="password" name="pass" bind:value={formData.pass} />
    <Button
      type="submit"
      class="cursor-pointer disabled:cursor-progress"
      color="red"
      onclick={handleUnLock}
    >
      UNLOCK
    </Button>
  </div>
{/snippet}

{#snippet upload()}
  <label
    class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
    for="file_input">Upload file</label
  >
  <input
    bind:files
    id="file"
    name="file"
    type="file"
    class="p-2 rounded-md border"
  />

  <Button
    type="submit"
    class="mb-5 cursor-pointer disabled:cursor-progress"
    color="red"
    onclick={handleSubmit}
  >
    Upload
  </Button>

  {#if files}
    <h2>Selected files:</h2>
    {#each Array.from(files) as file}
      <p>{file.name} ({file.size} bytes)</p>
    {/each}
  {/if}
{/snippet}
