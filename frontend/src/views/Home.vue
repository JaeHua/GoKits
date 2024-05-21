<template>

  <main class="container mx-auto px-4 py-12 md:px-6 lg:py-16">
    <div class="grid gap-8 md:grid-cols-2 lg:grid-cols-3">
      <div @click="handleClick" class="flex flex-col items-center justify-center space-y-4 rounded-lg bg-[#f8f5f2] p-6 shadow-lg transition-transform duration-300 hover:-translate-y-2 hover:shadow-xl">
        <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
            class="h-10 w-10 text-[#d4b8a3]"
        >
          <path d="M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5.08-1.25-.27-2.48-1-3.5.28-1.15.28-2.35 0-3.5 0 0-1 0-3 1.5-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5-.39.49-.68 1.05-.85 1.65-.17.6-.22 1.23-.15 1.85v4"></path>
          <path d="M9 18c-4.51 2-5-2-7-2"></path>
        </svg>
        <h3 class="text-xl font-semibold">Github</h3>
        <p class="text-sm text-gray-500">赞一个star吧⭐</p>
      </div>
      <div @click="$router.replace('todolist')" class="flex flex-col items-center justify-center space-y-4 rounded-lg bg-[#f8f5f2] p-6 shadow-lg transition-transform duration-300 hover:-translate-y-2 hover:shadow-xl">
        <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
            class="h-10 w-10 text-[#d4b8a3]"
        >
          <rect width="8" height="4" x="8" y="2" rx="1" ry="1"></rect>
          <path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"></path>
        </svg>
        <h3 class="text-xl font-semibold">To-Do List</h3>
        <p class="text-sm text-gray-500">Stay organized and on top of your tasks with our simple to-do list.</p>
      </div>
      <div class="flex flex-col items-center justify-center space-y-4 rounded-lg bg-[#f8f5f2] p-6 shadow-lg transition-transform duration-300 hover:-translate-y-2 hover:shadow-xl">
        <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
            class="h-10 w-10 text-[#d4b8a3]"
        >
          <circle cx="12" cy="12" r="10"></circle>
          <path d="M8 14s1.5 2 4 2 4-2 4-2"></path>
          <line x1="9" x2="9.01" y1="9" y2="9"></line>
          <line x1="15" x2="15.01" y1="9" y2="9"></line>
        </svg>
        <h3 class="text-xl font-semibold">Daily Mood Tracker</h3>
        <p class="text-sm text-gray-500">Monitor your daily mood and emotions with our intuitive tracker.</p>
      </div>
      <div class="flex flex-col items-center justify-center space-y-4 rounded-lg bg-[#f8f5f2] p-6 shadow-lg transition-transform duration-300 hover:-translate-y-2 hover:shadow-xl">
        <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
            class="h-10 w-10 text-[#d4b8a3]"
        >
          <path d="M8 2v4"></path>
          <path d="M16 2v4"></path>
          <rect width="18" height="18" x="3" y="4" rx="2"></rect>
          <path d="M3 10h18"></path>
        </svg>
        <h3 class="text-xl font-semibold">敬请期待</h3>
        <p class="text-sm text-gray-500">please looking forward to.</p>
      </div>

    </div>
    <div class="mt-12 ml-12 text-center">
      <h1 id="3x0kzxobllt" class="text-3xl font-bold text-[#d4b8a3] md:text-4xl lg:text-5xl animate-bounce">Simplify Your Life</h1>
      <p class="mt-4 text-gray-500 md:text-lg">
        Discover our suite of tools to help you stay organized and on track.
      </p>

    </div>
  </main>
</template>

<script setup>
const handleClick = () =>{
  window.location.href = 'https://github.com/JaeHua';
}

</script>
<style scoped>
</style>
