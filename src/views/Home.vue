<script setup lang="ts">
import UploadBox from '../components/UploadBox.vue'
</script>

<template>
  <div class="container mx-auto px-4 py-16 relative">
    <div class="text-center mb-16">
      <h2 class="text-5xl font-bold mb-6 bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
        Turn Your Long Videos into Viral Shorts
      </h2>
      <p class="text-2xl mb-8 max-w-2xl mx-auto opacity-80">
        Create engaging short-form content in seconds with AI-powered editing
      </p>
      
      <UploadBox />
    </div>

    <!-- Features Section -->
    <div class="grid md:grid-cols-3 gap-8 max-w-5xl mx-auto mt-16">
      <div class="p-6 rounded-xl bg-gray-50/80 dark:bg-gray-800/80 backdrop-blur-sm border border-gray-200 dark:border-gray-700">
        <div class="w-12 h-12 bg-primary/10 rounded-xl flex items-center justify-center mb-4">
          <span class="text-2xl">✂️</span>
        </div>
        <h3 class="text-xl font-semibold mb-3">AI-Powered Editing</h3>
        <p class="text-gray-600 dark:text-gray-400">
          Automatically detect key moments and create engaging short clips
        </p>
      </div>
      
      <div class="p-6 rounded-xl bg-gray-50/80 dark:bg-gray-800/80 backdrop-blur-sm border border-gray-200 dark:border-gray-700">
        <div class="w-12 h-12 bg-primary/10 rounded-xl flex items-center justify-center mb-4">
          <span class="text-2xl">🗒️</span>
        </div>
        <h3 class="text-xl font-semibold mb-3">Smart Captions</h3>
        <p class="text-gray-600 dark:text-gray-400">
          Auto-generated captions with engaging animations
        </p>
      </div>
      
      <div class="p-6 rounded-xl bg-gray-50/80 dark:bg-gray-800/80 backdrop-blur-sm border border-gray-200 dark:border-gray-700">
        <div class="w-12 h-12 bg-primary/10 rounded-xl flex items-center justify-center mb-4">
          <span class="text-2xl">🔗</span>
        </div>
        <h3 class="text-xl font-semibold mb-3">One-Click Share</h3>
        <p class="text-gray-600 dark:text-gray-400">
          Instantly share to TikTok, Instagram Reels, and YouTube Shorts
        </p>
      </div>
    </div>
  </div>
</template>