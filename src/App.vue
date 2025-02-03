<script setup lang="ts">
import ThemeToggle from './components/ThemeToggle.vue'
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { Home, Upload, History, Settings, Star, HelpCircle, Menu, X, Crown, CreditCard } from 'lucide-vue-next'

const router = useRouter()
const route = useRoute()

const viralStats = [
  { label: 'Views', value: '1.2M', trend: '+45%' },
  { label: 'Likes', value: '856K', trend: '+62%' },
  { label: 'Shares', value: '234K', trend: '+38%' }
]

const purpleBoxPosition = ref({ x: 0, y: 0 })
const isSidebarOpen = ref(true)

const updatePurpleBoxPosition = () => {
  purpleBoxPosition.value = {
    x: Math.random() * 80,
    y: Math.random() * 80
  }
}

const navItems = [
  { icon: Home, label: 'Dashboard', path: '/' },
  { icon: CreditCard, label: 'Subscription', path: '/subscription' },
  { icon: Settings, label: 'Settings', path: '/settings' },
  { icon: HelpCircle, label: 'Help', path: '/help' },
]

onMounted(() => {
  updatePurpleBoxPosition()
  setInterval(updatePurpleBoxPosition, 3000)
})

const toggleSidebar = () => {
  isSidebarOpen.value = !isSidebarOpen.value
}

const navigateTo = (path: string) => {
  router.push(path)
}
</script>

<template>
  <div class="min-h-screen bg-white dark:bg-gray-900 text-gray-900 dark:text-white transition-colors duration-200 relative overflow-hidden">
    <!-- Animated Background Elements -->
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
      
      <!-- Background Elements (kept from previous version) -->
    </div>

    <!-- Sidebar Toggle Button (Mobile) -->
    <button 
      @click="toggleSidebar"
      class="fixed top-4 left-4 z-50 p-2 rounded-lg bg-gray-200 dark:bg-gray-800 lg:hidden"
    >
      <Menu v-if="!isSidebarOpen" class="w-6 h-6" />
      <X v-else class="w-6 h-6" />
    </button>

    <!-- Side Navigation -->
    <nav 
      class="fixed top-0 left-0 h-full bg-white/80 dark:bg-gray-800/80 backdrop-blur-md w-64 shadow-xl transition-transform duration-300 z-40 transform"
      :class="isSidebarOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0'"
    >
      <div class="p-6">
        <div class="flex items-center gap-3 mb-8">
          
          <h1 class="text-2xl font-bold bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
            ChopTok
          </h1>
        </div>

        <div class="space-y-2">
          <button
            v-for="item in navItems"
            :key="item.label"
            @click="navigateTo(item.path)"
            class="w-full px-4 py-3 rounded-xl flex items-center gap-3 transition-colors duration-200"
            :class="route.path === item.path ? 'bg-primary/10 text-primary' : 'hover:bg-gray-100 dark:hover:bg-gray-700/50'"
          >
            <component :is="item.icon" class="w-5 h-5" />
            <span class="font-medium">{{ item.label }}</span>
          </button>
        </div>
      </div>
      <ThemeToggle />

      <!-- Pro Badge -->
      <div v-if="route.path !== '/upgrade'" class="absolute bottom-8 left-4 right-4">
        <div class="bg-gradient-to-r from-primary to-secondary p-4 rounded-xl text-white">
          <h3 class="font-semibold mb-1">Upgrade to Pro</h3>
          <p class="text-sm opacity-90">Get access to advanced features</p>
          <button 
            @click="navigateTo('/upgrade')"
            class="mt-3 bg-white text-primary px-4 py-1.5 rounded-lg text-sm font-medium hover:bg-gray-50 transition-colors"
          >
            Upgrade Now
          </button>
        </div>
      </div>
    </nav>

    
    <!-- Main Content -->
    <main class="lg:ml-64 transition-all duration-300" :class="{ 'ml-0': !isSidebarOpen }">
      <router-view></router-view>
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