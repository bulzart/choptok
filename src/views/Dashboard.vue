<script setup lang="ts">
import UploadBox from '../components/UploadBox.vue'
</script>

<template>
  <div class="p-8">
    <!-- Upload Section with Enhanced Animation -->
    <div class="bg-white/80 dark:bg-gray-800/80 backdrop-blur-sm rounded-xl p-8 shadow-lg mb-8 relative overflow-hidden">
      <!-- Animated Social Media Background -->
      <div class="absolute inset-0 pointer-events-none">
        <div v-for="i in 12" :key="i" 
             class="absolute animate-float opacity-5"
             :style="{
               left: `${Math.random() * 100}%`,
               top: `${Math.random() * 100}%`,
               animationDelay: `${i * 0.5}s`,
               transform: `scale(${0.3 + Math.random() * 0.3}) rotate(${Math.random() * 360}deg)`
             }">
          <img 
            :src="[
              'https://static.vecteezy.com/system/resources/previews/016/716/450/non_2x/tiktok-icon-free-png.png',
              'https://upload.wikimedia.org/wikipedia/commons/thumb/e/e7/Instagram_logo_2016.svg/2048px-Instagram_logo_2016.svg.png',
              'https://www.edigitalagency.com.au/wp-content/uploads/Youtube-logo-white-png.png'
            ][i % 3]"
            class="w-16 h-16 object-contain"
            alt="Social Media Icon"
          />
        </div>
      </div>
      
      <UploadBox />
    </div>

    <!-- Features Grid with Animations -->
    <div class="grid md:grid-cols-3 gap-6">
      <div class="group p-6 rounded-xl bg-white/80 dark:bg-gray-800/80 backdrop-blur-sm shadow-lg relative overflow-hidden">
        <div class="absolute inset-0 bg-gradient-to-r from-primary/10 to-secondary/10 transform scale-x-0 group-hover:scale-x-100 transition-transform duration-500 origin-left"></div>
        <div class="relative">
          <h3 class="text-xl font-semibold mb-3 flex items-center gap-2">
            <span class="text-2xl">✂️</span>
            AI-Powered Editing
          </h3>
          <p class="text-gray-600 dark:text-gray-400">
            Automatically detect key moments and create engaging short clips
          </p>
        </div>
      </div>
      
      <div class="group p-6 rounded-xl bg-white/80 dark:bg-gray-800/80 backdrop-blur-sm shadow-lg relative overflow-hidden">
        <div class="absolute inset-0 bg-gradient-to-r from-primary/10 to-secondary/10 transform scale-x-0 group-hover:scale-x-100 transition-transform duration-500 origin-left"></div>
        <div class="relative">
          <h3 class="text-xl font-semibold mb-3 flex items-center gap-2">
            <span class="text-2xl">🗒️</span>
            Smart Captions
          </h3>
          <p class="text-gray-600 dark:text-gray-400">
            Auto-generated captions with engaging animations
          </p>
        </div>
      </div>
      
      <div class="group p-6 rounded-xl bg-white/80 dark:bg-gray-800/80 backdrop-blur-sm shadow-lg relative overflow-hidden">
        <div class="absolute inset-0 bg-gradient-to-r from-primary/10 to-secondary/10 transform scale-x-0 group-hover:scale-x-100 transition-transform duration-500 origin-left"></div>
        <div class="relative">
          <h3 class="text-xl font-semibold mb-3 flex items-center gap-2">
            <span class="text-2xl">🔗</span>
            One-Click Share
          </h3>
          <p class="text-gray-600 dark:text-gray-400">
            Instantly share to TikTok, Instagram Reels, and YouTube Shorts
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
@keyframes float {
  0%, 100% {
    transform: translateY(0) rotate(0);
  }
  50% {
    transform: translateY(-20px) rotate(10deg);
  }
}

.animate-float {
  animation: float 6s ease-in-out infinite;
}
</style>