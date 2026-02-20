export interface ChatMessage {
  username: string;
  text: string;
  time: string;
}

export interface UrlParam {
  username: string;
  room: string;
  isHost: string;
  serverSideStreaming: string;
}
