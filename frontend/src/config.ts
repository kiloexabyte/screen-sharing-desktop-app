// App configuration - can be overridden via environment or Go backend
export const config = {
  livekitWsUrl:
    import.meta.env.VITE_LIVEKIT_WS_URL ||
    "wss://screensharing-web-app-nxkeok9z.livekit.cloud",
  apiBase:
    import.meta.env.VITE_API_BASE ||
    "https://screen-sharing-web-app-ebon.vercel.app",
};
