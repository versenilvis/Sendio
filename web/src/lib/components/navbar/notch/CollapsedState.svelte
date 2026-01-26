<script lang="ts">
	import avatar from '$lib/assets/avatar.png';
    let { name } = $props<{ name: string }>();
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

    <!-- name -->
    <div class="flex justify-center items-center overflow-hidden">
        <span class="text-[14px] font-black tracking-tight text-white/95 truncate leading-none">
            {name}
        </span>
    </div>

    <!-- noti bell -->
    <div class="flex items-center justify-end">
        <div class="flex items-center text-orange-300">
            <i class="bx bx-bell text-[21px] leading-none"></i>
        </div>
    </div>
</div>
