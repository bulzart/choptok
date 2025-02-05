<script setup lang="ts">
import { ref } from "vue";
import { Upload, Share2, Download, ArrowUpCircle, Youtube } from "lucide-vue-next";
import axios from "axios";

// Subscribe to a channel
const isHovering = ref(false);
const uploadProgress = ref(0);
const isUploading = ref(false);
const uploadComplete = ref(false);
const youtubeUrl = ref("");
const isYouTubeUploading = ref(false);
const spinningValue = ref("");
const showSubscriptionModal = ref(false);

const previewClips = ref([
  {
    id: 1,
    title: "Biggest Business Mistake? Not Starting NOW! 🚀",
    duration: "0:59",
    thumbnail: "https://picsum.photos/400/450",
  },
  {
    id: 2,
    title: "You're Losing Money By Ignoring This! 💰",
    duration: "0:45",
    thumbnail: "https://picsum.photos/400/451",
  },
  {
    id: 3,
    title: "Stop Doing This If You Want to Be Successful! ❌",
    duration: "0:52",
    thumbnail: "https://picsum.photos/400/452",
  },
]);

const generateIt = () => {
  alert(1);
};

const handleDrop = async(e: DragEvent) => {
  e.preventDefault();
  isHovering.value = false;
  const files = e.dataTransfer?.files;
  if (files && files.length > 0) {
    const file = files[0];
    if (file && ["video/mp4", "video/quicktime", "video/x-msvideo"].includes(file.type)) {
      await uploadChosedFile(file);
    } else {
      alert("Please upload a valid video file (MP4, MOV, or AVI)");
    }
  }
};

const pasteYtUrl = async () => {
  youtubeUrl.value = await navigator.clipboard.readText();
};

const uploadChosedFile = async(file: File) => {
  const formData = new FormData();
  formData.append('file', file);
  isUploading.value = true;
  uploadComplete.value = false;
  try {
    const response = await fetch('http://localhost:8080/upload', {
      method: 'POST',
      body: formData,
    });

    if (response.ok) {
      isUploading.value = false;
      uploadComplete.value = true;
      showSubscriptionModal.value = true; // Show the subscription modal
    } else {
      console.error('File upload failed');
    }
  } catch (error) {
    console.error('Error:', error);
  }
};

const handleFiles = async(files: FileList) => {
  const file = files[0];
  if (file && ['video/mp4', 'video/quicktime', 'video/x-msvideo'].includes(file.type)) {
    await uploadChosedFile(file);
  } else {
    alert('Please upload a valid video file (MP4, MOV, or AVI)');
  }
};

const resetUpload = () => {
  uploadComplete.value = false;
  isUploading.value = false;
  uploadProgress.value = 0;
  youtubeUrl.value = "";
  isYouTubeUploading.value = false;
};

const handleYouTubeUpload = () => {
  isYouTubeUploading.value = true;
  setTimeout(() => {
    isYouTubeUploading.value = false;
    uploadComplete.value = true;
  }, 2000); // Simulating processing time
};

const subscribe = () => {
  // Handle subscription logic here
  alert("Subscribed successfully!");
  showSubscriptionModal.value = false; // Close the modal after subscription
};
</script>

<template>
  <div class="w-full max-w-4xl mx-auto">
    <!-- Upload Box -->
    <div
      @dragover.prevent="isHovering = true"
      @dragleave.prevent="isHovering = false"
      @drop="handleDrop"
      class="w-full p-8 border-2 border-dashed rounded-xl transition-all duration-200 mb-8 backdrop-blur-sm"
      :class="[
        isHovering
          ? 'border-primary bg-primary/5'
          : 'border-gray-300 dark:border-gray-700',
      ]"
    >
      <div
        class="flex flex-col items-center gap-4"
        v-if="!uploadComplete && !isUploading"
      >
        <Upload class="w-12 h-12 text-gray-400" />
        <div class="text-center">
          <h3 class="text-xl font-semibold">Drag & Drop your video here</h3>
          <p class="text-md text-gray-600 dark:text-gray-400">
            Supports MP4, MOV, and AVI files
          </p>
        </div>

        <input
          v-if="!uploadComplete && !isUploading"
          type="file"
          accept="video/mp4,video/quicktime,video/x-msvideo"
          class="hidden"
          @change="(e) => e.target.files && handleFiles(e.target.files)"
          id="file-upload"
        />
        <label
          v-if="!uploadComplete && !isUploading"
          for="file-upload"
          class="btn btn-primary cursor-pointer text-xl"
        >
          Select Video
        </label>
        <div class="text-gray-600 dark:text-gray-400 text-md font-semibold">or</div>
      </div>

      <!-- YouTube Video Upload -->
      <div
        v-if="!uploadComplete && !isUploading"
        class="mt-6 flex flex-col items-center gap-4"
      >
        <h3 class="text-xl flex font-semibold">
          Upload directly from YouTube <Youtube class="w-5 h-5 text-primary my-1 mx-1" />
        </h3>

        <div class="flex items-center w-full max-w-lg">
          <input
            type="text"
            style="text-align: center"
            v-model="youtubeUrl"
            placeholder="Paste YouTube video URL..."
            @click="pasteYtUrl"
            class="input w-full rounded-lg px-4 py-3 border-1 focus:ring-primary transition-all"
          />
        </div>
        <input
          v-if="!uploadComplete && !isUploading"
          type="button"
          class="hidden"
          @click="generateIt"
        />
        <label
          v-if="!uploadComplete && !isUploading"
          class="btn btn-primary cursor-pointer text-xl"
        >
          Generate
        </label>
      </div>

      <div v-if="isUploading || isYouTubeUploading" class="mt-6 flex flex-col items-center">
        <div class="spinner-border animate-spin inline-block w-16 h-16 border-4 rounded-full border-t-primary border-solid" role="status">
          <span class="sr-only">Loading...</span>
        </div>
        <p class="text-sm text-center mt-2">Processing video...</p>
      </div>

      <div class="mt-5">
        <h3 class="text-2xl font-bold mb-6">Generate Clips⤵</h3>

        <div class="grid md:grid-cols-3 gap-3">
          <div
            v-for="clip in previewClips"
            :key="clip.id"
            class="bg-gray-50/80 dark:bg-gray-800/80 backdrop-blur-sm rounded-xl overflow-hidden shadow-lg hover:shadow-xl transition-shadow"
          >
            <div class="relative">
              <img :src="clip.thumbnail" :alt="clip.title" class="w-full object-cover" />
              <span
                class="absolute bottom-2 right-2 bg-black/70 text-white px-2 py-1 rounded text-sm"
              >
                {{ clip.duration }}
              </span>
            </div>
            <div class="p-4">
              <h3 class="font-semibold mb-3">{{ clip.title }}</h3>
              <div class="flex gap-2">
                <button
                  class="btn btn-primary py-2 flex-1 flex items-center justify-center gap-2"
                >
                  <Share2 class="w-4 h-4" />
                  Share
                </button>
                <button
                  class="btn bg-gray-200 dark:bg-gray-700 py-2 px-3 hover:bg-gray-300 dark:hover:bg-gray-600"
                >
                  <Download class="w-4 h-4" />
                </button>
              </div>
            </div>
          </div>
        </div>

        <div class="flex justify-center mt-12" v-if="uploadComplete">
          <button
            @click="resetUpload"
            class="group flex items-center gap-2 px-6 py-3 bg-gradient-to-r from-primary to-secondary text-white rounded-lg hover:opacity-90 transition-all duration-300 shadow-lg hover:shadow-xl"
          >
            <ArrowUpCircle
              class="w-5 h-5 group-hover:rotate-180 transition-transform duration-500"
            />
            Upload Another Video
          </button>
        </div>
      </div>
    </div>

    <!-- Enhanced Subscription Modal -->
    <div
  v-if="showSubscriptionModal"
  class="fixed inset-0 flex items-center justify-center z-50"
>
  <div
    class="modal-overlay absolute inset-0 bg-black bg-opacity-50 backdrop-blur-sm"
    @click="showSubscriptionModal = false"
  ></div>
  <div
    class="modal-container bg-white dark:bg-gray-800 p-8 rounded-2xl shadow-2xl w-full max-w-2xl mx-auto transition-transform transform scale-100"
  >
    <h2 class="text-3xl font-bold mb-4">Subscribe to Continue✅</h2>
    <p class="dark:text-gray-300 mb-6 text-lg">
      Unlock all features by subscribing for just $14.99/month.
    </p>
    <div class="mb-6">
      <h3 class="text-xl font-semibold mb-2">Subscription Features:</h3>
      <ul class="list-disc list-inside text-gray-600 dark:text-gray-400">
        <li>Video uploads & Generations</li>
        <li>High-quality video processing</li>
        <li>Priority support</li>
        <li>Exclusive content</li>
      </ul>
    </div>
    <div class="flex flex-col items-center gap-4">
      <button @click="subscribe" class="btn btn-primary">
        Subscribe Now
      </button>
    </div>
  </div>
</div>
      </div>
</template>
