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

<div class="grid grid-cols-[84px_1fr_84px] items-center h-full px-5 relative z-10 w-full">
    <!-- left -->
    <div class="flex items-center justify-start gap-3 h-full">
        <!-- home -->
        <a 
            href="/"
            onclick={(e) => e.stopPropagation()}
            class="group relative flex h-8 w-8 items-center justify-center rounded-full transition-all hover:scale-110 active:scale-95 text-white"
            title="Return Home"
        >
            <SendioLogo class="h-5 w-5" />
        </a>

        <!-- avatar (Login button) -->
        <a 
            href="/login"
            draggable="false"
            onclick={(e) => e.stopPropagation()}
            class="group relative flex h-7 w-7 items-center justify-center rounded-full p-0.5 transition-all hover:scale-110 active:scale-95 bg-linear-to-tr from-white/20 to-white/5 border border-white/20"
        >
            <div class="w-full h-full rounded-full bg-black flex items-center justify-center overflow-hidden">
                <img draggable="false" src={user.avatarUrl} alt="Login" class="w-full h-full object-cover" />
            </div>
             <!-- seal check -->
            <div class="absolute -top-0.5 -right-1 z-20 flex items-center justify-center">
                <i draggable="false" class="bx bxs-badge-check text-[#FFB800] text-[16px] drop-shadow-md"></i>
            </div>
        </a>
    </div>

    <!-- rotating info -->
    <div class="flex justify-center items-center overflow-hidden h-full relative">
        {#key activeIndex}
            <div 
                in:fly={{ y: 20, duration: 400, delay: 100 }}
                out:fly={{ y: -20, duration: 400 }}
                class="absolute inset-0 flex items-center justify-center"
            >
                {#if views[activeIndex].isBrand}
                    <div class="flex items-center gap-1.5">
                        <!-- <SendioLogo class="h-4 w-4 text-white" /> -->
                        <span class="text-white uppercase font-black tracking-tight text-[15px]">Sendio</span>
                    </div>
                {:else}
                    <span class="text-[14px] tracking-tight {views[activeIndex].color} truncate">
                        {views[activeIndex].value}
                    </span>
                {/if}
            </div>
        {/key}
    </div>

    <!-- right -->
    <div class="flex items-center justify-end gap-3 h-full">
        <!-- noti bell -->
        <div class="flex items-center text-orange-300 relative group cursor-pointer" title="Notifications">
            <i class="bx bx-bell text-[21px] leading-none transition-transform group-hover:rotate-12"></i>
            <!-- indicator -->
            <div class="absolute -top-1 -right-1 w-1.5 h-1.5 rounded-full bg-red-500 border border-black animate-pulse"></div>
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
            class="flex items-center text-white transition-all cursor-pointer p-1 rounded-full hover:scale-110 active:scale-95" 
            title="Search"
        >
            <i class='bx bx-search text-[20px]'></i>
        </div>
    </div>
</div>
