<script setup lang="ts">
import { useDark, useToggle } from '@vueuse/core'
import { Sun, Moon } from 'lucide-vue-next'

const isDark = useDark()
const toggleDark = useToggle(isDark)
</script>

<template>
  <button 
    @click="toggleDark()"
    class="fixed top-4 right-4 p-2 rounded-lg bg-gray-200 dark:bg-gray-800"
    aria-label="Toggle theme"
  >
    <Sun v-if="isDark" class="w-6 h-6" />
    <Moon v-else class="w-6 h-6" />
  </button>
</template>