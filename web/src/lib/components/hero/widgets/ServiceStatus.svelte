<script lang="ts">
</script>

<div class="bg-black rounded-3xl p-5 w-48 h-40 shadow-2xl flex flex-col justify-between group transition-all duration-700 select-none overflow-hidden relative">

    <div class="relative z-10 flex justify-between items-start">
        <div class="flex flex-col">
            <span class="text-[10px] font-bold text-white/60 tracking-tight uppercase">Response Rate</span>
            <span class="text-3xl font-bold text-white tracking-tighter mt-1">42.4<span class="text-lg text-white ml-1">%</span></span>
        </div>
        <div class="h-8 w-8 rounded-full bg-purple-500/20 flex items-center justify-center text-purple-400">
            <i class='bx bxs-bolt text-lg'></i>
        </div>
    </div>

    <!-- Colorful Magenta/Purple Bar Chart like Music/Heart Rate -->
    <div class="relative z-10 flex items-end gap-1.5 h-12 bg-white/5 rounded-2xl px-3 py-2">
        {#each [20, 35, 60, 45, 80, 50, 95, 70] as h, i}
            <div 
                class="flex-1 rounded-full transition-all duration-1000"
                style="height: {h}%; background: linear-gradient(to top, #7c3aed, #d946ef); opacity: {0.4 + (i * 0.08)}"
            ></div>
        {/each}
    </div>

    <div class="relative z-10 flex justify-between items-center text-[9px] font-bold text-white/60 tracking-tight uppercase">
        <span class="text-purple-400">High Engagement</span>
        <span class="bg-purple-500/20 text-purple-300 px-1.5 py-0.5 rounded uppercase">New Record</span>
    </div>
</div>
