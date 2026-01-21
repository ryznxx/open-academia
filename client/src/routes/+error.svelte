<script>
  import "../app.css";
</script>

<div class="min-h-screen flex items-center justify-center bg-gradient-to-br from-zinc-950 via-neutral-900 to-black text-neutral-200">
  <div class="max-w-md w-full px-6 text-center space-y-6">

    <!-- big code -->
    <h1 class="text-8xl font-black tracking-tight text-white">
      error
    </h1>

    <!-- vibe mahasiswa -->
    <p class="text-sm text-neutral-400 leading-relaxed">
      Lagi ngerjain tugas, kopi udah dingin, tab kebuka 23,
      terus halaman ini malah ngilang.  
      Yaudah. Tarik napas dulu.
    </p>

    <!-- divider -->
    <div class="flex items-center gap-3 justify-center opacity-40">
      <div class="h-px w-12 bg-neutral-600"></div>
      <span class="text-xs uppercase tracking-widest">system</span>
      <div class="h-px w-12 bg-neutral-600"></div>
    </div>

    <!-- action -->
    <div class="flex justify-center gap-4">
      <a
        href="/"
        class="px-5 py-2 text-sm rounded-md
               bg-white/5 border border-white/10
               hover:bg-white/10 hover:border-white/20
               transition">
        Balik
      </a>

      <button
        on:click={() => location.reload()}
        class="px-5 py-2 text-sm rounded-md
               bg-white/5 border border-white/10
               hover:bg-white/10 hover:border-white/20
               transition">
        Reload
      </button>
    </div>

    <!-- footer quote -->
    <p class="pt-8 text-xs text-neutral-500 italic">
      “Error ini bukan kegagalan.  
      Ini cuma sistem nyuruh lu istirahat bentar.”
    </p>

  </div>
</div>
