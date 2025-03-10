<script lang="ts">
  import LogDisplay, {
    type LogEntry,
  } from "$/lib/components/base/LogDisplay.svelte";
  import HomeCardList from "$/lib/components/home/HomeCardList.svelte";
  import {
    fetchBuild,
    fetchDeploy,
    fetchPull,
    fetchStartService,
    fetchStopService,
  } from "$/lib/hooks/site.hook.js";
  import { siteStore } from "$/lib/store/siteStore";
  import type { ISiteItem } from "$/lib/types/base";
  import { parseLogs } from "$/lib/utils/helper.js";
  import {
    faEarthAsia,
    faFolder,
    faFolderOpen,
  } from "@fortawesome/free-solid-svg-icons";
  import { tick } from "svelte";
  import { Button } from "svelte-5-ui-lib";
  import Fa from "svelte-fa";

  let { data }: { data: { slug: string } } = $props();
  let resData: ISiteItem | undefined = $state(undefined);
  let logs: LogEntry[] = $state([]);
  let loading = $state(false);
  const handlePull = async (repo: string) => {
    loading = true;
    await fetchPull(
      repo,
      async (msg) => {
        // console.log("SSE Message:", msg);
        logs = [
          ...logs,
          {
            type: "system",
            id: logs.length + 1,
            message: "SSE Message:" + msg,
            timestamp: new Date(),
          },
        ];
        await tick();
      }, // For "data:" events
      async (row) => {
        // console.log("Git Output:", row);
        logs = [...logs, ...parseLogs(row, logs.length)];
        await tick();
      } // For raw git logs
    );
    loading = false;
  };

  const handleBuild = async (repo: string) => {
    loading = true;
    await fetchBuild(
      repo,
      async (msg) => {
        // console.log("SSE Message:", msg);
        logs = [
          ...logs,
          {
            type: "system",
            id: logs.length + 1,
            message: "SSE Message:" + msg,
            timestamp: new Date(),
          },
        ];
        await tick();
      }, // For "data:" events
      async (row) => {
        // console.log("Git Output:", row);
        logs = [...logs, ...parseLogs(row, logs.length)];
        await tick();
      } // For raw git logs
    );
    loading = false;
  };

  const handleDeploy = async (repo: string) => {
    loading = true;
    await fetchDeploy(
      repo,
      async (msg) => {
        // console.log("SSE Message:", msg);
        logs = [
          ...logs,
          {
            type: "system",
            id: logs.length + 1,
            message: "SSE Message:" + msg,
            timestamp: new Date(),
          },
        ];
        await tick();
      }, // For "data:" events
      async (row) => {
        // console.log("Git Output:", row);
        logs = [...logs, ...parseLogs(row, logs.length)];
        await tick();
      } // For raw git logs
    );
    loading = false;
  };

  const handleStartService = async (repo: string) => {
    loading = true;
    await fetchStartService(
      repo,
      async (msg) => {
        // console.log("SSE Message:", msg);
        logs = [
          ...logs,
          {
            type: "system",
            id: logs.length + 1,
            message: "SSE Message:" + msg,
            timestamp: new Date(),
          },
        ];
        await tick();
      }, // For "data:" events
      async (row) => {
        // console.log("Git Output:", row);
        logs = [...logs, ...parseLogs(row, logs.length)];
        await tick();
      } // For raw git logs
    );
    loading = false;
  };

  const handleStopService = async (repo: string) => {
    loading = true;
    await fetchStopService(
      repo,
      async (msg) => {
        // console.log("SSE Message:", msg);
        logs = [
          ...logs,
          {
            type: "system",
            id: logs.length + 1,
            message: "SSE Message:" + msg,
            timestamp: new Date(),
          },
        ];
        await tick();
      }, // For "data:" events
      async (row) => {
        // console.log("Git Output:", row);
        logs = [...logs, ...parseLogs(row, logs.length)];
        await tick();
      } // For raw git logs
    );
    loading = false;
  };

  $effect(() => {
    const unsubscribe = siteStore.subscribe((items) => {
      resData = items.find((item) => item.site_title === data.slug);
    });
    return unsubscribe;
  });
</script>

<div class="flex w-full">
  <HomeCardList siteTitle={data.slug} />
  <div class="min-h-screen ml-56 w-full inline-block pl-0 p-5 bg-gray-800">
    <div class="h-full w-full p-6 rounded-4xl bg-gray-50">
      <h1 class="text-2xl text-orange-600 font-bold">
        {resData?.site_title.toLocaleUpperCase()}
      </h1>
      <div class="grid gap-1 my-5">
        <div class="flex items-center gap-3">
          <Fa icon={faFolder} />
          <p>{resData?.site_folder}</p>
        </div>
        <div class="flex items-center gap-3">
          <Fa icon={faFolderOpen} />
          <p>{resData?.site_deploy}</p>
        </div>
        <!-- <i class="fa-solid fa-earth-asia"></i> -->
        <div class="flex items-center gap-3">
          <Fa icon={faEarthAsia} />
          <a href={resData?.git_repository} class="text-orange-600"
            >{resData?.git_repository}</a
          >
        </div>
      </div>
      <Button
        disabled={loading}
        class="mb-5 cursor-pointer disabled:cursor-progress"
        color="teal"
        type="button"
        onclick={() => handlePull(resData?.site_title ?? "")}
      >
        PULL
      </Button>

      <Button
        disabled={loading}
        class="mb-5 cursor-pointer disabled:cursor-progress"
        color="purple"
        type="button"
        onclick={() => handleBuild(resData?.site_title ?? "")}
      >
        Build
      </Button>

      <Button
        disabled={loading}
        class="mb-5 cursor-pointer disabled:cursor-progress"
        color="orange"
        type="button"
        onclick={() => handleStopService(resData?.site_title ?? "")}
      >
        Stop Service
      </Button>

      <Button
        disabled={loading}
        class="mb-5 cursor-pointer disabled:cursor-progress"
        color="blue"
        type="button"
        onclick={() => handleDeploy(resData?.site_title ?? "")}
      >
        Deploy
      </Button>

      <Button
        disabled={loading}
        class="mb-5 cursor-pointer disabled:cursor-progress"
        color="orange"
        type="button"
        onclick={() => handleStartService(resData?.site_title ?? "")}
      >
        Start Service
      </Button>
      <br />
      <LogDisplay {logs} {loading} />
    </div>
  </div>
</div>
