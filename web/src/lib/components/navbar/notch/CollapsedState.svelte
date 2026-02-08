<script lang="ts">
    import { onMount } from 'svelte';
    import { fly } from 'svelte/transition';
    import SendioLogo from '$lib/components/icons/SendioLogo.svelte';
    import { user, getGrowthDisplay } from '$lib/stores/user.svelte';

    let { onSearch } = $props<{ onSearch: () => void }>();
    
    let activeIndex = $state(0);
    let views = $derived([
        { type: 'name', value: user.name, color: 'text-white font-black' },
        { type: 'growth', value: getGrowthDisplay(), color: 'text-emerald-400 font-semibold' },
        { type: 'quota', value: `${user.usage.sent.toLocaleString()} / ${user.usage.quota.toLocaleString()}`, color: 'text-white font-bold' },
        { type: 'brand', isBrand: true }
    ]);

    onMount(() => {
        const interval = setInterval(() => {
            activeIndex = (activeIndex + 1) % views.length;
        }, 1500);
        return () => clearInterval(interval);
    });
</script>

<div class="grid grid-cols-[84px_1fr_84px] items-center h-full px-5 relative z-10 w-full transition-all duration-400 ease-[cubic-bezier(0.3,1.5,0.4,1)]">
    <!-- left -->
    <div class="flex items-center justify-start gap-3 h-full">
        <!-- home -->
        <a 
            href="/"
            onclick={(e) => e.stopPropagation()}
            class="group relative flex h-9 w-9 items-center justify-center rounded-full transition-all duration-400 ease-[cubic-bezier(0.3,1.5,0.4,1)] hover:bg-white/10 active:scale-90 text-white"
            title="Return Home"
        >
            <SendioLogo class="h-5 w-5 transition-transform duration-400 group-hover:scale-110" />
        </a>

        <!-- avatar (Login button) -->
        <a 
            href="/login"
            draggable="false"
            onclick={(e) => e.stopPropagation()}
            class="group relative flex h-8 w-8 items-center justify-center rounded-full transition-all duration-400 ease-[cubic-bezier(0.3,1.5,0.4,1)] hover:bg-white/10 active:scale-90 bg-white/10 border border-white/20"
        >
            <div class="w-full h-full rounded-full flex items-center justify-center overflow-hidden transition-transform duration-400 group-hover:scale-110 relative">
                <!-- user icon fallback (always behind) -->
                <i class="bx bxs-user text-[18px] text-white/40 absolute z-0"></i>
                
                <img 
                    draggable="false" 
                    src={user.avatarUrl} 
                    alt="" 
                    class="w-full h-full object-cover relative z-10"
                />
            </div>
             <!-- seal check -->
            <div class="absolute -top-0.5 -right-1 z-20 flex items-center justify-center transition-transform duration-400 group-hover:scale-110 group-hover:translate-x-0.5 group-hover:-translate-y-0.5">
                <i draggable="false" class="bx bxs-badge-check text-[#FFB800] text-[16px] drop-shadow-md"></i>
            </div>
        </a>
    </div>

    <!-- rotating info -->
    <div class="flex justify-center items-center overflow-hidden h-full relative transition-all duration-500">
        {#key activeIndex}
            <div 
                in:fly={{ y: 20, duration: 500, easing: (t) => t * (2 - t) }}
                out:fly={{ y: -20, duration: 400 }}
                class="absolute inset-0 flex items-center justify-center"
            >
                {#if views[activeIndex].isBrand}
                    <div class="flex items-center gap-2">
                        <span class="text-white uppercase font-black tracking-widest text-[14px] transition-all duration-400">Sendio</span>
                    </div>
                {:else}
                    <span class="text-[14px] tracking-tight {views[activeIndex].color} truncate transition-all duration-400">
                        {views[activeIndex].value}
                    </span>
                {/if}
            </div>
        {/key}
    </div>

    <!-- right -->
    <div class="flex items-center justify-end gap-3 h-full">
        <!-- noti bell -->
        <div class="flex items-center justify-center w-9 h-9 text-orange-300 relative group cursor-pointer transition-all duration-400 ease-[cubic-bezier(0.3,1.5,0.4,1)] hover:bg-white/10 rounded-full" title="Notifications">
            <i class="bx bx-bell text-[22px] leading-none transition-all duration-400 group-hover:rotate-15 group-hover:scale-110"></i>
            <!-- indicator -->
            <div class="absolute top-1.5 right-1.5 w-2 h-2 rounded-full bg-red-500 border-2 border-black animate-pulse shadow-sm"></div>
        </div>

        <!-- search -->
        <div 
            role="button"
            tabindex="0"
            onclick={(e) => { e.stopPropagation(); onSearch(); }}
            onkeydown={(e) => {
                if (e.key === 'Enter') {
                    e.stopPropagation();
                    onSearch();
                }
            }}
            class="flex items-center justify-center w-9 h-9 text-white transition-all duration-400 ease-[cubic-bezier(0.3,1.5,0.4,1)] cursor-pointer rounded-full hover:bg-white/10 active:scale-90" 
            title="Search"
        >
            <i class='bx bx-search text-[21px] transition-transform duration-400 group-hover:scale-110'></i>
        </div>
    </div>
</div>
