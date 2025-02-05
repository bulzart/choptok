<script setup lang="ts">
const videos = [
  {
    id: 1,
    title: "Morning Routine for Success",
    views: "450K",
    likes: "52K",
    thumbnail: "https://picsum.photos/400/600"
  },
  {
    id: 2,
    title: "5 Productivity Hacks",
    views: "280K",
    likes: "35K",
    thumbnail: "https://picsum.photos/400/601"
  },
  {
    id: 3,
    title: "Business Tips for 2024",
    views: "620K",
    likes: "78K",
    thumbnail: "https://picsum.photos/400/602"
  }
]
</script>

<template>
  <div class="p-8">
    <h1 class="text-2xl font-bold mb-6">My Videos</h1>
    <div class="grid md:grid-cols-3 gap-6">
      <div 
        v-for="video in videos" 
        :key="video.id"
        class="bg-white/80 dark:bg-gray-800/80 backdrop-blur-sm rounded-xl overflow-hidden shadow-lg"
      >
        <img :src="video.thumbnail" :alt="video.title" class="w-full aspect-video object-cover" />
        <div class="p-4">
          <h3 class="font-semibold mb-2">{{ video.title }}</h3>
          <div class="flex justify-between text-sm text-gray-600 dark:text-gray-400">
            <span>{{ video.views }} views</span>
            <span>{{ video.likes }} likes</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>