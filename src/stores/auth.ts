import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'

export const useAuthStore = defineStore('auth', () => {
    const isAuthenticated = ref(true)
    const loading = ref(false)
    const router = useRouter()

    function signIn(email: string, password: string) {
        loading.value = true
        // Simulate API call

    }

    function signUp(email: string, password: string) {
        loading.value = true
        // Simulate API call
        setTimeout(() => {
            isAuthenticated.value = true
            loading.value = false
            router.push('/')
        }, 1000)
    }

    function signOut() {
        loading.value = true
        // Simulate API call
        setTimeout(() => {
            isAuthenticated.value = false
            loading.value = false
            router.push('/auth/login')
        }, 500)
    }

    return {
        isAuthenticated,
        loading,
        signIn,
        signUp,
        signOut
    }
})