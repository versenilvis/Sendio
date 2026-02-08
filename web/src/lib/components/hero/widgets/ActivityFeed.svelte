<script lang="ts">
</script>

<div class="bg-black rounded-[32px] p-6 w-64 h-56 shadow-2xl flex flex-col justify-between group transition-all duration-700 select-none overflow-hidden relative">
    
    <!-- Vivid Header without icon -->
    <div class="w-full flex justify-between items-start relative z-10">
        <div class="flex flex-col">
            <span class="text-[10px] font-bold text-white/60 tracking-tight uppercase">Reached Audience</span>
            <div class="flex items-end gap-2 mt-1">
                <span class="text-3xl font-bold text-white tracking-tighter leading-none">1.24M</span>
                <span class="text-[11px] font-black text-emerald-400 bg-emerald-400/20 px-2 py-1 rounded-lg tracking-tighter mb-[2px] leading-none">+12.4%</span>
            </div>
        </div>
    </div>

    <!-- Colorful Spectrum Waveform -->
    <div class="relative h-20 w-full flex items-end justify-center py-2">
        <div class="absolute inset-0 bg-blue-500/5 blur-3xl rounded-full"></div>
        <svg viewBox="0 0 100 40" class="w-full h-full overflow-visible relative z-10">
            <defs>
                <linearGradient id="spectrum" x1="0%" y1="0%" x2="100%" y2="0%">
                    <stop offset="0%" style="stop-color:#3b82f6" />
                    <stop offset="25%" style="stop-color:#6366f1" />
                    <stop offset="50%" style="stop-color:#a855f7" />
                    <stop offset="75%" style="stop-color:#ec4899" />
                    <stop offset="100%" style="stop-color:#f43f5e" />
                </linearGradient>
            </defs>
            <path 
                d="M0,38 Q10,38 20,25 T40,20 T60,25 T80,10 T100,5" 
                fill="none" 
                stroke="url(#spectrum)" 
                stroke-width="4" 
                stroke-linecap="round"
                class="animate-draw"
                stroke-dasharray="200"
                stroke-dashoffset="200"
            />
        </svg>
        <!-- Glow Cursor -->
        <div class="absolute right-[2%] top-[15%] h-2.5 w-2.5 rounded-full bg-rose-500 shadow-[0_0_15px_#f43f5e]"></div>
    </div>

    <!-- Bottom Metadata with aesthetic improvements -->
    <div class="flex justify-between items-end relative z-10">
        <div class="flex -space-x-3">
            <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=Felix" class="h-8 w-8 rounded-full border-2 border-black bg-zinc-800 object-cover" alt="" />
            <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=Aneka" class="h-8 w-8 rounded-full border-2 border-black bg-zinc-800 object-cover" alt="" />
            <img src="https://api.dicebear.com/7.x/avataaars/svg?seed=Liam" class="h-8 w-8 rounded-full border-2 border-black bg-zinc-800 object-cover" alt="" />
        </div>
        <div class="flex flex-col items-start min-w-[100px] ml-4">
            <span class="text-[9px] font-bold text-white/60 tracking-tight uppercase mb-1">Top Region</span>
            <span class="text-[10px] font-bold text-white py-1 rounded-full uppercase tracking-tight truncate w-full">United States</span>
        </div>
    </div>
</div>

<style>
    @keyframes draw {
        to { stroke-dashoffset: 0; }
    }
    .animate-draw {
        animation: draw 3s cubic-bezier(0.4, 0, 0.2, 1) forwards;
    }
</style>
