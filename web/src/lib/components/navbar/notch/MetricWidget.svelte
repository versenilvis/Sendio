<script lang="ts">
	import avatar from '$lib/assets/avatar.png';
	let { name } = $props<{ name: string }>();
</script>

<div class="bg-white/2 border border-white/5 rounded-[40px] p-6 mb-6 flex gap-10 shadow-xl shadow-black/20">
    <!-- round square box -->
    <div class="w-[210px] aspect-square flex flex-col items-center justify-center bg-black/40 rounded-[32px] p-5 border border-white/5 shadow-inner shrink-0">
        <div class="relative mb-6">
            <div class="absolute inset-0 bg-blue-500/10 blur-3xl rounded-full"></div>
            <!-- pfp -->
            <div class="relative h-20 w-20">
                <div class="w-full h-full rounded-xl bg-white/5 border border-white/10 flex items-center justify-center overflow-hidden">
                    <img src={avatar} alt="User Avatar" class="w-full h-full object-cover" />
                </div>
                <!-- seal check -->
                <div class="absolute -top-1.5 -right-1.5 h-8 w-8 flex items-center justify-center z-10">
                    <i class="bx bxs-badge-check text-yellow-500 text-[28px]"></i>
                </div>
            </div>
        </div>
        <div class="text-center overflow-hidden w-full">
            <div class="text-xl font-bold text-white leading-tight truncate">{name}</div>
            <div class="text-xs font-medium text-white/50 truncate mt-2">admin@sendio.dev</div>
        </div>
    </div>

    <!-- Status -->
    <div class="flex-1 flex flex-col justify-between pt-1">
        <!-- quota progress -->
        <div class="flex flex-col gap-3">
            <div class="flex justify-between items-center">
                <div class="flex items-baseline gap-2">
                    <span class="text-4xl font-extrabold text-white tracking-tighter">5,863</span>
                    <span class="text-lg font-bold text-white/30">/ 8,000</span>
                </div>
                <div class="flex flex-col items-end">
                    <span class="text-[11px] font-bold text-white/80 uppercase mb-1 flex items-center leading-none">
                        Details
                        <i class='bx bx-chevron-right text-[14px]'></i>
                    </span>
                    <div class="h-0.5 w-15 bg-white/40 rounded-full"></div>
                </div>
            </div>
            
            <div class="relative w-full h-8 flex items-center">
                <!-- background track -->
                <div class="absolute inset-0 h-3 top-2.5 bg-white/5 rounded-full border border-white/5 overflow-hidden">
                    <!-- progress fill -->
                    <div 
                        class="h-full bg-white rounded-full transition-all duration-1000 ease-out shadow-[0_0_20px_white]" 
                        style="width: 73.2%"
                    ></div>
                </div>
            </div>
        </div>

        <!-- line chart -->
        <div class="flex flex-col gap-3 mt-4">
            <div class="flex justify-between items-center text-[10px] font-bold text-white uppercase px-1">
                <span>last 7 days usage</span>
                <span class="text-emerald-400 font-bold">+12.4%</span>
            </div>
            
            <div class="h-20 w-full relative">
                <svg viewBox="0 0 100 40" class="w-full h-full overflow-visible">
                    <defs>
                        <linearGradient id="lineGradient" x1="0" y1="0" x2="0" y2="1">
                            <stop offset="0%" stop-color="#10b981" stop-opacity="0.3" />
                            <stop offset="100%" stop-color="#10b981" stop-opacity="0" />
                        </linearGradient>
                    </defs>
                    <!-- area fill -->
                    <path 
                        d="M 0,22 L 16.6,9 L 33.3,19 L 50,4 L 66.6,14 L 83.3,6 L 100,11 L 100,40 L 0,40 Z" 
                        fill="url(#lineGradient)"
                        class="transition-all duration-1000"
                    />
                    <!-- line -->
                    <path 
                        d="M 0,22 L 16.6,9 L 33.3,19 L 50,4 L 66.6,14 L 83.3,6 L 100,11" 
                        fill="none" 
                        stroke="#10b981" 
                        stroke-width="2" 
                        stroke-linecap="round" 
                        stroke-linejoin="round"
                        class="transition-all duration-1000"
                    />
                    <!-- data points -->
                    {#each [22, 9, 19, 4, 14, 6, 11] as y, i}
                        <circle cx={i * 16.6} cy={y} r="2.5" fill="white" stroke="#10b981" stroke-width="1.5" />
                    {/each}
                </svg>
            </div>

            <!-- days label -->
            <div class="flex justify-between px-2 text-[10px] font-bold text-white/50 uppercase">
                <span>Mon</span><span>Tue</span><span>Wed</span><span>Thu</span><span>Fri</span><span>Sat</span><span>Sun</span>
            </div>
        </div>
    </div>
</div>
