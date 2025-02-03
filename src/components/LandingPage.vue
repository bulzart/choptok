<script setup lang="ts">
import { ref } from 'vue';
import { Upload, Youtube, Zap } from 'lucide-vue-next';

const features = [
  {
    icon: Upload,
    title: 'Upload Your Video',
    description: 'Simply upload your video, and our AI will analyze and chop it into engaging TikTok/Reel clips.'
  },
  {
    icon: Youtube,
    title: 'Paste YouTube Link',
    description: 'Paste a YouTube URL, and we’ll extract the video, analyze it, and generate shareable clips for you.'
  },
  {
    icon: Zap,
    title: 'AI-Powered Editing',
    description: 'Our AI adds captions, transitions, and effects to make your videos viral-ready.'
  }
];

const testimonials = [
  {
    quote: 'ChopTok saved me hours of editing! The AI-generated clips are so professional and engaging.',
    author: 'Sarah T., Social Media Manager'
  },
  {
    quote: 'I love how easy it is to turn my long YouTube videos into TikTok-ready content. Highly recommend!',
    author: 'James L., Content Creator'
  }
];

// New Example: Joe Rogan Podcast transformation demo
const examples = [
  {
    title: "Joe Rogan Podcast Transformation",
    // Original video thumbnail (e.g. a 30-min podcast)
    original: 'https://placehold.co/600x400/EEE/31343C?text=Joe+Rogan+Podcast+30min',
    // Three generated TikTok/Reel clips with captions
    generated: [
      {
        thumbnail: 'https://placehold.co/300x200/31343C/EEE?text=Key+Moment+1',
        caption: 'Key Moment 1 – Insightful Discussion'
      },
      {
        thumbnail: 'https://placehold.co/300x200/31343C/EEE?text=Key+Moment+2',
        caption: 'Key Moment 2 – Controversial Topic'
      },
      {
        thumbnail: 'https://placehold.co/300x200/31343C/EEE?text=Key+Moment+3',
        caption: 'Key Moment 3 – Funny Highlight'
      }
    ],
    description: 'Watch how a 30-minute Joe Rogan podcast is transformed into 3 engaging, 45-second reels highlighting the best moments.'
  }
];

const isDarkMode = ref(false);
const toggleDarkMode = () => {
  isDarkMode.value = !isDarkMode.value;
  document.documentElement.classList.toggle('dark', isDarkMode.value);
};
</script>

<template>
  <div class="min-h-screen bg-white dark:bg-gray-900 text-gray-900 dark:text-white transition-colors duration-200">
    <!-- Header -->
    <header class="relative bg-gradient-to-r from-primary to-secondary py-20 overflow-hidden">
      <div class="absolute inset-0 overflow-hidden pointer-events-none">
        <!-- Animated Background Elements -->
        <div class="absolute w-32 h-32 bg-purple-600/20 rounded-lg blur-xl animate-pulse" :style="{ left: '10%', top: '20%' }"></div>
        <div class="absolute w-48 h-48 bg-purple-600/20 rounded-lg blur-xl animate-pulse" :style="{ left: '70%', top: '50%' }"></div>
      </div>
      <div class="container mx-auto px-6 text-center relative z-10">
        <h1 class="text-5xl font-bold text-white mb-4">
          Turn Your Videos into Viral TikTok & Reels with ChopTok
        </h1>
        <p class="text-lg text-white/80 mb-8">
          Upload your video or paste a YouTube link – let our AI analyze, chop, and transform your content into shareable clips.
        </p>
        <div class="flex justify-center gap-4">
          <button class="bg-white text-primary px-8 py-3 rounded-lg font-semibold hover:bg-gray-100 transition-colors">
            Get Started
          </button>
          <button class="bg-transparent border border-white text-white px-8 py-3 rounded-lg font-semibold hover:bg-white/10 transition-colors">
            How It Works
          </button>
        </div>
      </div>
    </header>

    <!-- Features Section -->
    <section class="container mx-auto px-6 py-16">
      <div class="text-center mb-12">
        <h2 class="text-3xl font-bold mb-4">How It Works</h2>
        <p class="text-gray-600 dark:text-gray-400">
          ChopTok makes it easy to create viral-ready content in just a few steps.
        </p>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
        <div v-for="(feature, index) in features" :key="index" class="bg-white dark:bg-gray-800 p-6 rounded-xl shadow-lg text-center">
          <component :is="feature.icon" class="w-12 h-12 text-primary mx-auto mb-4" />
          <h3 class="text-xl font-semibold mb-2">{{ feature.title }}</h3>
          <p class="text-gray-600 dark:text-gray-400">{{ feature.description }}</p>
        </div>
      </div>
    </section>

    <!-- Example Section: Joe Rogan Podcast Transformation -->
    <section class="bg-gray-50 dark:bg-gray-800 py-16">
      <div class="container mx-auto px-6">
        <div class="text-center mb-12">
          <h2 class="text-3xl font-bold mb-4">See It in Action</h2>
          <p class="text-gray-600 dark:text-gray-400">
            Experience how ChopTok transforms a long-form podcast into bite-sized, engaging clips.
          </p>
        </div>
        <div class="mb-12">
          <h3 class="text-xl font-bold text-center mb-4">{{ examples[0].title }}</h3>
          <div class="flex flex-col md:flex-row gap-8 justify-center">
            <!-- Original Podcast Thumbnail -->
            <div class="flex-1">
              <h4 class="text-lg font-semibold text-center mb-2">Original Podcast</h4>
              <img :src="examples[0].original" alt="Original Podcast" class="rounded-lg shadow-md hover:scale-105 transition-transform" />
            </div>
            <!-- Generated Key Moments -->
            <div class="flex-1">
              <h4 class="text-lg font-semibold text-center mb-2">Generated Key Moments</h4>
              <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
                <div v-for="(clip, i) in examples[0].generated" :key="i" class="relative group">
                  <img :src="clip.thumbnail" alt="Key Moment" class="rounded-lg shadow-md hover:scale-105 transition-transform" />
                  <div class="absolute inset-0 bg-black bg-opacity-50 flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity">
                    <span class="text-white text-sm text-center px-2">{{ clip.caption }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <p class="text-gray-600 dark:text-gray-300 mt-4 text-center">
            {{ examples[0].description }}
          </p>
        </div>
      </div>
    </section>

    <!-- Testimonials Section -->
    <section class="container mx-auto px-6 py-16">
      <div class="text-center mb-12">
        <h2 class="text-3xl font-bold mb-4">What Our Users Say</h2>
        <p class="text-gray-600 dark:text-gray-400">
          Hear from creators who are already using ChopTok to grow their audience.
        </p>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div v-for="(testimonial, index) in testimonials" :key="index" class="bg-white dark:bg-gray-700 p-6 rounded-xl shadow-lg">
          <p class="text-gray-600 dark:text-gray-300 italic mb-4">"{{ testimonial.quote }}"</p>
          <p class="text-gray-900 dark:text-white font-semibold">{{ testimonial.author }}</p>
        </div>
      </div>
    </section>

    <!-- Call-to-Action Section -->
    <section class="bg-gradient-to-r from-primary to-secondary py-20">
      <div class="container mx-auto px-6 text-center">
        <h2 class="text-3xl font-bold text-white mb-4">Ready to Go Viral?</h2>
        <p class="text-lg text-white/80 mb-8">
          Start creating TikTok and Reel-ready videos in minutes.
        </p>
        <button class="bg-white text-primary px-8 py-3 rounded-lg font-semibold hover:bg-gray-100 transition-colors">
          Get Started Now
        </button>
      </div>
    </section>

    <!-- Footer -->
    <footer class="bg-gray-900 text-white py-8">
      <div class="container mx-auto px-6 text-center">
        <p class="text-sm text-gray-400">&copy; 2023 ChopTok. All rights reserved.</p>
      </div>
    </footer>
  </div>
</template>

<style scoped>
@keyframes pulse {
  0%, 100% { opacity: 0.2; }
  50% { opacity: 0.5; }
}

.animate-pulse {
  animation: pulse 3s ease-in-out infinite;
}

.hover\:scale-105:hover {
  transform: scale(1.05);
}
</style>
