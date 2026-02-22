import { ref } from "vue";

export interface MediaDeviceOption {
  deviceId: string;
  label: string;
  kind: MediaDeviceKind;
}

const videoDevices = ref<MediaDeviceOption[]>([]);
const audioDevices = ref<MediaDeviceOption[]>([]);
const audioOutputDevices = ref<MediaDeviceOption[]>([]);
const selectedVideoDeviceId = ref<string>("");
const selectedAudioDeviceId = ref<string>("");

export function useMediaDevices() {
  const enumerateDevices = async () => {
    // Request permissions first so labels are populated
    try {
      const stream = await navigator.mediaDevices.getUserMedia({
        video: true,
        audio: true,
      });
      stream.getTracks().forEach((track) => track.stop());
    } catch {
      // Permission denied or no devices, continue anyway
    }

    const devices = await navigator.mediaDevices.enumerateDevices();

    videoDevices.value = devices
      .filter((d) => d.kind === "videoinput")
      .map((d) => ({
        deviceId: d.deviceId,
        label: d.label || `Camera ${d.deviceId.slice(0, 8)}`,
        kind: d.kind,
      }));

    audioDevices.value = devices
      .filter((d) => d.kind === "audioinput")
      .map((d) => ({
        deviceId: d.deviceId,
        label: d.label || `Microphone ${d.deviceId.slice(0, 8)}`,
        kind: d.kind,
      }));

    audioOutputDevices.value = devices
      .filter((d) => d.kind === "audiooutput")
      .map((d) => ({
        deviceId: d.deviceId,
        label: d.label || `Speaker ${d.deviceId.slice(0, 8)}`,
        kind: d.kind,
      }));

    if (videoDevices.value.length && !selectedVideoDeviceId.value) {
      selectedVideoDeviceId.value = videoDevices.value[0].deviceId;
    }
    if (audioDevices.value.length && !selectedAudioDeviceId.value) {
      selectedAudioDeviceId.value = audioDevices.value[0].deviceId;
    }
  };

  const getStream = async (
    videoDeviceId?: string,
    audioDeviceId?: string,
  ): Promise<MediaStream> => {
    const constraints: MediaStreamConstraints = {
      video: videoDeviceId
        ? {
            deviceId: { exact: videoDeviceId },
            width: { ideal: 2560, max: 2560 },
            height: { ideal: 1440, max: 1440 },
            frameRate: { ideal: 60, max: 60 },
          }
        : false,
      audio: audioDeviceId
        ? {
            deviceId: { exact: audioDeviceId },
            autoGainControl: false,
            channelCount: 2,
            echoCancellation: false,
            noiseSuppression: false,
            sampleRate: 48000,
            sampleSize: 16,
          }
        : false,
    };

    return navigator.mediaDevices.getUserMedia(constraints);
  };

  const getScreenStream = async (): Promise<MediaStream> => {
    return navigator.mediaDevices.getDisplayMedia({
      video: {
        width: { ideal: 2560, max: 2560 },
        height: { ideal: 1440, max: 1440 },
        frameRate: { ideal: 60, max: 60 },
      },
      audio: {
        autoGainControl: false,
        channelCount: 2,
        echoCancellation: false,
        noiseSuppression: false,
        sampleRate: 48000,
        sampleSize: 16,
      },
    });
  };

  return {
    videoDevices,
    audioDevices,
    audioOutputDevices,
    selectedVideoDeviceId,
    selectedAudioDeviceId,
    enumerateDevices,
    getStream,
    getScreenStream,
  };
}
