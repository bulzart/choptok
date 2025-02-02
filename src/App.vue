<script setup lang="ts">
import ThemeToggle from './components/ThemeToggle.vue'
import UploadBox from './components/UploadBox.vue'
import Pusher from 'pusher-js';

const pusher = new Pusher('1ac366cc7e1e8ef0dd7c', {
    cluster: 'eu',
});

// Subscribe to a channel
const channel = pusher.subscribe('choptok');

// Bind to an event


const viralStats = [
  { label: 'Views', value: '3.72M', trend: '+45%' },
  { label: 'Likes', value: '856K', trend: '+62%' },
  { label: 'Shares', value: '234K', trend: '+38%' }
]
</script>

<template>
  <div class="min-h-screen bg-white dark:bg-gray-900 text-gray-900 dark:text-white transition-colors duration-200 relative overflow-hidden">
    <!-- Animated Background Elements -->
    <div class="absolute inset-0 overflow-hidden pointer-events-none">
     
      
      <!-- TikTok-style Scrolling Phone Frames -->
      <div class="absolute top-0 -right-0 w-32 h-screen animate-scroll-down">
        <div class="relative w-full h-full">
          <div v-for="i in 1" :key="`phone-${i}`"
               class="absolute w-full aspect-[9/16] rounded-3xl border-8 border-gray-800 dark:border-gray-700 overflow-hidden transform"
               :style="{
                 top: `${i * 50}%`,
                 opacity: 0.3,
                 transform: `rotate(${5 - i}deg)`
               }">
            <div class="w-full h-full bg-gradient-to-br from-pink-500 to-purple-700"></div>
            <div class="absolute inset-0 flex items-center justify-center">
              <span class="text-white text-4xl font-bold">{{ ['🔥', '🚀'][i-1] }}</span>
            </div>
          </div>
        </div>
      </div>


      
      <!-- Viral Stats Charts -->
      <div class="absolute top-20 left-10 space-y-4 opacity-30">
        <div v-for="(stat, index) in viralStats" :key="index"
             class="bg-white/10 backdrop-blur-sm rounded-lg p-4 w-64 transform hover:scale-105 transition-transform"
             :style="{ animationDelay: `${index * 0.2}s` }">
          <div class="flex justify-between items-center">
            <span class="text-sm font-semibold">{{ stat.label }}</span>
            <span class="text-green-400">{{ stat.trend }}</span>
          </div>
          <div class="text-2xl font-bold mt-1">{{ stat.value }}</div>
          <div class="h-2 bg-gray-700 rounded-full mt-2 overflow-hidden">
            <div class="h-full bg-gradient-to-r from-primary to-secondary animate-grow"
                 :style="{ width: `${70 + index * 10}%` }"></div>
          </div>
        </div>
      </div>

         <!-- Floating Social Icons -->
        <!-- Floating Social Icons -->
        <div v-for="i in 8" :key="i" 
           class="absolute animate-float"
           :style="{
             left: `${Math.random() * 100}%`,
             top: `${Math.random() * 100}%`,
             animationDelay: `${Math.random() * 5}s`,
             opacity: '0.15',
             fontSize: `${Math.random() * 2 + 2}rem`
           }">
        <div class="transform" :style="{ transform: `rotate(${Math.random() * 360}deg)` }">
          {{ ['📱', '🎥', '🎬', '🤖', '✨', '🎵', '🎮', '🎪'][i-1] }}
        </div>
      </div>

      <!-- Animated Circles -->
      <div v-for="i in 3" :key="`circle-${i}`"
           class="absolute rounded-full animate-pulse mix-blend-overlay"
           :class="[
             i === 1 ? 'bg-primary/30' : i === 2 ? 'bg-secondary/30' : 'bg-purple-500/30',
             `w-${[96, 80, 64][i-1]} h-${[96, 80, 64][i-1]}`
           ]"
           :style="{
             left: `${Math.random() * 100}%`,
             top: `${Math.random() * 100}%`,
             animationDelay: `${i * 2}s`,
             filter: 'blur(60px)'
           }">
      </div>

    </div>

    <ThemeToggle />
    
    <main class="container mx-auto px-4 py-16 relative">
      <!-- Hero Section -->
      <div class="text-center mb-16">
        <h1 class="text-7xl font-bold mb-6 bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
          ChopTok
        </h1>
        <p class="text-2xl mb-8 max-w-2xl mx-auto">
          Turn Your Long Videos into Viral Shorts – Instantly⚡️
        </p>
        
        <!-- Upload Section -->
        <UploadBox />
      </div>

      <!-- Features Section -->
      <div class="grid md:grid-cols-3 gap-8 max-w-5xl mx-auto mt-16">
        <div class="p-6 rounded-xl bg-gray-50/80 dark:bg-gray-800/80 backdrop-blur-sm">
          <h3 class="text-2xl font-semibold mb-3">AI-Powered Editing ✂️</h3>
          <p class="text-gray-600 dark:text-gray-400">
            Automatically detect key moments and create engaging short clips
          </p>
        </div>
        
        <div class="p-6 rounded-xl bg-gray-50/80 dark:bg-gray-800/80 backdrop-blur-sm">
          <h3 class="text-2xl font-semibold mb-3">Smart Captions 🗒️</h3>
          <p class="text-gray-600 dark:text-gray-400">
            Auto-generated captions with engaging animations
          </p>
        </div>
        
        <div class="p-6 rounded-xl bg-gray-50/80 dark:bg-gray-800/80 backdrop-blur-sm">
          <h3 class="text-2xl font-semibold mb-3">One-Click Share 🔗</h3>
          <p class="text-gray-600 dark:text-gray-400">
            Instantly share to TikTok, Instagram Reels, and YouTube Shorts
          </p>
        </div>
      </div>
    </main>
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

@keyframes pulse-1 {
  0%, 100% { opacity: 0.2; }
  50% { opacity: 0.3; }
}

@keyframes pulse-2 {
  0%, 100% { opacity: 0.15; }
  50% { opacity: 0.25; }
}

@keyframes pulse-3 {
  0%, 100% { opacity: 0.1; }
  50% { opacity: 0.2; }
}

@keyframes pulse-4 {
  0%, 100% { opacity: 0.05; }
  50% { opacity: 0.15; }
}

@keyframes scroll-down {
  0% { transform: translateY(-50%); }
  100% { transform: translateY(0%); }
}

@keyframes scroll-up {
  0% { transform: translateY(0%); }
  100% { transform: translateY(-50%); }
}

@keyframes grow {
  0% { transform: scaleX(0); }
  100% { transform: scaleX(1); }
}

.animate-pulse-1 {
  animation: pulse-1 4s ease-in-out infinite;
}

.animate-pulse-2 {
  animation: pulse-2 4s ease-in-out infinite;
  animation-delay: 1s;
}

.animate-pulse-3 {
  animation: pulse-3 4s ease-in-out infinite;
  animation-delay: 2s;
}

.animate-pulse-4 {
  animation: pulse-4 4s ease-in-out infinite;
  animation-delay: 3s;
}

.animate-scroll-down {
  animation: scroll-down 20s linear infinite;
}

.animate-scroll-up {
  animation: scroll-up 20s linear infinite;
}

.animate-grow {
  animation: grow 2s ease-out forwards;
  transform-origin: left;
}
</style>