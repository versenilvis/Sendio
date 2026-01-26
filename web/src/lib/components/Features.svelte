<script lang="ts">
	
	let { 
		theme = 'dark',
		title = 'Deeper integration.',
		desc = 'Every interaction, refined. Powering the next generation of email infrastructure.',
		linkText = 'Learn more'
	} = $props();
</script>

<section class="w-full h-[540px] md:h-[600px] flex flex-col items-center pt-16 text-center overflow-hidden relative
	{theme === 'dark' ? 'bg-black text-white' : 'bg-[#f5f5f7] text-black'}">
	
	<div class="z-20 flex flex-col items-center px-6">

		<h2 class="text-4xl md:text-6xl font-bold mb-4 tracking-tight">
			{title}
		</h2>
		
		<p class="max-w-md text-lg md:text-xl font-medium mb-8 leading-relaxed opacity-80">
			{desc}
		</p>

		<div class="flex items-center gap-8">
			<a 
				href="/features" 
				class="px-6 py-2 text-sm md:text-base font-bold rounded-full transition-all hover:scale-105 active:scale-95 shadow-xl
				{theme === 'dark' ? 'bg-white text-black hover:bg-white/90' : 'bg-blue-600 text-white hover:bg-blue-700'}"
			>
				{linkText}
			</a>
		</div>
	</div>

	<div class="absolute inset-x-0 bottom-[-5%] h-1/2 md:h-2/3 flex items-end justify-center pointer-events-none">
		<div class="w-full max-w-5xl h-full rounded-t-[60px] md:rounded-t-[80px] border-t border-x flex items-center justify-center overflow-hidden
			{theme === 'dark' 
				? 'bg-linear-to-b from-white/10 to-transparent border-white/10' 
				: 'bg-linear-to-b from-black/5 to-transparent border-black/5'}">
			
			{#if theme === 'dark'}
				<div class="absolute inset-0 bg-radial from-blue-500/10 to-transparent"></div>
			{/if}

			<span class="italic text-2xl md:text-4xl font-black uppercase tracking-[0.3em] select-none opacity-5">
				Visualization
			</span>
		</div>
	</div>
</section>

<style>
</style>
