<script lang="ts">
    import { userStore } from '$lib/stores/user.svelte';

    let { onSignOut } = $props<{
        onSignOut: () => void;
    }>();

    const menuItems = [
        { label: 'Dashboard', icon: 'bx-grid-alt', href: '/dashboard' },
        { label: 'Notification', icon: 'bx-bell', href: '/notifications' },
        { label: 'Subscription', icon: 'bx-credit-card', href: '/subscription', badge: 'PRO' },
        { label: 'Settings', icon: 'bx-cog', href: '/settings' },
    ];

    const bottomItems = [
        { label: 'Keyboard shortcut', icon: 'bx-command', href: '/shortcuts' },
        { label: 'Help center', icon: 'bx-help-circle', href: '/help' },
    ];
</script>

<div class="flex flex-col h-full w-full overflow-hidden">
    <!-- Header: Re-added User Info -->
    <div class="px-5 pt-2 pb-4 flex items-center justify-between border-b border-white/5">
        <div class="flex flex-col gap-0.5">
            <span class="text-[15px] font-bold text-white leading-tight">
                {userStore.user?.name || 'User'}
            </span>
            <span class="text-[13px] text-white/40 font-medium">
                {userStore.user?.email || 'email@example.com'}
            </span>
        </div>
        <!-- Avatar with ring -->
        <div class="relative group cursor-default">
            <div class="absolute inset-0 rounded-full bg-linear-to-tr from-pink-500 via-orange-400 to-emerald-400 p-[2px] animate-gradient-xy">
                <div class="w-full h-full rounded-full bg-black p-px">
                    <img 
                        src={userStore.user?.avatarUrl} 
                        alt="" 
                        class="w-full h-full rounded-full object-cover"
                    />
                </div>
            </div>
            <div class="h-10 w-10"></div>
        </div>
    </div>

    <!-- Upgrade Banner - Slim & Premium on Dark -->
    <div class="px-5 py-4">
        <div
            class="relative group cursor-pointer overflow-hidden rounded-xl p-3 transition-transform active:scale-[0.98]"
            style="background: linear-gradient(to right, #90D5EC, #FDC365, #F484A0, #C1A0D8)"
        >
            <div class="relative z-10 flex items-center justify-center">
                <div class="flex items-center gap-2 h-8">
                    <div class="flex h-8 items-center justify-center">
                        <i class='bx bxs-crown text-black text-[16px]'></i>
                    </div>
                    <span class="text-sm font-bold text-black tracking-tight">Upgrade profile</span>
                </div>
                <!-- <div class="px-2 py-0.5 rounded-full bg-white text-emerald-600 text-[9px] font-black uppercase tracking-widest shadow-sm">
                    PRO
                </div> -->
            </div>
            <!-- subtle shine effect -->
            <div class="absolute inset-0 bg-linear-to-r from-transparent via-white/30 to-transparent -translate-x-full group-hover:translate-x-full transition-transform duration-1000"></div>
        </div>
    </div>

    <!-- Main Menu -->
    <div class="px-2 py-1">
        {#each menuItems as item}
            <a 
                href={item.href}
                class="flex items-center justify-between px-4 py-3 rounded-xl hover:bg-white/5 transition-all group"
            >
                <div class="flex items-center gap-3">
                    <div class="flex h-5 w-5 items-center justify-center text-white/40 group-hover:text-white transition-colors">
                        <i class='bx {item.icon} text-xl'></i>
                    </div>
                    <span class="text-[14px] font-semibold text-white/70 group-hover:text-white transition-colors">{item.label}</span>
                </div>
                {#if item.badge}
                    <div
                        class="px-2 py-0.5 rounded-md text-black text-[9px] font-black uppercase tracking-wider"
                        style="background: linear-gradient(to right, #90D5EC, #FDC365, #F484A0, #C1A0D8)"
                    >
                        {item.badge}
                    </div>
                {/if}
            </a>
        {/each}
    </div>

    <div class="h-px bg-white/5 mx-5 mt-2 mb-1"></div>

    <!-- Bottom Menu -->
    <div class="px-2 py-2">
        {#each bottomItems as item}
            <a 
                href={item.href}
                class="flex items-center gap-3 px-4 py-3 rounded-xl hover:bg-white/5 transition-all group"
            >
                <div class="flex h-5 w-5 items-center justify-center text-white/40 group-hover:text-white transition-colors">
                    <i class='bx {item.icon} text-xl'></i>
                </div>
                <span class="text-[14px] font-semibold text-white/70 group-hover:text-white transition-colors">{item.label}</span>
            </a>
        {/each}

        <!-- Sign Out -->
        <button 
            onclick={(e) => { e.stopPropagation(); onSignOut(); }}
            class="w-full flex items-center gap-3 px-4 py-3 rounded-xl hover:bg-red-500/20 transition-all group mt-1 cursor-pointer"
        >
            <div class="flex h-5 w-5 items-center justify-center text-red-500/80 group-hover:text-red-500 transition-colors rotate-180">
                <i class='bx bx-log-out-circle text-xl'></i>
            </div>
            <span class="text-[14px] font-semibold text-red-500/80 group-hover:text-red-500 transition-colors">Sign out</span>
        </button>
    </div>
</div>

<style>
    @keyframes gradient-xy {
        0%, 100% {
            background-size: 400% 400%;
            background-position: left center;
        }
        50% {
            background-size: 200% 200%;
            background-position: right center;
        }
    }
    .animate-gradient-xy {
        animation: gradient-xy 6s ease infinite;
    }
</style>
