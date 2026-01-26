<script lang="ts">
	import SendioLogo from '$lib/components/icons/SendioLogo.svelte';
</script>

<div class="pb-12 flex justify-between items-center text-white/10 border-t border-white/5 pt-6">
    <div class="flex items-center gap-2">
        <!-- logo -->
        <a href="/" class="flex items-center gap-2.5 text-xl font-bold tracking-tight text-white">
            <div class="flex h-8 w-8 items-center justify-center rounded-lg bg-blue-600/5 text-blue-600">
                <SendioLogo class="h-6 w-6" />
            </div>
            <span class="text-white">Sendio</span>
        </a>
    </div>
    <div class="flex items-center gap-6">
        <span class="text-[12px] font-bold uppercase underline text-white/40">Setting</span>
    </div>
</div>
