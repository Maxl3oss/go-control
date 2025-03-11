<script>
  import HomeCardList from "$/lib/components/home/HomeCardList.svelte";
  import { uploadFile } from "$/lib/hooks/site.hook";
  import { Button } from "svelte-5-ui-lib";

	let files = $state();

  const handleSubmit = async() => {
    const response = await uploadFile(files[0]);
    if(response) {
      files = undefined
    }
  }
</script>

<div class="flex w-full">
  <HomeCardList siteTitle={"form"} />
  <div class="min-h-screen ml-56 w-full inline-block pl-0 p-5 bg-gray-800">
    <div class="h-full w-full p-6 rounded-4xl bg-gray-50">
      <label class="block mb-2 text-sm font-medium text-gray-900 dark:text-white" for="file_input">Upload file</label>
      <input bind:files id="file" name="file" type="file" class="p-2 rounded-xl border">

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
    </div>
  </div>
</div>
