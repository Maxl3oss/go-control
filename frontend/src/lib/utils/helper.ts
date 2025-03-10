import type { LogEntry } from "../components/base/LogDisplay.svelte";

export const BASE_URL = import.meta.env.VITE_BACKEND_API_URL ?? "";

function removeAnsiCodes(str: string) {
  return str.replace(/\u001b\[[0-9;]*m/g, '');  // Regex to remove ANSI escape codes
}

function decodeAllUnicodeEscapes(text: string): string {
  return removeAnsiCodes(text.replace(/\\u([0-9a-fA-F]{4})/g, (_, hex) => String.fromCharCode(parseInt(hex, 16))));
}

export function parseLogs(str: string, arrLen: number): LogEntry[] {
  const rawLogs = decodeAllUnicodeEscapes(str);
  const logEntries: LogEntry[] = [];
  const lines = rawLogs.split("\n");

  lines.forEach((line) => {
    const trimmedLine = line.trim();
    if (!trimmedLine) return;

    let type: "info" | "warning" | "error" | "success" | "system" | "cmd" = "info";

    // Determine log type
    if (trimmedLine.includes("✓")) type = "success";
    if (trimmedLine.includes("copied")) type = "success";
    if (trimmedLine.includes("!")) type = "warning";
    if (trimmedLine.includes("ignore")) type = "warning";
    if (trimmedLine.includes("error")) type = "error";
    if (trimmedLine.includes("$")) type = "cmd";

    // Add the entry
    logEntries.push({
      id: (arrLen + 1),
      message: trimmedLine,
      type,
      timestamp: new Date(),
    });
  });

  return logEntries;
}