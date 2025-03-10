<script lang="ts">
  import {
    faCircleCheck,
    faCircleNodes,
    faFaceSmile,
    faSortDown,
  } from "@fortawesome/free-solid-svg-icons";
  import Fa from "svelte-fa";

  export type LogEntry = {
    id: number;
    message: string;
    type: "info" | "warning" | "error" | "success" | "system" | "cmd";
    timestamp: Date;
  };

  let element: HTMLDivElement | null = $state(null);
  let { logs, loading } = $props<{ logs: LogEntry[]; loading: boolean }>();
  let dataLogs: LogEntry[] = $state([]);

  // Computed property to format timestamp
  function formatTimestamp(date: Date) {
    return date.toLocaleTimeString("en-US", {
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit",
    });
  }

  const scrollToBottom = (node: HTMLDivElement) => {
    node.scroll({ top: node.scrollHeight, behavior: "smooth" });
  };

  $effect(() => {
    dataLogs = logs;
  });
</script>

<div class="relative log-display rounded-md">
  <div bind:this={element} class="h-[32rem] overflow-y-auto">
    {#each dataLogs as log (log.id)}
      <div
        class="log-entry"
        class:info={log.type === "info"}
        class:warning={log.type === "warning"}
        class:error={log.type === "error"}
        class:success={log.type === "success"}
        class:system={log.type === "system"}
        class:cmd={log.type === "cmd"}
      >
        <span class="timestamp">{formatTimestamp(log.timestamp)}</span>
        <span class="message">{log.message}</span>
      </div>
    {/each}
  </div>
  <div class="absolute bottom-5 left-5 text-orange-600">
    {#if dataLogs.length <= 0}
      <Fa icon={faFaceSmile} class="text-xl" />
    {:else if loading}
      <Fa icon={faCircleNodes} class="text-xl animate-spin" />
    {:else}
      <Fa icon={faCircleCheck} class="text-xl" />
    {/if}
  </div>

  <div
    class="absolute bottom-5 right-5 red-500 flex items-center justify-center w-8 h-8 bg-orange-600 hover:bg-orange-400 cursor-pointer text-white rounded-full"
  >
    <button type="button" onclick={() => element && scrollToBottom(element)}>
      <Fa icon={faSortDown} />
    </button>
  </div>
</div>

<style lang="postcss">
  .log-display {
    flex-grow: 1;
    overflow-y: auto;
    padding: 10px;
    background-color: #2d2d2d;
  }

  .log-entry {
    display: flex;
    align-items: center;
    margin-bottom: 5px;
  }

  .timestamp {
    margin-right: 10px;
    color: #777;
    font-size: 0.8em;
  }

  .message {
    flex-grow: 1;
  }

  .info {
    color: #e0e0e0;
  }
  .warning {
    color: #ffa500;
  }
  .error {
    color: #ff4136;
  }
  .success {
    color: #2ecc40;
  }
  .system {
    color: #a1887f;
  }
  .cmd {
    color: #673ab7;
  }
</style>
