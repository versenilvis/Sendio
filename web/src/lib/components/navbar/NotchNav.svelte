<script lang="ts">
	import { fade } from 'svelte/transition';
	import CollapsedState from './notch/CollapsedState.svelte';
	import MetricWidget from './notch/MetricWidget.svelte';
	import NavGrid from './notch/NavContent.svelte';
	import NavFooter from './notch/NavFooter.svelte';

	let isExpanded = $state(false);
	let isHovered = $state(false);
	let islandRef: HTMLElement;

	function handleKeydown(e: KeyboardEvent) {
		if (e.key === 'Escape') isExpanded = false;
	}

	let name = "Verse Admin"
</script>

<svelte:window 
	onclick={(e) => {
		if (islandRef && !islandRef.contains(e.target as Node)) {
			isExpanded = false;
		}
	}} 
	onkeydown={handleKeydown}
/>

<div class="fixed top-0 left-0 right-0 z-50 flex justify-center pointer-events-none select-none">
	<div class="relative notch-wrapper {isExpanded ? 'expanded' : ''}">
		<div 
			bind:this={islandRef}
			role="button"
			tabindex="0"
			aria-label="Toggle Control Center"
			onmouseenter={() => isHovered = true}
			onmouseleave={() => isHovered = false}
			onclick={(e) => { e.stopPropagation(); isExpanded = !isExpanded; }}
			onkeydown={(e) => { if (e.key === 'Enter' || e.key === ' ') { e.preventDefault(); isExpanded = !isExpanded; } }}
			class="relative flex flex-col bg-black pointer-events-auto transition-[width,height,border-radius,padding] duration-400 ease-[cubic-bezier(0.2,1,0.2,1)] overflow-hidden outline-none
				{isExpanded ? 'w-[920px] h-[640px] rounded-b-[56px] pt-12 cursor-default' : isHovered ? 'w-[380px] h-[64px] rounded-b-[32px] cursor-pointer' : 'w-[320px] h-[52px] rounded-b-[24px] cursor-pointer'}"
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
					<CollapsedState {name} />
				</div>
			{/if}

			<!-- expanded state -->
			{#if isExpanded}
				<div 
					in:fade={{ duration: 250, delay: 150 }}
					out:fade={{ duration: 0 }}
					class="w-full px-12 relative z-10 text-left h-full flex flex-col justify-center"
				>
					<MetricWidget {name} />
					<NavGrid />
					<NavFooter />
				</div>
			{/if}
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
