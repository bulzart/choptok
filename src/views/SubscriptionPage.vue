<script setup lang="ts">
import { ref } from 'vue'
import { Check, Crown, Zap, Star } from 'lucide-vue-next'

const plans = [

  {
    name: 'Pro',
    price: '$14.99',
    period: 'per month',
    description: 'Best for content creators',
    icon: Crown,
    features: [
      'Unlimited video uploads',
      'AI-powered editing',
      'HD quality exports',
      'Priority support',
      'Custom watermarks',
      'Advanced analytics',
      'Batch processing'
    ],
    buttonText: 'Subscribe',
    isPopular: true,
    disabled: false
  }
]

const selectedPlan = ref('monthly')
</script>

<template>
  <div class="min-h-screen py-16 px-4">
    <div class="max-w-7xl mx-auto">
      <!-- Header -->
      <div class="text-center mb-16">
        <h1 class="text-4xl font-bold mb-4 bg-gradient-to-r from-primary to-secondary bg-clip-text text-transparent">
          Choose Your Plan
        </h1>
        <p class="text-xl text-gray-600 dark:text-gray-400 max-w-2xl mx-auto">
          Get access to powerful features that help you create viral content faster
        </p>
      </div>

      <!-- Pricing Toggle -->
      <div class="flex justify-center mb-12">
        <div class="bg-gray-100 dark:bg-gray-800 p-1 rounded-xl inline-flex">
          <button
            @click="selectedPlan = 'monthly'"
            class="px-6 py-2 rounded-lg text-sm font-medium transition-colors"
            :class="selectedPlan === 'monthly' ? 'bg-white dark:bg-gray-700 shadow-sm' : 'text-gray-600 dark:text-gray-400'"
          >
            Monthly
          </button>
          
        </div>
      </div>

      <!-- Pricing Cards -->
      <div class="justify-center gap-8 max-w-7xl mx-auto ">
        <div
          v-for="plan in plans"
          :key="plan.name"
          class="relative rounded-2xl backdrop-blur-sm border transition-all duration-200 hover:shadow-xl"
          :class="[
            plan.isPopular 
              ? 'border-primary shadow-lg shadow-primary/20' 
              : 'border-gray-200 dark:border-gray-700'
          ]"
        >
          <!-- Popular Badge -->
          <div
            v-if="plan.isPopular"
            class="absolute -top-4 left-1/2 -translate-x-1/2 bg-gradient-to-r from-primary to-secondary text-white px-4 py-1 rounded-full text-sm font-medium"
          >
            Most Popular
          </div>

          <div class="p-8">
            <!-- Plan Header -->
            <div class="flex items-center gap-4 mb-6">
              <div 
                class="w-12 h-12 rounded-xl flex items-center justify-center"
                :class="plan.isPopular ? 'bg-primary/10' : 'bg-gray-100 dark:bg-gray-800'"
              >
                <component 
                  :is="plan.icon" 
                  class="w-6 h-6"
                  :class="plan.isPopular ? 'text-primary' : ''"
                />
              </div>
              <div>
                <h3 class="text-xl font-bold">{{ plan.name }}</h3>
                <p class="text-sm text-gray-600 dark:text-gray-400">{{ plan.description }}</p>
              </div>
            </div>

            <!-- Price -->
            <div class="mb-6">
              <div class="flex items-baseline">
                <span class="text-4xl font-bold">
                  {{ selectedPlan === 'annual' ? `${parseInt(plan.price) * 0.8}` : plan.price }}
                </span>
                <span class="text-gray-600 dark:text-gray-400 ml-2">
                  {{ plan.period }}
                </span>
              </div>
            </div>

            <!-- Features -->
            <ul class="space-y-4 mb-8">
              <li
                v-for="feature in plan.features"
                :key="feature"
                class="flex items-start gap-3"
              >
                <Check class="w-5 h-5 text-green-500 shrink-0 mt-0.5" />
                <span class="text-gray-600 dark:text-gray-400">{{ feature }}</span>
              </li>
            </ul>

            <!-- Action Button -->
            <button
              class="w-full py-3 px-6 rounded-xl font-medium transition-all duration-200"
              :class="[
                plan.isPopular 
                  ? 'bg-gradient-to-r from-primary to-secondary text-white hover:opacity-90' 
                  : 'bg-gray-100 dark:bg-gray-800 hover:bg-gray-200 dark:hover:bg-gray-700',
                plan.disabled ? 'opacity-50 cursor-not-allowed' : ''
              ]"
              :disabled="plan.disabled"
            >
              {{ plan.buttonText }}
            </button>
          </div>
        </div>
      </div>

      <!-- FAQ Section -->
      <div class="mt-24 max-w-3xl mx-auto">
        <h2 class="text-2xl font-bold text-center mb-8">Frequently Asked Questions</h2>
        <div class="space-y-6">
          <div class="bg-gray-50/80 dark:bg-gray-800/80 rounded-xl p-6">
            <h3 class="font-semibold mb-2">Can I change my plan later?</h3>
            <p class="text-gray-600 dark:text-gray-400">Yes, you can upgrade, downgrade, or cancel your plan at any time from your account settings.</p>
          </div>
          <div class="bg-gray-50/80 dark:bg-gray-800/80 rounded-xl p-6">
            <h3 class="font-semibold mb-2">What payment methods do you accept?</h3>
            <p class="text-gray-600 dark:text-gray-400">We accept all major credit cards, PayPal, and Apple Pay.</p>
          </div>
          <div class="bg-gray-50/80 dark:bg-gray-800/80 rounded-xl p-6">
            <h3 class="font-semibold mb-2">Is there a free trial?</h3>
            <p class="text-gray-600 dark:text-gray-400">Yes, you can try ChopTok Pro free for 14 days. No credit card required.</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>