<script lang="ts">
	import { fade } from 'svelte/transition';
	let withSendio = $state(true);
</script>

<section class="w-full bg-black py-16 relative overflow-hidden flex flex-col items-center select-none">
	
	<!-- line -->
	<div class="absolute inset-x-0 top-[45%] h-1 pointer-events-none transition-all duration-1000 ease-in-out z-0"
		style="transform: translateY(-50%) rotate({withSendio ? '-3deg' : '3deg'}) scaleX(1.5); 
			background: {withSendio ? '#10b981' : '#ef4444'}; 
			box-shadow: 0 0 50px {withSendio ? 'rgba(16,185,129,0.5)' : 'rgba(239,68,68,0.5)'};">
	</div>

	<!-- header and desc -->
	<div class="relative z-10 text-center px-6 min-h-[120px] flex flex-col items-center justify-center">
		<h2 class="text-4xl md:text-6xl font-black text-white mb-4 transition-all duration-300 max-w-3xl mx-auto tracking-tight">
			Build 
			<span class="transition-colors duration-500 {withSendio ? 'text-emerald-400' : 'text-red-500'}">
				{withSendio ? 'without the friction' : 'until they break'}
			</span>
		</h2>
		<p class="text-white opacity-80 font-medium text-lg md:text-xl transition-opacity duration-500">
			{withSendio ? 'Sendio maintains your infrastructure 24/7' : 'Standard API limits and unstable connections'}
		</p>
	</div>

	<!-- toggle buton -->
	<div class="relative z-10 flex items-center gap-6 mb-16">
		<span class="text-sm font-bold uppercase tracking-wider {withSendio ? 'text-white' : 'text-zinc-700'} transition-colors">With</span>
		<button 
			onclick={() => withSendio = !withSendio}
			aria-label="Toggle Comparison"
			class="w-16 h-8 rounded-full bg-zinc-900 border border-white/10 p-1.5 transition-all duration-300 relative flex items-center group cursor-pointer"
		>
			<div class="w-5 h-5 rounded-full bg-white transition-all duration-500 ease-[cubic-bezier(0.34,1.56,0.64,1)] shadow-xl"
				style="transform: translateX({withSendio ? '0' : '32px'});">
			</div>
		</button>
		<span class="text-sm font-bold uppercase tracking-wider {!withSendio ? 'text-white' : 'text-zinc-700'} transition-colors">Without</span>
	</div>

	<!-- Main -->
	<div class="relative z-10 w-full max-w-6xl grid grid-cols-1 lg:grid-cols-2 gap-8 px-6">
		
		<!-- left panel -->
		<div class="rounded-2xl bg-zinc-900/50 backdrop-blur-2xl border border-white/5 p-8 flex flex-col relative overflow-hidden h-[500px] transition-all duration-500 shadow-2xl">
			
			<div class="flex items-center justify-between mb-8">
				<div class="flex items-center gap-3">
					<span class="text-lg font-bold text-white uppercase tracking-tight">Status</span>
					<div class="px-2.5 py-0.5 rounded-full text-[9px] font-black uppercase tracking-widest transition-colors duration-500
						{withSendio ? 'bg-emerald-500 text-black' : 'bg-red-500 text-white animate-pulse'}">
						{withSendio ? 'Healthy' : 'Degraded'}
					</div>
					<div class="px-2.5 py-0.5 rounded-full text-[9px] font-black uppercase tracking-widest bg-white/5 text-white/40 border border-white/5">
						{withSendio ? '+1,000 clients' : '+0 client'}
					</div>
				</div>
				{#if !withSendio}
					<span class="text-red-500 font-black tracking-tighter text-sm">$5,920 wasted</span>
				{/if}
			</div>

			<div class="space-y-4 flex-1 overflow-y-auto pr-2 custom-scrollbar">
				<div class="text-white/20 font-bold uppercase tracking-wider text-[9px] mb-4 transition-opacity duration-500">Log output:</div>
				
				{#each [
					'Establishing idempotent connection...',
					'Verifying HMAC signatures...',
					'Analyzing delivery patterns...',
					'Executing smart-backoff retry...',
					'Mapping dynamic templates...',
					'Zero-latency pulse initiated...'
				] as item}
					<div class="flex gap-3 items-center">
						<div class="w-1 h-1 rounded-full shrink-0 transition-colors duration-500 {withSendio ? 'bg-emerald-500 shadow-[0_0_10px_#10b981]' : 'bg-red-500 shadow-[0_0_10px_#ef4444]'}"></div>
						<p class="text-gray-400 text-[12px] font-medium tracking-tight">{item}</p>
					</div>
				{/each}
			</div>

			<!-- badges overlay -->
			<div class="absolute inset-0 pointer-events-none">
				<!-- badge 1 - only with -->
				{#if withSendio}
					<div in:fade={{duration: 400}} out:fade={{duration: 200}} 
						class="absolute top-24 -right-2 rotate-15 px-3 py-1.5 rounded-full text-[9px] font-bold uppercase tracking-wide bg-emerald-500 text-black shadow-lg">
						army of experts
					</div>
				{/if}
				
				<!-- badge 2 - dynamic -->
				<div class="absolute bottom-24 mb-10 -left-1 px-3 py-1.5 rounded-full text-[9px] font-bold uppercase tracking-wide transition-all duration-500
					{withSendio ? 'bg-emerald-400 text-black opacity-100' : 'bg-red-600 text-white translate-x-10 scale-105'} shadow-xl">
					{withSendio ? 'reliable as core' : 'unstable api credits'}
				</div>
			</div>

			{#if !withSendio}
				<div class="mt-6 space-y-2 animate-in fade-in slide-in-from-bottom-5 duration-500">
					<div class="bg-red-500 text-white px-4 py-2 rounded-xl font-bold uppercase tracking-wider text-[9px] flex items-center justify-between shadow-2xl">
						<span>EXHAUSTED</span>
					</div>
					<div class="bg-zinc-900 border border-red-500/30 p-4 rounded-xl flex items-center gap-3">
						<i class='bx bx-error text-red-500 text-xl'></i>
						<span class="text-[10px] font-bold uppercase tracking-tight text-white">Critical Error</span>
					</div>
				</div>
			{/if}
		</div>

		<!-- right panel -->
		<div class="rounded-2xl border border-white/5 flex flex-col relative overflow-hidden transition-all duration-700 h-[500px]
			{withSendio ? 'bg-[#f8f9fa] shadow-xl' : 'bg-zinc-950 shadow-2xl'}">
			
			<!-- browser bar -->
			<div class="flex items-center gap-3 px-6 py-4 transition-colors duration-700 
				{withSendio ? 'bg-zinc-200/50' : 'bg-zinc-900/50'}">
				<div class="flex gap-1.5">
					<div class="w-2.5 h-2.5 rounded-full bg-[#ff5f57]"></div>
					<div class="w-2.5 h-2.5 rounded-full bg-[#febc2e]"></div>
					<div class="w-2.5 h-2.5 rounded-full bg-[#28c840]"></div>
				</div>
				<div class="flex-1 max-w-xs mx-auto h-7 rounded-lg flex items-center px-4 text-[10px] font-bold transition-all duration-700
					{withSendio ? 'bg-white text-black/20' : 'bg-black text-white/20'}">
					localhost:3000
				</div>
			</div>

			<!-- viewport content -->
			<div class="flex-1 p-8 flex flex-col relative">
				<div class="space-y-6 transition-all duration-700 {withSendio ? 'opacity-20' : 'opacity-10'}">
					<div class="h-8 w-1/3 rounded-lg {withSendio ? 'bg-black' : 'bg-white'}"></div>
					<div class="h-24 w-full rounded-xl {withSendio ? 'bg-black' : 'bg-white'}"></div>
					<div class="grid grid-cols-2 gap-4">
						<div class="h-32 rounded-xl {withSendio ? 'bg-black' : 'bg-white'}"></div>
						<div class="h-32 rounded-xl {withSendio ? 'bg-black' : 'bg-white'}"></div>
					</div>
				</div>

				<!-- center modal - with -->
				{#if withSendio}
					<div in:fade={{duration: 500}} out:fade={{duration: 200}}
						class="absolute inset-0 flex items-center justify-center p-8 transition-all duration-500">
						<div class="bg-emerald-500 text-black px-8 py-3 rounded-2xl text-xl font-bold uppercase tracking-tight shadow-2xl">
							Success
						</div>
					</div>
				{/if}

				<!-- center modal - without -->
				{#if !withSendio}
					<div in:fade={{duration: 500}} out:fade={{duration: 200}}
						class="absolute inset-0 flex items-center justify-center p-8 transition-all duration-700">
						<div class="bg-red-950/90 backdrop-blur-2xl border border-red-500/40 p-8 rounded-2xl text-center max-w-sm shadow-2xl">
							<i class='bx bx-error-alt text-red-500 text-5xl mb-4 animate-bounce'></i>
							<h3 class="text-red-500 text-lg font-bold uppercase tracking-tight mb-2">ERROR</h3>
							<p class="text-white/40 text-[11px] font-medium leading-relaxed">
								Fatal memory leak.
							</p>
						</div>
					</div>
				{/if}

				<!-- floating badges - right side -->
				<div class="absolute top-[20%] px-4 py-2 rounded-full text-[9px] font-bold uppercase tracking-wide transition-all duration-700
					{withSendio ? 'bg-emerald-500 text-black shadow-lg opacity-100 rotate-6 translate-x-0' : 'bg-red-500 text-white -rotate-12 translate-x-10 opacity-100'}">
					{withSendio ? 'uptime' : 'err'}
				</div>

				<div class="absolute bottom-12 left-1/2 -translate-x-1/2 px-5 py-2 rounded-full text-[10px] font-bold uppercase tracking-wide transition-all duration-500
					{withSendio ? 'bg-emerald-400 text-black opacity-100 shadow-xl' : 'bg-red-600 text-white translate-y-12'}">
					{withSendio ? 'autonomous' : 'failed'}
				</div>
			</div>
		</div>
	</div>

</section>

<style>
	.custom-scrollbar::-webkit-scrollbar {
		width: 3px;
	}
	.custom-scrollbar::-webkit-scrollbar-thumb {
		background: rgba(255, 255, 255, 0.1);
		border-radius: 10px;
	}
</style>

