// App configuration - can be overridden via environment or Go backend
export const config = {
  livekitWsUrl: import.meta.env.VITE_LIVEKIT_WS_URL || "wss://your-livekit-server.com",
  apiBase: import.meta.env.VITE_API_BASE || "https://screen-sharing-web-app-ebon.vercel.app",
};
