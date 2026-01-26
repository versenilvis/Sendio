<script lang="ts">
	import { fade, slide } from 'svelte/transition';
	import { cubicOut } from 'svelte/easing';

	let { open = false, children } = $props();
</script>

{#if open}
	<div
		transition:fade={{ duration: 200 }}
		class="glass fixed inset-0 top-12 z-40 border-t border-(--color-border)/20 md:hidden"
	>
		<div
			transition:slide={{ duration: 400, easing: cubicOut }}
			class="flex h-full flex-col justify-between p-6"
		>
			<div class="space-y-6">
				{@render children()}
			</div>
			
			<div class="flex flex-col gap-3 pb-8">
				<button class="h-12 w-full rounded-2xl border border-(--color-border) bg-(--color-bg-card) font-bold">
					Sign In
				</button>
				<button class="h-12 w-full rounded-2xl bg-(--color-accent) font-bold text-white">
					Get Started
				</button>
			</div>
		</div>
	</div>
{/if}
