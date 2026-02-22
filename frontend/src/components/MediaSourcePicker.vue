<script setup lang="ts">
import { onMounted, ref } from "vue";
import Dialog from "primevue/dialog";
import Button from "primevue/button";
import {
  useMediaDevices,
  type MediaDeviceOption,
} from "../composables/useMediaDevices";

const emit = defineEmits<{
  (e: "select", stream: MediaStream): void;
  (e: "close"): void;
}>();

const props = defineProps<{
  visible: boolean;
}>();

const {
  videoDevices,
  audioDevices,
  selectedVideoDeviceId,
  selectedAudioDeviceId,
  enumerateDevices,
  getStream,
  getScreenStream,
} = useMediaDevices();

const error = ref("");

onMounted(async () => {
  await enumerateDevices();
});

const startScreenShare = async () => {
  error.value = "";
  try {
    const stream = await getScreenStream();
    emit("select", stream);
  } catch (err: any) {
    error.value = `Screen share failed: ${err.message}`;
  }
};

const startCamera = async () => {
  error.value = "";
  try {
    const stream = await getStream(
      selectedVideoDeviceId.value,
      selectedAudioDeviceId.value,
    );
    emit("select", stream);
  } catch (err: any) {
    error.value = `Camera failed: ${err.message}`;
  }
};
</script>

<template>
  <Dialog
    :visible="props.visible"
    header="Select Media Source"
    :modal="true"
    :closable="true"
    @update:visible="(val: boolean) => !val && emit('close')"
    :style="{ width: '450px' }"
  >
    <div class="flex flex-col gap-4">
      <div>
        <Button
          @click="startScreenShare"
          class="w-full"
          severity="info"
          outlined
        >
          Share Screen
        </Button>
      </div>

      <div class="border-t pt-4">
        <h3 class="mb-2 font-semibold">Or use camera</h3>

        <div v-if="videoDevices.length" class="mb-3">
          <label class="mb-1 block text-sm">Video</label>
          <select
            v-model="selectedVideoDeviceId"
            class="w-full rounded border p-2 text-sm"
          >
            <option
              v-for="device in videoDevices"
              :key="device.deviceId"
              :value="device.deviceId"
            >
              {{ device.label }}
            </option>
          </select>
        </div>
        <div v-else class="mb-3 text-sm text-gray-500">No cameras detected</div>

        <div v-if="audioDevices.length" class="mb-3">
          <label class="mb-1 block text-sm">Audio</label>
          <select
            v-model="selectedAudioDeviceId"
            class="w-full rounded border p-2 text-sm"
          >
            <option
              v-for="device in audioDevices"
              :key="device.deviceId"
              :value="device.deviceId"
            >
              {{ device.label }}
            </option>
          </select>
        </div>
        <div v-else class="mb-3 text-sm text-gray-500">
          No microphones detected
        </div>

        <Button
          @click="startCamera"
          class="w-full"
          :disabled="!videoDevices.length && !audioDevices.length"
          outlined
        >
          Start Camera
        </Button>
      </div>

      <div v-if="error" class="text-sm text-red-500">{{ error }}</div>
    </div>
  </Dialog>
</template>
