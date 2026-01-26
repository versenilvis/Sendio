<script lang="ts">
	import avatar from '$lib/assets/avatar.png';
    import { onMount } from 'svelte';
    import { fly } from 'svelte/transition';

    let { name } = $props<{ name: string }>();
    
    let activeIndex = $state(0);
    let views = $derived([
        { type: 'name', value: name, color: 'text-white/95 font-black' },
        { type: 'growth', value: '+12.4%', color: 'text-emerald-400 font-semibold' },
        { type: 'quota', value: '5,863 / 8,000', color: 'text-white/95 font-bold' }
    ]);

    onMount(() => {
        const interval = setInterval(() => {
            activeIndex = (activeIndex + 1) % views.length;
        }, 2000);
        return () => clearInterval(interval);
    });
</script>

<div class="grid grid-cols-[40px_1fr_40px] items-center h-full px-5 relative z-10 w-full">
    <!-- avatar -->
    <div class="flex items-center justify-start">
        <div class="flex h-8 w-8 items-center justify-center rounded-full p-0.5 shadow-lg">
            <div class="w-full h-full rounded-full bg-black flex items-center justify-center overflow-hidden">
                <img src={avatar} alt="User Avatar" class="w-full h-full object-cover" />
            </div>
        </div>
    </div>

    <!-- rotating info -->
    <div class="flex justify-center items-center overflow-hidden h-full relative">
        {#key activeIndex}
            <div 
                in:fly={{ y: 20, duration: 400, delay: 100 }}
                out:fly={{ y: -20, duration: 400 }}
                class="absolute inset-0 flex items-center justify-center"
            >
                <span class="text-[14px] tracking-tight {views[activeIndex].color} truncate leading-none">
                    {views[activeIndex].value}
                </span>
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
