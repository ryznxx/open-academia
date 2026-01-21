<script lang="ts">
  import { config } from "$lib/config/config";
  import { fly, fade } from "svelte/transition";

  const content = {
    title: "Single Sign On (SSO)",
    subtitle: config.signinWelcomingMessage || "Selamat datang kembali di dashboard Anda.",
    footerText: "Belum punya akun?",
    footerAction: "Hubungi Admin",
    buttonText: "Masuk ke Akun",
    ssoText: "Atau masuk dengan"
  };

  // State
  let isNIMMode = false;
  let identifier = "";
  let password = "";
  let isSubmitting = false;

  const handleSubmit = () => {
    isSubmitting = true;
    setTimeout(() => { isSubmitting = false; }, 2000);
  };

  // SSO Handler
  const handleSSO = (provider: string) => {
    console.log(`Logging in with ${provider}`);
  };
</script>

<div class="min-h-screen w-full bg-white flex text-gray-800 overflow-hidden">
  
  <div class="hidden md:flex w-1/2 bg-gray-50 border-r border-gray-100 p-20 flex-col justify-between relative overflow-hidden">
    <div class="absolute top-[-10%] left-[-10%] w-64 h-64 bg-gray-200/50 rounded-full blur-3xl"></div>
    <div class="absolute bottom-[5%] right-[-5%] w-96 h-96 bg-gray-200/30 rounded-full blur-3xl"></div>
    
    <div class="space-y-6 relative z-10">
      <div class="h-12 w-12 bg-gray-900 rounded-xl flex items-center justify-center text-white text-xl font-bold shadow-lg">
        {config.appName?.charAt(0) || 'A'}
      </div>
      <div class="space-y-4">
        <h2 class="text-4xl font-bold tracking-tight text-gray-900 leading-tight">
          {content.title}
        </h2>
        <p class="text-xl text-gray-400 max-w-md leading-relaxed font-medium">
          {content.subtitle}
        </p>
      </div>
    </div>

    <div class="space-y-4 relative z-10">
      <div class="flex gap-1.5">
        {#each Array(3) as _, i}
            <div class="h-1 {i === 0 ? 'w-12 bg-gray-900' : 'w-2 bg-gray-200'} rounded-full"></div>
        {/each}
      </div>
      <p class="text-xs text-gray-400 font-bold tracking-[0.2em] uppercase">
        {config.appName} Digital Ecosystem
      </p>
    </div>
  </div>

  <div class="flex-1 flex items-center justify-center p-8 md:p-16 bg-white relative">
    <div class="w-full max-w-sm relative z-10">
      
      <div class="space-y-3 mb-8">
        <button 
          on:click={() => handleSSO('Google')}
          class="w-full flex items-center justify-center gap-3 py-3 border border-gray-200 rounded-xl hover:bg-gray-50 transition-all font-medium text-sm shadow-sm active:scale-[0.98]"
        >
          <svg class="w-5 h-5" viewBox="0 0 24 24">
            <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"/><path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/><path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l3.66-2.84z" fill="#FBBC05"/><path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/>
          </svg>
          Masuk dengan Google
        </button>
      </div>

      <div class="relative flex py-5 items-center">
        <div class="flex-grow border-t border-gray-100"></div>
        <span class="flex-shrink mx-4 text-gray-400 text-xs font-bold uppercase tracking-widest">{content.ssoText}</span>
        <div class="flex-grow border-t border-gray-100"></div>
      </div>
      
      <div class="mb-8 p-1.5 bg-gray-100 rounded-2xl flex relative">
        <div 
          class="absolute top-1.5 bottom-1.5 w-[calc(50%-6px)] bg-white rounded-xl shadow-sm transition-all duration-300 ease-out"
          class:translate-x-full={isNIMMode}
          class:translate-x-0={!isNIMMode}
        ></div>
        
        <button 
          type="button"
          class="flex-1 py-2 text-xs font-bold z-10 transition-colors duration-300"
          class:text-gray-900={!isNIMMode}
          class:text-gray-400={isNIMMode}
          on:click={() => isNIMMode = false}
        >
          Email
        </button>
        <button 
          type="button"
          class="flex-1 py-2 text-xs font-bold z-10 transition-colors duration-300"
          class:text-gray-900={isNIMMode}
          class:text-gray-400={!isNIMMode}
          on:click={() => isNIMMode = true}
        >
          NIM
        </button>
      </div>

      <form on:submit|preventDefault={handleSubmit} class="space-y-5">
        <div class="space-y-2">
          <label for="identifier" class="text-[11px] font-bold text-gray-400 uppercase tracking-widest ml-1">
            {#if isNIMMode}
              <span in:fly={{ y: 5, duration: 200 }}>Nomor Induk Mahasiswa</span>
            {:else}
              <span in:fly={{ y: 5, duration: 200 }}>Email Address</span>
            {/if}
          </label>
          
          <div class="relative">
            {#if isNIMMode}
              <input
                key="nim" in:fade
                type="text" bind:value={identifier} placeholder="220101xxx"
                class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-gray-900 focus:bg-white transition-all outline-none"
                required
              />
            {:else}
              <input
                key="email" in:fade
                type="email" bind:value={identifier} placeholder="name@academia.edu"
                class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-gray-900 focus:bg-white transition-all outline-none"
                required
              />
            {/if}
          </div>
        </div>

        <div class="space-y-2">
          <div class="flex justify-between items-center ml-1">
            <label for="password" class="text-[11px] font-bold text-gray-400 uppercase tracking-widest">Password</label>
            <a href="/" class="text-xs font-semibold text-gray-400 hover:text-gray-900 transition-colors">Lupa?</a>
          </div>
          <input
            type="password" bind:value={password} placeholder="••••••••"
            class="w-full px-4 py-3 bg-gray-50 border border-gray-200 rounded-xl focus:ring-2 focus:ring-gray-900 focus:bg-white transition-all outline-none"
            required
          />
        </div>

        <button
          type="submit"
          disabled={isSubmitting}
          class="w-full py-4 bg-gray-900 hover:bg-black text-white rounded-xl font-bold transition-all flex items-center justify-center gap-3 shadow-xl shadow-gray-200 active:scale-[0.99] disabled:opacity-70 mt-2"
        >
          {#if isSubmitting}
            <div class="w-5 h-5 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
          {/if}
          {content.buttonText}
        </button>
      </form>

      <div class="mt-10 pt-8 border-t border-gray-50 text-center">
        <p class="text-xs font-medium text-gray-400">
          {content.footerText} 
          <a href="/" class="text-gray-900 font-bold hover:underline ml-1 uppercase tracking-tighter">
            {content.footerAction}
          </a>
        </p>
      </div>
    </div>
  </div>
</div>