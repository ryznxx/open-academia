<script lang="ts">
  import { config } from "$lib/config/config";

  const content = {
    title: "Sign in",
    subtitle: config.signinWelcomingMessage || "Selamat datang kembali di dashboard Anda.",
    footerText: "Belum punya akun?",
    footerAction: "Hubungi Admin",
    buttonText: "Masuk ke Akun"
  };

  const placeholders = {
    email: "email@contoh.com",
    password: "••••••••"
  };

  let email = "";
  let password = "";
  let isSubmitting = false;

  const handleSubmit = () => {
    isSubmitting = true;
    setTimeout(() => { isSubmitting = false; }, 2000);
  };
</script>

<div class="min-h-screen w-full bg-white flex text-gray-800">
  
  <div class="hidden md:flex w-1/2 bg-gray-50 border-r border-gray-100 p-20 flex-col justify-between">
    <div class="space-y-6">
      <div class="h-12 w-12 bg-gray-900 rounded-xl flex items-center justify-center text-white text-xl font-bold shadow-lg">
        {config.appName?.charAt(0) || 'A'}
      </div>
      <div class="space-y-4">
        <h2 class="text-4xl font-bold tracking-tight text-gray-900 leading-tight">
          {content.title}
        </h2>
        <p class="text-xl text-gray-500 max-w-md leading-relaxed">
          {content.subtitle}
        </p>
      </div>
    </div>

    <div class="space-y-4">
      <div class="flex gap-3">
        <div class="h-1 w-12 bg-gray-900 rounded-full"></div>
        <div class="h-1 w-4 bg-gray-200 rounded-full"></div>
        <div class="h-1 w-4 bg-gray-200 rounded-full"></div>
      </div>
      <p class="text-sm text-gray-400 font-medium tracking-widest uppercase">
        Internal Data Management
      </p>
    </div>
  </div>

  <div class="flex-1 flex items-center justify-center p-8 md:p-16 bg-white">
    <div class="w-full max-w-sm">
      <div class="md:hidden mb-8 h-10 w-10 bg-gray-900 rounded-lg flex items-center justify-center text-white font-bold">
        {config.appName?.charAt(0) || 'A'}
      </div>

      <form on:submit|preventDefault={handleSubmit} class="space-y-6">
        <div class="space-y-2">
          <label for="email" class="text-sm font-semibold text-gray-700 tracking-tight">Email Address</label>
          <input
            type="email"
            bind:value={email}
            placeholder={placeholders.email}
            class="w-full px-4 py-3.5 mt-2 bg-gray-50 border border-gray-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-gray-900 focus:bg-white transition-all placeholder:text-gray-400"
            required
          />
        </div>

        <div class="space-y-2">
          <div class="flex justify-between items-center">
            <label for="password" class="text-sm font-semibold text-gray-700 tracking-tight">Password</label>
            <a href="/" class="text-xs font-medium text-gray-400 hover:text-gray-900 transition-colors">Lupa sandi?</a>
          </div>
          <input
            type="password"
            bind:value={password}
            placeholder={placeholders.password}
            class="w-full px-4 py-3.5 bg-gray-50 border border-gray-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-gray-900 focus:bg-white transition-all placeholder:text-gray-400"
            required
          />
        </div>

        <button
          type="submit"
          disabled={isSubmitting}
          class="w-full py-4 px-4 bg-gray-900 hover:bg-black text-white rounded-xl font-semibold transition-all flex items-center justify-center gap-3 shadow-xl shadow-gray-200 active:scale-[0.99] disabled:opacity-70"
        >
          {#if isSubmitting}
            <div class="w-5 h-5 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
          {/if}
          {content.buttonText}
        </button>
      </form>

      <div class="mt-12 pt-8 border-t border-gray-100 text-center">
        <p class="text-sm text-gray-500">
          {content.footerText} 
          <a href="/" class="text-gray-900 font-bold hover:underline underline-offset-4 ml-1">
            {content.footerAction}
          </a>
        </p>
      </div>
    </div>
  </div>

</div>
