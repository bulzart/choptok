<script setup lang="ts">
import ThemeToggle from './components/ThemeToggle.vue'
import { ref } from 'vue'
import { Menu, Home, Video, TrendingUp, Settings, HelpCircle, LogOut,CreditCard } from 'lucide-vue-next'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from './stores/auth'

const router = useRouter()
const route = useRoute()
const auth = useAuthStore()

const purpleBoxPosition = ref({ x: 0, y: 0 })
const isSidebarCollapsed = ref(false)

const updatePurpleBoxPosition = () => {
  purpleBoxPosition.value = {
    x: Math.random() * 80,
    y: Math.random() * 80
  }
}

const navItems = [
  { icon: Home, label: 'Dashboard', path: '/' },
  { icon: Video, label: 'My Videos', path: '/videos' },
  { icon: CreditCard, label: 'Subscription', path: '/subscriptions' },
  { icon: Settings, label: 'Settings', path: '/settings' },
  { icon: HelpCircle, label: 'Help', path: '/help' },
]

setInterval(updatePurpleBoxPosition, 3000)
</script>

<template>
  <div class="min-h-screen bg-white dark:bg-gray-900 text-gray-900 dark:text-white transition-colors duration-200 flex">
    <!-- Only show sidebar if user is authenticated -->
    <aside 
      v-if="auth.isAuthenticated"
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
            <router-link 
              :to="item.path"
              class="flex items-center gap-3 px-3 py-2 rounded-lg transition-colors"
              :class="route.path === item.path ? 'bg-primary/10 text-primary' : 'hover:bg-gray-200 dark:hover:bg-gray-700'"
            >
              <component :is="item.icon" class="w-5 h-5" />
              <span :class="isSidebarCollapsed ? 'hidden' : 'block'">{{ item.label }}</span>
            </router-link>
          </li>
          <!-- Logout Button -->
          <li>
            <button 
              @click="auth.signOut"
              class="w-full flex items-center gap-3 px-3 py-2 rounded-lg transition-colors hover:bg-gray-200 dark:hover:bg-gray-700 text-red-500"
            >
              <LogOut class="w-5 h-5" />
              <span :class="isSidebarCollapsed ? 'hidden' : 'block'">Logout</span>
            </button>
          </li>
        </ul>
      </nav>
    </aside>

    <!-- Main Content -->
    <main 
      class="flex-1 transition-all duration-300 overflow-hidden"
      :class="auth.isAuthenticated ? (isSidebarCollapsed ? 'ml-20' : 'ml-64') : ''"
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
        <ThemeToggle />
        <router-view></router-view>
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