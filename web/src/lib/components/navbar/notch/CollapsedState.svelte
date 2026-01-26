<script lang="ts">
	import avatar from '$lib/assets/avatar.png';
    import { onMount } from 'svelte';
    import { fly } from 'svelte/transition';
    import SendioLogo from '$lib/components/icons/SendioLogo.svelte';

    let { name } = $props<{ name: string }>();
    
    let activeIndex = $state(0);
    let views = $derived([
        { type: 'name', value: name, color: 'text-white/95 font-black' },
        { type: 'growth', value: '+12.4%', color: 'text-emerald-400 font-semibold' },
        { type: 'quota', value: '5,863 / 8,000', color: 'text-white/95 font-bold' },
        { type: 'brand', isBrand: true }
    ]);

    onMount(() => {
        const interval = setInterval(() => {
            activeIndex = (activeIndex + 1) % views.length;
        }, 1500);
        return () => clearInterval(interval);
    });
</script>

<div class="grid grid-cols-[40px_1fr_40px] items-center h-full px-5 relative z-10 w-full">
    <!-- avatar (Login button) -->
    <div class="flex items-center justify-start h-full">
        <a 
            href="/login"
            onclick={(e) => e.stopPropagation()}
            class="group relative flex h-9 w-9 items-center justify-center rounded-full p-0.5 transition-all hover:scale-110 active:scale-95 bg-linear-to-tr from-white/10 to-white/5 border border-white/10 overflow-visible"
        >
            <div class="w-full h-full rounded-full bg-black flex items-center justify-center overflow-hidden shadow-2xl relative">
                <img src={avatar} alt="Login" class="w-full h-full object-cover" />
            </div>
            <!-- seal check -->
            <div class="absolute -top-0.5 -right-1 z-20 flex items-center justify-center">
                <i class="bx bxs-badge-check text-[#FFB800] text-[16px] drop-shadow-md"></i>
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
                        <div class="flex h-5 w-5 items-center justify-center text-white">
                            <SendioLogo class="h-full w-full" />
                        </div>
                        <span class="text-white uppercase font-black tracking-tight text-[16px]">Sendio</span>
                    </div>
                {:else}
                    <span class="text-[14px] tracking-tight {views[activeIndex].color} truncate leading-none">
                        {views[activeIndex].value}
                    </span>
                {/if}
            </div>
        {/key}
    </div>

    <!-- noti bell -->
    <div class="flex items-center justify-end">
        <div class="flex items-center text-orange-300">
            <i class="bx bx-bell text-[21px] leading-none"></i>
        </div>
    </div>
</div>
