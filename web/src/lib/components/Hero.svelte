<script lang="ts">
	import handModel from '$lib/assets/hand_model.png';
	import Heat from './icons/Heat.svelte';
	import Mail from './icons/Mail.svelte';
</script>

	
	<div class="w-full h-full min-h-screen grid grid-cols-1 lg:grid-cols-2 relative z-10">
		
		<div class="flex flex-col justify-center items-start pl-6 md:pl-[10%] pr-6 lg:pr-0 pt-32 pb-12 lg:py-0 space-y-10 z-30 h-full">
			<div class="flex flex-col space-y-6 w-full max-w-2xl">
				<h1 class="text-[48px] md:text-[64px] lg:text-[80px] xl:text-[96px] font-bold leading-[1.05] tracking-tight text-blue-600 animate-slide-up flex flex-wrap items-center gap-x-4">
					<span class="flex items-center gap-4">
						Emails
						<Mail class="h-24 w-24" />
					</span>
					<br class="hidden md:block" />
					<span class="text-black w-full">for developers</span>
				</h1>
				<p class="text-xl md:text-2xl font-light leading-relaxed text-gray-500 animate-slide-up [animation-delay:150ms] opacity-0">
					Simple to integrate, impossible to outgrow. <br/>
					Modern email infrastructure to keep contact with your clients.
				</p>
			</div>

			<div class="flex flex-col space-y-12 w-full animate-slide-up [animation-delay:300ms] opacity-0">
				<div class="flex flex-col sm:flex-row items-start sm:items-center gap-8">
					<a href="/dashboard" class="group flex items-center gap-4 rounded-full bg-black py-4 px-8 text-white transition-all hover:bg-black/80 hover:shadow-xl active:scale-95">
						<span class="text-xl md:text-2xl font-medium">Try for free</span>
						<Heat class="h-6 w-6" />
					</a>
					
					<a href="#docs" class="text-lg md:text-xl font-medium text-gray-600 hover:text-black transition-colors px-4">
						Documentation
					</a>
				</div>
			</div>
		</div>

		<div class="relative w-full h-[50vh] lg:h-full flex items-end justify-center lg:justify-end pointer-events-none lg:overflow-visible overflow-hidden">
			<div class="absolute bottom-0 right-0 w-[90%] md:w-[70%] lg:w-[145%] xl:w-[135%] max-w-[900px] lg:max-w-none lg:h-[95%] flex items-end justify-end">
				<img 
					src={handModel} 
					alt="Sendio App Dashboard" 
					class="w-full h-full object-contain object-bottom drop-shadow-[-40px_20px_60px_rgba(0,0,0,0.15)] bg-blend-overlay animate-fade-in translate-x-[10%] lg:translate-x-0"
				/>
				<!-- White gradient overlay to blend the hand model 
				<div class="absolute inset-x-0 bottom-0 h-[25%] bg-linear-to-t from-white via-white/40 to-transparent z-20 pointer-events-none"></div>
				<div class="absolute inset-y-0 left-0 w-[8%] bg-linear-to-r from-white via-white/20 to-transparent z-20 pointer-events-none"></div> -->
			</div>
		</div>

	</div>



<style>
	@keyframes slideUp {
		from { opacity: 0; transform: translateY(30px); }
		to { opacity: 1; transform: translateY(0); }
	}
	@keyframes fadeIn {
		from { opacity: 0; transform: scale(0.98) translateY(20px); }
		to { opacity: 1; transform: scale(1) translateY(0); }
	}
	.animate-slide-up {
		animation: slideUp 0.8s cubic-bezier(0.2, 1, 0.2, 1) forwards;
	}
	.animate-fade-in {
		animation: fadeIn 1.2s cubic-bezier(0.2, 1, 0.2, 1) forwards;
	}
</style>
