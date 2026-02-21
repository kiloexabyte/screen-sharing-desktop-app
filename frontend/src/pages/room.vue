<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount, provide } from "vue";
import { useRoute, useRouter } from "vue-router";
import Dialog from "primevue/dialog";
import Button from "primevue/button";
import Chat from "../components/Chat.vue";
import RoomInfo from "../components/RoomInfo.vue";
import { useChatMessage } from "../composables/useChatMessage";
import { useLiveKit } from "../composables/useLiveKit";
import { config } from "../config";
import type { UrlParam } from "../types";

// page data
const localVideo = ref<HTMLMediaElement | null>(null);
const { chatMessages } = useChatMessage();
const currentRoom = ref<string>("");
const currentHost = ref<string>("");
const serverSideStreamingEnabled = ref<boolean>(false);

// UI state
const chatIsOpen = ref(true);
const dialogVisible = ref<boolean>(false);
const failureMessage = ref<string>("");

const {
  hostRoom,
  leaveRoom,
  joinRoom,
  sendMessageLiveKit,
  toggleScreenshareP2P,
  toggleScreenshare,
  cleanUpData,
  participantNames,
  isServerSideStreaming,
} = useLiveKit();

const route = useRoute();
const router = useRouter();

// URL param
const { username, room, isHost, serverSideStreaming } =
  route.query as Partial<UrlParam>;

const leave = async () => {
  await leaveRoom();
  router.push("/");
};

const adjustVolume = (event: KeyboardEvent) => {
  if (!localVideo.value) return;

  const volumeChangeAmount = 0.1;
  switch (event.key) {
    case "ArrowUp":
      localVideo.value.volume = Math.min(
        localVideo.value.volume + volumeChangeAmount,
        1,
      );
      break;
    case "ArrowDown":
      localVideo.value.volume = Math.max(
        localVideo.value.volume - volumeChangeAmount,
        0,
      );
      break;
  }
};

const toggleFullScreen = (): void => {
  if (!localVideo.value) return;

  if (document.fullscreenElement) {
    if (document.exitFullscreen) {
      document.exitFullscreen();
    } else if ((document as any).mozCancelFullScreen) {
      (document as any).mozCancelFullScreen();
    } else if ((document as any).webkitExitFullscreen) {
      (document as any).webkitExitFullscreen();
    } else if ((document as any).msExitFullscreen) {
      (document as any).msExitFullscreen();
    }
  } else if (localVideo.value.requestFullscreen) {
    localVideo.value.requestFullscreen();
  } else if ((localVideo.value as any).mozRequestFullScreen) {
    (localVideo.value as any).mozRequestFullScreen();
  } else if ((localVideo.value as any).webkitRequestFullscreen) {
    (localVideo.value as any).webkitRequestFullscreen();
  } else if ((localVideo.value as any).msRequestFullscreen) {
    (localVideo.value as any).msRequestFullscreen();
  }
};

const preventPlayPause = (event: MouseEvent): void => {
  event.preventDefault();
  toggleFullScreen();
};

const handleToggleStream = async () => {
  if (localVideo.value) {
    if (serverSideStreamingEnabled.value) {
      await toggleScreenshare(localVideo.value);
    } else {
      await toggleScreenshareP2P(localVideo.value);
    }
  }
};

const handleToggleChat = async () => {
  chatIsOpen.value = !chatIsOpen.value;
};

const handleCloseDialog = async () => {
  failureMessage.value = "";
  dialogVisible.value = false;
  router.push("/");
};

onMounted(async () => {
  if (!username || !room) {
    router.push("/");
    return;
  }
  serverSideStreamingEnabled.value = serverSideStreaming === "true";
  isServerSideStreaming.value = serverSideStreaming === "true";

  const res = await fetch(
    `${config.apiBase}/api/livekit/roomCheck?roomName=${room}`,
    { method: "GET" },
  );

  if (!res.ok) {
    dialogVisible.value = true;
    failureMessage.value = "Error Checking if room already exist";
    return;
  }

  const data = await res.json();

  try {
    if (isHost === "true") {
      if (data.roomExist) {
        dialogVisible.value = true;
        failureMessage.value = "Room already exist";
        return;
      }
      await hostRoom(
        room.toString() ?? "",
        username.toString() ?? "",
        serverSideStreamingEnabled.value,
        localVideo.value,
      );
      currentHost.value = username;
      currentRoom.value = room;
    } else {
      if (!data.roomExist) {
        dialogVisible.value = true;
        failureMessage.value = "Room does not exist";
        return;
      }

      if (localVideo.value) {
        const { host } = await joinRoom(
          room.toString() ?? "",
          username.toString() ?? "",
          localVideo.value,
        );
        currentHost.value = host;
        currentRoom.value = room;
      }
    }
  } catch (err: any) {
    dialogVisible.value = true;
    failureMessage.value = err.toString();
  }

  globalThis.addEventListener("keydown", adjustVolume);

  if (localVideo.value) {
    localVideo.value.addEventListener("click", preventPlayPause);
  }
});

onBeforeUnmount(async () => {
  await cleanUpData();
  globalThis.removeEventListener("keydown", adjustVolume);
  if (localVideo.value) {
    localVideo.value.removeEventListener("click", preventPlayPause);
  }
});

provide("sendMessage", sendMessageLiveKit);
provide("handleToggleStream", handleToggleStream);
provide("ToggleChat", handleToggleChat);
provide("leaveRoom", leave);
</script>

<template>
  <div class="h-screen bg-black">
    <div class="flex h-[90%]">
      <video
        autoplay
        playsinline
        ref="localVideo"
        :class="chatIsOpen ? 'w-5/6' : 'w-full'"
        :muted="isHost === 'true'"
      />
      <Chat
        v-if="chatIsOpen"
        :chats="chatMessages"
        :class="chatIsOpen ? 'w-1/6' : 'w-0'"
      />
    </div>

    <div class="h-[10%]">
      <RoomInfo
        :roomName="currentRoom"
        :usernames="participantNames"
        :username="username ?? ''"
        :host="currentHost"
        :isHost="isHost ?? ''"
        :usingSFU="serverSideStreamingEnabled"
      />
    </div>
  </div>

  <Dialog
    v-model:visible="dialogVisible"
    header="Failed to join/host"
    @hide="
      () => {
        dialogVisible = false;
      }
    "
  >
    <p>Hosting/Joining room failed, error message: {{ failureMessage }}</p>
    <Button type="button" @click="handleCloseDialog">Close</Button>
  </Dialog>
</template>
