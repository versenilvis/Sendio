<script lang="ts">
	import CheckCopy from '../../icons/CheckCopy.svelte';
	import { user, getDisplayEmail, getUsagePercent, getGrowthDisplay } from '$lib/stores/user.svelte';

	let nameCopied = $state(false);
	let emailCopied = $state(false);

	function handleCopyName() {
		navigator.clipboard.writeText(user.name);
		nameCopied = true;
		setTimeout(() => (nameCopied = false), 1000);
	}

	function handleCopyEmail() {
		navigator.clipboard.writeText(user.email);
		emailCopied = true;
		setTimeout(() => (emailCopied = false), 1000);
	}
</script>

<div class="bg-white/2 border border-white/5 rounded-[40px] p-6 mb-6 flex gap-10 shadow-xl shadow-black/20">
    <div class="w-[210px] aspect-square rounded-[32px] overflow-hidden relative border border-white/5 shadow-2xl shrink-0 group select-none">
        <!-- background avatar -->
        <img draggable="false" src={user.avatarUrl} alt="User" class="absolute inset-0 w-full h-full object-cover transition-transform duration-1000" />
        
        <!-- black gradient overlay from bottom -->
        <div class="absolute inset-0 bg-linear-to-t from-black via-black/40 to-transparent pointer-events-none z-10"></div>
        
        <!-- seal check in corner -->
        <div class="absolute top-4 right-4 z-20">
            
        </div>

        <!-- text inside at bottom -->
        <div class="absolute inset-x-0 bottom-0 p-6 z-20 flex flex-col gap-1">
            <button 
                onclick={handleCopyName}
                class="group transition-all active:scale-95 text-left cursor-pointer"
            >
                <h3 class="text-lg font-bold leading-[1.3] tracking-tight drop-shadow-lg text-white">
                    {user.name}{#if user.isPremium}<i title="premium" class="bx bxs-badge-check text-[#FFB800] text-[18px] inline align-middle ml-0.5 drop-shadow-[0_4px_10px_rgba(0,0,0,0.5)]"></i>{/if}{#if nameCopied}<span class="inline-flex align-middle ml-1"><CheckCopy class="h-4 w-4" /></span>{/if}
                </h3>
            </button>
            
            <button 
                onclick={handleCopyEmail}
                class="group relative text-[11px] font-medium tracking-wide lowercase flex items-center gap-2 transition-all active:scale-95 text-left cursor-pointer text-white/50"
            >
                <span>{getDisplayEmail()}</span>
                {#if emailCopied}
                    <div class="shrink-0 flex items-center justify-center text-white select-none">
                        <CheckCopy class="h-3.5 w-3.5" />
                    </div>
                {/if}
            </button>
        </div>
    </div>

    <!-- Status -->
    <div class="flex-1 flex flex-col justify-between pt-1">
        <!-- quota progress -->
        <div class="flex flex-col gap-3">
            <div class="flex justify-between items-center">
                <div class="flex items-baseline gap-2">
                    <span class="text-4xl font-extrabold text-white tracking-tighter">{user.usage.sent.toLocaleString()}</span>
                    <span class="text-lg font-bold text-white/30">/ {user.usage.quota.toLocaleString()}</span>
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
                        style="width: {getUsagePercent()}%"
                    ></div>
                </div>
            </div>
        </div>

        <!-- line chart -->
        <div class="flex flex-col gap-3 mt-4">
            <div class="flex justify-between items-center text-[10px] font-bold text-white uppercase px-1">
                <span>last 7 days usage</span>
                <span class="text-emerald-400 font-bold">{getGrowthDisplay()}</span>
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
