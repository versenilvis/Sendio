<script lang="ts">
	import { fade } from 'svelte/transition';
	import { tick } from 'svelte';
	import CollapsedState from './notch/CollapsedState.svelte';
	import MetricWidget from './notch/MetricWidget.svelte';
	import NavGrid from './notch/NavContent.svelte';
	import SendioLogo from '$lib/components/icons/SendioLogo.svelte';

	let isExpanded = $state(false);
	let isHovered = $state(false);
	let islandRef: HTMLElement;
	let searchInput = $state<HTMLInputElement>();

	async function handleSearchClick() {
		isExpanded = true;
		await tick();
		if (searchInput) searchInput.focus();
	}

	function handleKeydown(e: KeyboardEvent) {
		if (e.key === 'Escape') isExpanded = false;
	}

	let name = "Admin Verse"
</script>

<svelte:window 
	onclick={(e) => {
		if (islandRef && !islandRef.contains(e.target as Node)) {
			isExpanded = false;
		}
	}} 
	onkeydown={handleKeydown}
/>

<div class="fixed top-0 left-0 right-0 z-50 pointer-events-none select-none">
	<!-- brand -->
	<!-- <div class="absolute left-4 top-2">
		<a href="/" class="pointer-events-auto flex items-center group" aria-label="Home">
			<div class="flex h-12 w-12 items-center justify-center rounded-lg text-black transition-transform">
				<SendioLogo class="h-8 w-8" />
			</div>
			<span class="text-2xl font-black uppercase text-black">Sendio</span>
		</a>
	</div> -->

	<!-- Notch Section -->
	<div class="flex justify-center">
		<div class="relative notch-wrapper {isExpanded ? 'expanded' : ''}">
		<div 
			bind:this={islandRef}
			role="button"
			tabindex="0"
			aria-label="Toggle Control Center"
			onmouseenter={() => isHovered = true}
			onmouseleave={() => isHovered = false}
			onclick={(e) => { 
				if (!isExpanded) {
					e.stopPropagation(); 
					isExpanded = true; 
				}
			}}
			onkeydown={(e) => { 
				if (!isExpanded && (e.key === 'Enter' || e.key === ' ')) { 
					e.preventDefault(); 
					isExpanded = true; 
				} 
			}}
			class="relative flex flex-col bg-black pointer-events-auto transition-[width,height,border-radius,padding] duration-400 ease-[cubic-bezier(0.2,1,0.2,1)] overflow-hidden outline-none
				{isExpanded ? 'w-[920px] h-[640px] rounded-b-[48px] pt-8 cursor-default' : isHovered ? 'w-[440px] h-[70px] rounded-b-[36px] cursor-pointer' : 'w-[380px] h-[60px] rounded-b-[30px] cursor-pointer'}"
		>
			<!-- subtle surface shine -->
			<div class="absolute inset-0 bg-linear-to-b from-white/8 to-transparent pointer-events-none"></div>

			<!-- collapsed state -->
			{#if !isExpanded}
				<div 
					in:fade={{ duration: 100, delay: 100 }}
					out:fade={{ duration: 0 }}
					class="w-full h-full"
				>
					<CollapsedState {name} onSearch={handleSearchClick} />
				</div>
			{/if}

			<!-- expanded state -->
			{#if isExpanded}
				<div 
					in:fade={{ duration: 250, delay: 150 }}
					out:fade={{ duration: 0 }}
					class="w-full px-12 relative z-10 text-left h-full flex flex-col"
				>
					<!-- Top Row: Brand & Compact Search -->
					<div class="flex items-center justify-between mb-8">
						<!-- logo -->
						<a href="/" class="flex items-center text-xl font-black tracking-tight text-white transition-opacity hover:opacity-80" onclick={(e) => e.stopPropagation()}>
							<div class="flex h-8 w-8 items-center justify-center rounded-lg">
								<SendioLogo class="h-6 w-6" />
							</div>
							<span class="text-white uppercase">Sendio</span>
						</a>

						<!-- Brand & Utilities -->
						<div class="flex items-center gap-5">
							

							<!-- Compact Search -->
							<div class="relative group w-48">
								<div class="absolute inset-y-0 left-4 flex items-center pointer-events-none text-white/80 group-focus-within:text-emerald-400 transition-colors">
									<i class='bx bx-search text-lg'></i>
								</div>
								<input 
									bind:this={searchInput}
									type="text" 
									placeholder="Quick search..." 
									class="w-full bg-white/5 border border-white/10 rounded-xl py-2.5 pl-11 pr-4 text-white placeholder:text-white/80 outline-none focus:bg-white/10 focus:border-emerald-500/50 transition-all text-xs font-medium"
									onclick={(e) => e.stopPropagation()}
								/>
								
								<!-- <div class="absolute inset-y-0 right-4 flex items-center pointer-events-none">
									<kbd class="flex items-center gap-0.5 px-1.5 py-0.5 rounded-md bg-white/5 border border-white/10 text-white/80 text-[8px] uppercase font-black tracking-tighter">
										<span class="text-[9px]">⌘</span>K
									</kbd>
								</div> -->
							</div>
							<button title="Settings" onclick={(e) => e.stopPropagation()} class="text-white/40 hover:text-white transition-all cursor-pointer">
								<i class='bx bx-cog text-2xl hover:rotate-90 transition-transform'></i>
							</button>
						</div>
					</div>

					<MetricWidget {name} />
					<NavGrid />
					
				</div>
			{/if}
		</div>
	</div>
</div>
</div>

<style>
	.notch-wrapper {
		filter: drop-shadow(0 30px 60px rgba(0,0,0,0.8));
		-webkit-filter: drop-shadow(0 30px 60px rgba(0,0,0,0.8));
	}

	.notch-wrapper::before,
	.notch-wrapper::after {
		content: "";
		position: absolute;
		top: 0;
		width: 40px;
		height: 40px;
		background: black;
		pointer-events: none;
		z-index: -1;
		transition: all 0.4s cubic-bezier(0.2, 1, 0.2, 1);
		opacity: 1;
	}

	.notch-wrapper.expanded::before,
	.notch-wrapper.expanded::after {
		opacity: 0;
		transform: translateY(-100%);
	}

	.notch-wrapper::before {
		left: -39px;
		mask: radial-gradient(circle at 0 40px, transparent 40px, black 41px);
		-webkit-mask: radial-gradient(circle at 0 40px, transparent 40px, black 41px);
	}

	.notch-wrapper::after {
		right: -39px;
		mask: radial-gradient(circle at 40px 40px, transparent 40px, black 41px);
		-webkit-mask: radial-gradient(circle at 40px 40px, transparent 40px, black 41px);
	}

</style>
