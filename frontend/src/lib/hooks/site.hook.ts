import { BASE_URL } from "../utils/helper";

export const fetchGetSite = async () => {
  try {
    const response = await fetch(`${BASE_URL}/get-site`);
    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    const resData = await response.json();
    return resData.data;
  } catch (err) {
    return [];
  }
};

const eventStreamHelper = async (path: string, repo: string, onMessage: (msg: string) => void, onRawOutput: (msg: string) => void) => {
  try {
    const response = await fetch(`${BASE_URL}/${repo.toLowerCase()}/${path}`, {
      method: 'POST',
      headers: {
        Accept: 'text/event-stream', // Request SSE response
      },
    });

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    const reader = response.body?.getReader();
    const decoder = new TextDecoder();

    if (!reader) {
      throw new Error('ReadableStream not supported');
    }

    let buffer = '';

    while (true) {
      const { done, value } = await reader.read();
      if (done) break;

      buffer += decoder.decode(value, { stream: true });

      // Process each line
      const lines = buffer.split('\n');
      buffer = lines.pop() || ''; // Save the last incomplete line

      for (const line of lines) {
        if (line.startsWith('data:')) {
          const message = line.replace('data:', '').trim();
          onMessage(message); // Callback for SSE messages
        } else if (!line.startsWith('event:') && line.trim() !== '') {
          onRawOutput(line.trim()); // Callback for non-SSE output (Git messages)
        }
      }
    }
  } catch (err) {
    console.error('Error fetching data:', err);
  }
};

export const fetchPull = async (repo: string, onMessage: (msg: string) => void, onRawOutput: (msg: string) => void) => {
  return eventStreamHelper('pull', repo, onMessage, onRawOutput);
};

export const fetchBuild = async (repo: string, onMessage: (msg: string) => void, onRawOutput: (msg: string) => void) => {
  return eventStreamHelper('build', repo, onMessage, onRawOutput);
};

export const fetchStopService = async (repo: string, onMessage: (msg: string) => void, onRawOutput: (msg: string) => void) => {
  return eventStreamHelper('stop-service', repo, onMessage, onRawOutput);
};

export const fetchDeploy = async (repo: string, onMessage: (msg: string) => void, onRawOutput: (msg: string) => void) => {
  return eventStreamHelper('deploy', repo, onMessage, onRawOutput);
};

export const fetchStartService = async (repo: string, onMessage: (msg: string) => void, onRawOutput: (msg: string) => void) => {
  return eventStreamHelper('start-service', repo, onMessage, onRawOutput);
};