<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '../../stores/auth'

const auth = useAuthStore()
const email = ref('')
const password = ref('')
const error = ref('')
const isLoading = ref(false)

async function handleSubmit() {
  try {
    error.value = ''
    isLoading.value = true
    await auth.signIn(email.value, password.value)
  } catch (e: any) {
    error.value = e.message
  } finally {
    isLoading.value = false
  }
}
</script>

<template>
  <div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900 py-12 px-4 sm:px-6 lg:px-8 relative overflow-hidden">
    <!-- Animated Background Elements -->
    <div class="absolute inset-0 overflow-hidden">
      <!-- Animated Gradient Orbs -->
      <div v-for="i in 3" :key="`orb-${i}`"
           class="absolute rounded-full mix-blend-overlay animate-float"
           :style="{
             width: `${200 + i * 100}px`,
             height: `${200 + i * 100}px`,
             background: i === 1 ? 'radial-gradient(circle, rgba(100,108,255,0.3) 0%, rgba(0,0,0,0) 70%)' :
                        i === 2 ? 'radial-gradient(circle, rgba(66,184,131,0.3) 0%, rgba(0,0,0,0) 70%)' :
                                 'radial-gradient(circle, rgba(147,51,234,0.3) 0%, rgba(0,0,0,0) 70%)',
             left: `${(i * 30) - 20}%`,
             top: `${(i * 20) - 10}%`,
             animationDelay: `${i * 0.5}s`,
             filter: 'blur(60px)'
           }">
      </div>

      <!-- Social Media Icons Animation -->
      <div v-for="i in 5" :key="`icon-${i}`" 
           class="absolute animate-float opacity-10"
           :style="{
             left: `${Math.random() * 100}%`,
             top: `${Math.random() * 100}%`,
             animationDelay: `${i * 0.5}s`,
             transform: `scale(${0.5 + Math.random() * 0.5})`
           }">
        <img 
          :src="[
            'https://static.vecteezy.com/system/resources/thumbnails/025/225/149/small_2x/3d-illustration-icon-of-camera-film-video-with-circular-or-round-podium-png.png',
            'https://freesvg.org/img/kdisknav.png',
            'https://static.vecteezy.com/system/resources/thumbnails/025/225/149/small_2x/3d-illustration-icon-of-camera-film-video-with-circular-or-round-podium-png.png',
            'https://freesvg.org/img/kdisknav.png',
            'https://storage.needpix.com/rsynced_images/multimedia-98385_1280.png'
          ][i -1]"
          class="w-16 h-16 object-contain"
          alt="Social Media Icon"
        />
      </div>
    </div>

    <!-- Login Form -->
    <div class="max-w-md w-full relative">
      <!-- Glassmorphism Card -->
      <div class="backdrop-blur-xl bg-white/10 p-8 rounded-2xl shadow-2xl border border-white/20">
        <!-- Logo Animation -->
        <div class="text-center mb-8">
          <div class="relative inline-block">
            <div class="absolute inset-0 bg-gradient-to-r from-primary via-purple-500 to-secondary blur-2xl opacity-30 animate-pulse"></div>
            <h1 class="relative text-5xl font-bold bg-gradient-to-r from-primary via-purple-500 to-secondary bg-clip-text text-transparent animate-gradient">
              ChopTok
            </h1>
          </div>
          <p class="mt-4 text-gray-300 text-lg">Transform your content into viral shorts</p>
        </div>

        <form @submit.prevent="handleSubmit" class="space-y-6">
          <!-- Email Input -->
          <div class="group">
            <div class="relative">
              <input
                type="email"
                v-model="email"
                required
                class="w-full px-4 py-3 bg-gray-800/50 border border-gray-700 rounded-lg focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all duration-300 text-white placeholder-gray-400"
                placeholder="Email address"
              />
              <div class="inset-0 border border-primary/0 rounded-lg group-hover:border-primary/20 transition-all duration-300"></div>
            </div>
          </div>

          <!-- Password Input -->
          <div class="group">
            <div class="relative">
              <input
                type="password"
                v-model="password"
                required
                class="w-full px-4 py-3 bg-gray-800/50 border border-gray-700 rounded-lg focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all duration-300 text-white placeholder-gray-400"
                placeholder="Password"
              />
              <div class="inset-0 border border-primary/0 rounded-lg group-hover:border-primary/20 transition-all duration-300"></div>
            </div>
          </div>

          <!-- Submit Button -->
          <button
            type="submit"
            :disabled="auth.loading"
            class="relative w-full group overflow-hidden"
          >
            <div class="-inset-0.5 bg-gradient-to-r from-primary to-secondary rounded-lg blur opacity-60 group-hover:opacity-100 transition duration-300 animate-pulse"></div>
            <div class="relative px-6 py-3 bg-gray-900 rounded-lg leading-none flex items-center justify-center">
              <span class="text-white font-medium transition duration-300">
                {{ auth.loading ? 'Signing in...' : 'Sign in' }}
              </span>
            </div>
          </button>

          <!-- Register Link -->
          <div class="text-center mt-6">
            <router-link
              to="/auth/register"
              class="text-gray-300 hover:text-white transition-colors duration-300"
            >
              Don't have an account? 
              <span class="font-medium text-primary hover:text-primary/80">Sign up</span>
            </router-link>
          </div>
        </form>

        <!-- Error Message -->
        <div v-if="error" class="mt-4 text-red-400 text-center text-sm">
          {{ error }}
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

@keyframes gradient {
  0% { background-position: 0% 50%; }
  50% { background-position: 100% 50%; }
  100% { background-position: 0% 50%; }
}

.animate-gradient {
  background-size: 200% auto;
  animation: gradient 4s linear infinite;
}
</style>