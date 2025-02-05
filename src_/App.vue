<script setup lang="ts">
import ThemeToggle from './components/ThemeToggle.vue'
import UploadBox from './components/UploadBox.vue'
import { ref, onMounted } from 'vue'
import { Menu, Home, Video, TrendingUp, Settings, HelpCircle } from 'lucide-vue-next'

const viralStats = [
  { label: 'Views', value: '1.2M', trend: '+45%' },
  { label: 'Likes', value: '856K', trend: '+62%' },
  { label: 'Shares', value: '234K', trend: '+38%' }
]

const purpleBoxPosition = ref({ x: 0, y: 0 })
const isSidebarCollapsed = ref(false)

const updatePurpleBoxPosition = () => {
  purpleBoxPosition.value = {
    x: Math.random() * 80,
    y: Math.random() * 80
  }
}

const navItems = [
  { icon: Home, label: 'Dashboard', active: true },
  { icon: Video, label: 'My Videos', active: false },
  { icon: TrendingUp, label: 'Analytics', active: false },
  { icon: Settings, label: 'Settings', active: false },
  { icon: HelpCircle, label: 'Help', active: false },
]

onMounted(() => {
  updatePurpleBoxPosition()
  setInterval(updatePurpleBoxPosition, 3000)
})
</script>

<template>
  <div class="min-h-screen bg-white dark:bg-gray-900 text-gray-900 dark:text-white transition-colors duration-200 flex">
    <ThemeToggle />
    <!-- Sidebar -->
    <aside 
      class="fixed left-0 top-0 h-full bg-gray-50 dark:bg-gray-800 transition-all duration-300 z-50 border-r border-gray-200 dark:border-gray-700"
      :class="isSidebarCollapsed ? 'w-20' : 'w-64'"
    >
      <!-- Logo -->
      <div class="h-16 flex items-center justify-between px-4 border-b border-gray-200 dark:border-gray-700">
        <h1 
          class="font-bold text-xl bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent"
          :class="isSidebarCollapsed ? 'hidden' : 'block'"
        >
          ChopTok
        </h1>
        <button 
          @click="isSidebarCollapsed = !isSidebarCollapsed"
          class="p-2 rounded-lg hover:bg-gray-200 dark:hover:bg-gray-700 transition-colors"
        >
          <Menu class="w-5 h-5" />
        </button>
      </div>

      <!-- Navigation -->
      <nav class="p-4">
        <ul class="space-y-2">
          <li v-for="item in navItems" :key="item.label">
            <a 
              href="#" 
              class="flex items-center gap-3 px-3 py-2 rounded-lg transition-colors"
              :class="item.active ? 'bg-primary/10 text-primary' : 'hover:bg-gray-200 dark:hover:bg-gray-700'"
            >
              <component :is="item.icon" class="w-5 h-5" />
              <span :class="isSidebarCollapsed ? 'hidden' : 'block'">{{ item.label }}</span>
            </a>
          </li>
        </ul>
      </nav>
    </aside>

    <!-- Main Content -->
    <main 
      class="flex-1 transition-all duration-300 overflow-hidden"
      :class="isSidebarCollapsed ? 'ml-20' : 'ml-64'"
    >
      <!-- Background Elements -->
      <div class="absolute inset-0 overflow-hidden pointer-events-none">
        <!-- Random Purple Box -->
        <div 
          class="absolute w-32 h-32 bg-purple-600/20 rounded-lg blur-xl animate-pulse"
          :style="{
            left: `${purpleBoxPosition.x}%`,
            top: `${purpleBoxPosition.y}%`,
            transition: 'all 1.5s ease-in-out'
          }"
        ></div>
        
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

      <div class="relative">
       
        
        <!-- Dashboard Content -->
        <div class="p-8">
          <!-- Stats Grid -->
        

          <!-- Upload Section -->
          <div class="bg-white/80 dark:bg-gray-800/80 backdrop-blur-sm rounded-xl p-8 shadow-lg mb-8">
            <UploadBox />
          </div>

          <!-- Features Grid -->
          <div class="grid md:grid-cols-3 gap-6">
            <div class="p-6 rounded-xl bg-white/80 dark:bg-gray-800/80 backdrop-blur-sm shadow-lg">
              <h3 class="text-xl font-semibold mb-3">AI-Powered Editing ✂️</h3>
              <p class="text-gray-600 dark:text-gray-400">
                Automatically detect key moments and create engaging short clips
              </p>
            </div>
            
            <div class="p-6 rounded-xl bg-white/80 dark:bg-gray-800/80 backdrop-blur-sm shadow-lg">
              <h3 class="text-xl font-semibold mb-3">Smart Captions 🗒️</h3>
              <p class="text-gray-600 dark:text-gray-400">
                Auto-generated captions with engaging animations
              </p>
            </div>
            
            <div class="p-6 rounded-xl bg-white/80 dark:bg-gray-800/80 backdrop-blur-sm shadow-lg">
              <h3 class="text-xl font-semibold mb-3">One-Click Share 🔗</h3>
              <p class="text-gray-600 dark:text-gray-400">
                Instantly share to TikTok, Instagram Reels, and YouTube Shorts
              </p>
            </div>
          </div>
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

@keyframes grow {
  0% { transform: scaleX(0); }
  100% { transform: scaleX(1); }
}

.animate-grow {
  animation: grow 2s ease-out forwards;
  transform-origin: left;
}
</style>